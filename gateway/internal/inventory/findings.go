package inventory

import (
	"context"
	"fmt"
	"strings"
)

// findings answers the scope questions that depend only on the dictionary
// (#1, #2, #11, #12, #13, #14 of docs/escopo-v1.md) with evidence.
func findings(ctx context.Context, q *querier, res *Result) []Finding {
	var out []Finding
	if t := res.Table("PCDATAS"); t != nil && t.Dict != nil {
		f := Finding{ID: "1", Title: "PCDATAS tem uma linha por data ou por filial?"}
		f.Evidence = append(f.Evidence, "Colunas: "+columnNames(t))
		pk := "nenhuma"
		if k := t.Dict.PrimaryKey(); k != nil {
			pk = k.Name + " (" + strings.Join(k.Columns, ", ") + ")"
		}
		f.Evidence = append(f.Evidence, "Chave primária: "+pk)
		f.Evidence = append(f.Evidence, fmt.Sprintf("Coluna CODFILIAL existe: %v", t.Dict.Column("CODFILIAL") != nil))
		if t.Dict.Column("DATA") != nil {
			if total, distinct, err := q.distinctCount(ctx, qualifiedOf(t), "DATA"); err == nil {
				f.Evidence = append(f.Evidence, fmt.Sprintf("COUNT(*) = %d, COUNT(DISTINCT DATA) = %d", total, distinct))
				if total == distinct {
					f.Answer = "Uma linha por data: a data é única na tabela."
				} else {
					f.Answer = fmt.Sprintf("Mais de uma linha por data (%d linhas para %d datas); a tabela é por filial ou tem duplicatas.", total, distinct)
				}
			} else {
				f.Evidence = append(f.Evidence, "Contagem não executada: "+err.Error())
			}
		}
		if f.Answer == "" {
			f.Answer = "Depende da contagem (rodar sem --no-count)."
		}
		out = append(out, f)
	}

	f2 := Finding{ID: "2", Title: "Chaves reais: NUMTRANSITEM em PCMOV, NUMSEQ em PCPEDI e PCITEM, chave de PCSECAO"}
	for _, name := range []string{"PCMOV", "PCPEDI", "PCITEM", "PCSECAO"} {
		t := res.Table(name)
		if t == nil || t.Dict == nil {
			continue
		}
		line := name + ": "
		if k := t.Dict.PrimaryKey(); k != nil {
			line += "PK " + k.Name + " (" + strings.Join(k.Columns, ", ") + ")"
		} else {
			line += "sem PK"
		}
		var uniques []string
		for _, u := range t.Dict.UniqueConstraints() {
			uniques = append(uniques, u.Name+"("+strings.Join(u.Columns, ",")+")")
		}
		seen := map[string]bool{}
		for _, u := range t.Dict.Constraints {
			seen[u.Name] = true
		}
		for _, u := range t.Dict.UniqueIndexes() {
			if !seen[u.Name] {
				uniques = append(uniques, u.Name+"("+strings.Join(u.Columns, ",")+")")
			}
		}
		if len(uniques) > 0 {
			line += "; únicos: " + strings.Join(uniques, " ")
		}
		if k := t.Spec.PrimaryKey; k.VerifiedUnique != nil {
			line += fmt.Sprintf("; unicidade verificada da chave adotada (%s): %v", strings.Join(k.Columns, ", "), *k.VerifiedUnique)
			if k.NullKeyRows != nil && *k.NullKeyRows > 0 {
				line += fmt.Sprintf(" (%d linhas com chave nula)", *k.NullKeyRows)
			}
		}
		f2.Evidence = append(f2.Evidence, line)
	}
	if len(f2.Evidence) > 0 {
		f2.Answer = "Chaves adotadas por tabela: " + keysSummary(res, "PCMOV", "PCPEDI", "PCITEM", "PCSECAO")
		out = append(out, f2)
	}

	if t := res.Table("PCRATEIOCENTROCUSTO"); t != nil && t.Dict != nil {
		f := Finding{ID: "11", Title: "Colunas de PCRATEIOCENTROCUSTO (valor ou percentual de rateio)"}
		f.Evidence = append(f.Evidence, "Colunas: "+columnNames(t))
		var found []string
		for _, c := range t.Dict.Columns {
			up := strings.ToUpper(c.Name)
			if strings.Contains(up, "PERC") || strings.Contains(up, "VALOR") || strings.HasPrefix(up, "VL") {
				found = append(found, c.Name+" "+c.TypeString())
			}
		}
		if len(found) > 0 {
			f.Answer = "Colunas candidatas ao rateio: " + strings.Join(found, ", ") + ". Confirmar a semântica com o negócio."
		} else {
			f.Answer = "Nenhuma coluna com PERC/VALOR/VL no nome; ver a lista completa."
		}
		out = append(out, f)
	}

	f12 := Finding{ID: "12", Title: "Tipos reais das chaves e versão do Oracle"}
	f12.Evidence = append(f12.Evidence, "Versão: "+res.DBVersion)
	for _, col := range []string{"CODFILIAL", "CODCLI", "CODPROD", "CODUSUR", "NUMTRANSITEM", "NUMTRANSVENDA", "NUMPED", "RECNUM"} {
		types := map[string][]string{}
		var order []string
		for _, t := range res.Tables {
			if t.Dict == nil {
				continue
			}
			if c := t.Dict.Column(col); c != nil {
				ts := c.TypeString()
				if _, seen := types[ts]; !seen {
					order = append(order, ts)
				}
				types[ts] = append(types[ts], t.Spec.Table)
			}
		}
		if len(types) == 0 {
			continue
		}
		var parts []string
		for _, typ := range order {
			parts = append(parts, fmt.Sprintf("%s em %d tabelas (%s)", typ, len(types[typ]), strings.Join(types[typ], ", ")))
		}
		f12.Evidence = append(f12.Evidence, col+": "+strings.Join(parts, "; "))
	}
	f12.Answer = "Tipos por coluna acima; cod_filial é texto quando CODFILIAL for VARCHAR2 em todas as tabelas."
	out = append(out, f12)

	f13 := Finding{ID: "13", Title: "A nota de entrada guarda o pedido de compra de origem?"}
	for _, name := range []string{"PCNFENT", "PCMOV"} {
		t := res.Table(name)
		if t == nil || t.Dict == nil {
			continue
		}
		var cands []string
		for _, c := range t.Dict.Columns {
			if strings.Contains(strings.ToUpper(c.Name), "PED") {
				cands = append(cands, c.Name+" "+c.TypeString())
			}
		}
		f13.Evidence = append(f13.Evidence, name+": colunas com PED no nome: "+strings.Join(cands, ", "))
		if t.Dict.Column("NUMPED") != nil {
			if total, filled, err := q.fillRate(ctx, qualifiedOf(t), "NUMPED"); err == nil {
				f13.Evidence = append(f13.Evidence, fmt.Sprintf("%s.NUMPED preenchido em %d de %d linhas", name, filled, total))
			}
		}
	}
	if len(f13.Evidence) > 0 {
		f13.Answer = "Ver as colunas candidatas e a taxa de preenchimento; confirmar com o negócio qual delas liga a entrada ao pedido de compra."
		out = append(out, f13)
	}

	if t := res.Table("PCPRODUT"); t != nil && t.Dict != nil {
		f := Finding{ID: "14", Title: "O que PCPRODUT.DTVENC representa?"}
		if c := t.Dict.Column("DTVENC"); c != nil {
			f.Evidence = append(f.Evidence, fmt.Sprintf("DTVENC %s, comentário no dicionário: %q", c.TypeString(), c.Comment))
			if total, filled, err := q.fillRate(ctx, qualifiedOf(t), "DTVENC"); err == nil {
				f.Evidence = append(f.Evidence, fmt.Sprintf("Preenchida em %d de %d produtos", filled, total))
			}
			if rs, err := q.rangeStats(ctx, qualifiedOf(t), "DTVENC", "YYYY-MM-DD"); err == nil {
				f.Evidence = append(f.Evidence, fmt.Sprintf("Mínimo %s, máximo %s", rs.Min, rs.Max))
			}
			f.Answer = "O dicionário não define a semântica; a taxa de preenchimento e a faixa de datas indicam se é validade de produto (poucos valores, futuros) ou outra data. Confirmar com o negócio."
		} else {
			f.Answer = "PCPRODUT não tem a coluna DTVENC nesta base."
		}
		out = append(out, f)
	}
	fw := Finding{ID: "Q9", Title: "Colunas de controle de alteração (candidatas a watermark secundário)"}
	for _, t := range res.Tables {
		if t.Dict == nil {
			continue
		}
		var cands []string
		for _, c := range t.Dict.Columns {
			up := strings.ToUpper(c.Name)
			if (strings.Contains(up, "ALTER") || strings.Contains(up, "ATUALIZ") || strings.Contains(up, "INCLUS") || strings.Contains(up, "CRIACAO")) && (c.DataType == "DATE" || strings.HasPrefix(c.DataType, "TIMESTAMP")) {
				cands = append(cands, c.Name)
			}
		}
		if len(cands) > 0 {
			fw.Evidence = append(fw.Evidence, t.Spec.Table+": "+strings.Join(cands, ", "))
		}
	}
	if len(fw.Evidence) == 0 {
		fw.Evidence = append(fw.Evidence, "Nenhuma tabela tem coluna DATE/TIMESTAMP com ALTER, ATUALIZ, INCLUS ou CRIACAO no nome.")
	}
	fw.Answer = "Só as tabelas listadas têm carimbo de alteração no dicionário; nas demais o watermark é a data de negócio mais a janela móvel. Resta validar, por teste controlado, se o carimbo muda a cada alteração."
	out = append(out, fw)
	return out
}

func columnNames(t *TableResult) string {
	var names []string
	for _, c := range t.Dict.Columns {
		names = append(names, c.Name+" "+c.TypeString())
	}
	return strings.Join(names, ", ")
}

func qualifiedOf(t *TableResult) string {
	return `"` + t.Dict.Owner + `"."` + t.Dict.Name + `"`
}

func keysSummary(res *Result, names ...string) string {
	var parts []string
	for _, n := range names {
		if t := res.Table(n); t != nil {
			parts = append(parts, fmt.Sprintf("%s = (%s) [%s]", n, strings.Join(t.Spec.PrimaryKey.Columns, ", "), t.Spec.PrimaryKey.Source))
		}
	}
	return strings.Join(parts, "; ")
}
