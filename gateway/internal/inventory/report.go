package inventory

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/the12company/mint/internal/catalog"
)

// WriteSpecs writes one <TABLE>.yaml per inventoried table. Approved files
// are preserved unless force is set.
func WriteSpecs(dir string, res *Result, force bool) (written, skipped []string, err error) {
	for _, t := range res.Tables {
		existing, lerr := catalog.LoadTableSpec(dir, t.Spec.Table)
		if lerr == nil && existing.Status == catalog.StatusApproved && !force {
			skipped = append(skipped, t.Spec.Table)
			continue
		}
		if lerr != nil && !errors.Is(lerr, os.ErrNotExist) {
			// An invalid existing file is overwritten, but noted.
			t.Spec.Questions = append(t.Spec.Questions, "Arquivo anterior inválido foi substituído: "+lerr.Error())
		}
		if err := catalog.SaveTableSpec(dir, t.Spec); err != nil {
			return written, skipped, err
		}
		written = append(written, t.Spec.Table)
	}
	return written, skipped, nil
}

// WriteDictionaries writes docs/inventory/<TABLE>.md with the full column
// list of each table (names and types only, no data).
func WriteDictionaries(dir string, res *Result) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for _, t := range res.Tables {
		if t.Dict == nil {
			continue
		}
		var b bytes.Buffer
		fmt.Fprintf(&b, "# %s.%s — dicionário\n\nGerado por `gateway inventory` em %s. %d colunas.\n\n", t.Dict.Owner, t.Dict.Name, res.InventoriedAt.Format("2006-01-02"), len(t.Dict.Columns))
		b.WriteString("| # | Coluna | Tipo Oracle | Nulo | Comentário |\n|---|---|---|---|---|\n")
		for _, c := range t.Dict.Columns {
			null := "N"
			if c.Nullable {
				null = "S"
			}
			fmt.Fprintf(&b, "| %d | %s | %s | %s | %s |\n", c.ID, c.Name, c.TypeString(), null, mdEscape(c.Comment))
		}
		b.WriteString("\n## Chaves e índices\n\n")
		for _, k := range t.Dict.Constraints {
			fmt.Fprintf(&b, "- Constraint %s (%s): %s\n", k.Name, k.Type, strings.Join(k.Columns, ", "))
		}
		for _, i := range t.Dict.Indexes {
			u := ""
			if i.Unique {
				u = " UNIQUE"
			}
			fmt.Fprintf(&b, "- Índice %s%s: %s\n", i.Name, u, strings.Join(i.Columns, ", "))
		}
		if err := os.WriteFile(filepath.Join(dir, t.Spec.Table+".md"), b.Bytes(), 0o644); err != nil {
			return err
		}
	}
	return nil
}

// WriteReport writes docs/tables-inventory.md.
func WriteReport(path string, res *Result, dictDir string) error {
	var b bytes.Buffer
	fmt.Fprintf(&b, "# Inventário das tabelas de origem (E1.2)\n\nGerado por `gateway inventory` em %s contra %s. Os arquivos por tabela estão em `sources/winthor/<TABELA>.yaml` (status `proposed` até a aprovação) e a lista completa de colunas de cada tabela em `%s/<TABELA>.md`.\n\n", res.InventoriedAt.Format("2006-01-02 15:04"), res.DBVersion, filepath.ToSlash(dictDir))
	b.WriteString("## Resumo\n\n| Tabela | Onda | Owner | Linhas (dicionário) | COUNT(*) | MB est. | Chave | Estratégia | Watermark | Partição | Riscos |\n|---|---|---|---|---|---|---|---|---|---|---|\n")
	for _, t := range res.Tables {
		s := t.Spec
		fmt.Fprintf(&b, "| %s | %d | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			s.Table, s.Wave, s.Owner, i64(s.Stats.NumRowsDictionary), i64(s.Stats.Count), f64(s.Stats.EstimatedMB),
			keyCell(s), s.Strategy, refCell(s.Watermark), refCell(s.Partition), riskCell(s))
	}

	b.WriteString("\n## Respostas aos pontos do escopo que dependem só do dicionário\n\n")
	for _, f := range res.Findings {
		fmt.Fprintf(&b, "### #%s — %s\n\n", f.ID, f.Title)
		for _, e := range f.Evidence {
			fmt.Fprintf(&b, "- %s\n", mdEscape(e))
		}
		fmt.Fprintf(&b, "\n**Resposta:** %s\n\n", mdEscape(f.Answer))
	}

	b.WriteString("## Detalhe por tabela\n\n")
	for _, t := range res.Tables {
		s := t.Spec
		fmt.Fprintf(&b, "### %s (onda %d) — %s\n\n", s.Table, s.Wave, s.Description)
		if t.Err != nil {
			fmt.Fprintf(&b, "**Não inventariada:** %s\n\n", mdEscape(t.Err.Error()))
			continue
		}
		fmt.Fprintf(&b, "- Owner `%s`, %d colunas no dicionário, %d extraídas. Estratégia `%s`", s.Owner, s.Stats.DictionaryColumns, len(s.Columns), s.Strategy)
		if s.Watermark != nil {
			fmt.Fprintf(&b, ", watermark `%s` (%s)", s.Watermark.Column, s.Watermark.Kind)
		}
		if s.ReprocessDays > 0 {
			fmt.Fprintf(&b, ", reprocessamento de %d dias", s.ReprocessDays)
		}
		if len(s.OpenWhenNull) > 0 {
			fmt.Fprintf(&b, ", em aberto quando `%s` nulo", strings.Join(s.OpenWhenNull, "`, `"))
		}
		b.WriteString(".\n")
		fmt.Fprintf(&b, "- Chave: %s", keyCell(s))
		if s.PrimaryKey.VerifiedUnique != nil {
			fmt.Fprintf(&b, "; unicidade verificada: %v (%d grupos duplicados, %s linhas com chave nula)", *s.PrimaryKey.VerifiedUnique, *s.PrimaryKey.DuplicateGroups, i64(s.PrimaryKey.NullKeyRows))
		}
		b.WriteString(".\n")
		fmt.Fprintf(&b, "- Linhas: dicionário %s (analisado em %s), COUNT(*) %s, tamanho estimado %s MB, linha média %s bytes.\n", i64(s.Stats.NumRowsDictionary), nz(s.Stats.LastAnalyzed), i64(s.Stats.Count), f64(s.Stats.EstimatedMB), i64(s.Stats.AvgRowLen))
		if s.Stats.WatermarkMin != "" || s.Stats.WatermarkMax != "" {
			fmt.Fprintf(&b, "- Watermark: mínimo %s, máximo %s, nulos %s.\n", s.Stats.WatermarkMin, s.Stats.WatermarkMax, i64(s.Stats.WatermarkNulls))
		}
		if len(t.ByMonth) > 0 {
			b.WriteString("- Distribuição por mês (watermark): ")
			var parts []string
			for _, m := range t.ByMonth {
				parts = append(parts, fmt.Sprintf("%s=%d", m.Month, m.Rows))
			}
			b.WriteString(strings.Join(parts, ", ") + "\n")
		}
		if len(s.MissingColumns) > 0 {
			fmt.Fprintf(&b, "- **Colunas do escopo que não existem na base:** %s\n", strings.Join(s.MissingColumns, ", "))
		}
		if len(s.Risks) > 0 {
			fmt.Fprintf(&b, "- Riscos: %s\n", mdEscape(strings.Join(s.Risks, "; ")))
		}
		for _, q := range s.Questions {
			fmt.Fprintf(&b, "- Pergunta: %s\n", mdEscape(q))
		}
		b.WriteString("\n| Coluna | Oracle | ClickHouse | Nulo | Uso | Destino | Nota |\n|---|---|---|---|---|---|---|\n")
		for _, c := range s.Columns {
			null := ""
			if c.Nullable {
				null = "S"
			}
			fmt.Fprintf(&b, "| %s | %s | %s | %s | %s | %s | %s |\n", c.Name, c.OracleType, c.ClickHouseType, null, c.Use, c.Ref, mdEscape(c.Note))
		}
		b.WriteString("\n")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, b.Bytes(), 0o644)
}

func keyCell(s *catalog.TableSpec) string {
	if len(s.PrimaryKey.Columns) == 0 {
		return "(nenhuma)"
	}
	return strings.Join(s.PrimaryKey.Columns, ", ") + " [" + s.PrimaryKey.Source + "]"
}

func refCell(r *catalog.ColumnRef) string {
	if r == nil {
		return ""
	}
	return r.Column + " (" + r.Kind + ")"
}

func riskCell(s *catalog.TableSpec) string {
	var short []string
	for _, r := range s.Risks {
		if i := strings.IndexAny(r, ":("); i > 0 {
			r = strings.TrimSpace(r[:i])
		}
		short = appendUnique(short, r)
	}
	return strings.Join(short, ", ")
}

func i64(p *int64) string {
	if p == nil {
		return "?"
	}
	return fmt.Sprintf("%d", *p)
}

func f64(p *float64) string {
	if p == nil {
		return "?"
	}
	return fmt.Sprintf("%.0f", *p)
}

func nz(s string) string {
	if s == "" {
		return "?"
	}
	return s
}

func mdEscape(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "|", `\|`), "\n", " ")
}
