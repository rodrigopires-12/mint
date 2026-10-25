package doctor

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/the12company/mint/internal/catalog"
	"github.com/the12company/mint/internal/source"
	"github.com/the12company/mint/internal/source/oraerr"
)

// capabilities are optional privileges that unlock features in later
// deliverables (ROWID partitioning, AS OF SCN, session monitoring).
func (c *checker) capabilities(ctx context.Context, sysPrivs []string, version string) {
	probe := func(name, stmt, purpose string) {
		var n sql.NullString
		err := c.query(ctx, stmt, func(rows source.Rows) error { return rows.Scan(&n) })
		if err != nil {
			c.r.add("capacidade opcional", Info, "%s: não (%s). Uso: %s", name, oraerr.Humanize(err), purpose)
			return
		}
		c.r.add("capacidade opcional", Info, "%s: sim. Uso: %s", name, purpose)
	}
	hasFlashback := false
	for _, p := range sysPrivs {
		if p == "FLASHBACK ANY TABLE" {
			hasFlashback = true
		}
	}
	c.r.add("capacidade opcional", Info, "FLASHBACK ANY TABLE: %s. Uso: AS OF SCN (ADR 0004)", yesNo(hasFlashback))
	probe("DBMS_FLASHBACK", "SELECT TO_CHAR(DBMS_FLASHBACK.GET_SYSTEM_CHANGE_NUMBER) FROM DUAL", "obter o SCN para AS OF SCN (ADR 0004)")
	probe("DBA_EXTENTS", "SELECT TO_CHAR(COUNT(*)) FROM DBA_EXTENTS WHERE ROWNUM < 1", "particionar por faixas de ROWID (E1.5)")
	probe("V$SESSION", "SELECT TO_CHAR(COUNT(*)) FROM V$SESSION WHERE ROWNUM < 1", "conferir o teto de sessões (E1.5)")
	if version != "" {
		major := strings.SplitN(version, ".", 2)[0]
		feature := "só ROWNUM (11g)"
		if len(major) >= 2 && major >= "12" {
			feature = "FETCH FIRST disponível (12c+)"
		}
		c.r.add("capacidade opcional", Info, "versão major %s: %s", major, feature)
	}
}

func yesNo(b bool) string {
	if b {
		return "sim"
	}
	return "não"
}

// tables checks that every catalog table exists and can be read.
func (c *checker) tables(ctx context.Context, tables []string) {
	for _, t := range tables {
		if !catalog.ValidIdentifier(t) { // defence in depth: the catalog already validated
			c.r.add("tabelas", Fail, "nome de tabela inválido no catálogo: %q", t)
			return
		}
	}
	inList := "'" + strings.Join(tables, "','") + "'"

	owners := map[string][]string{}
	err := c.query(ctx, "SELECT table_name, owner FROM ALL_TABLES WHERE table_name IN ("+inList+") ORDER BY table_name, owner",
		func(rows source.Rows) error {
			var t, o string
			if err := rows.Scan(&t, &o); err != nil {
				return err
			}
			owners[t] = append(owners[t], o)
			return nil
		})
	if err != nil {
		c.r.add("tabelas", Fail, "ALL_TABLES: %s", oraerr.Humanize(err))
		return
	}
	synonyms := map[string]string{}
	err = c.query(ctx, "SELECT synonym_name, table_owner FROM ALL_SYNONYMS WHERE owner IN ('PUBLIC', USER) AND synonym_name IN ("+inList+")",
		func(rows source.Rows) error {
			var s string
			var o sql.NullString
			if err := rows.Scan(&s, &o); err != nil {
				return err
			}
			synonyms[s] = o.String
			return nil
		})
	if err != nil {
		c.r.add("tabelas", Warn, "ALL_SYNONYMS: %s", oraerr.Humanize(err))
	}

	ownerCount := map[string]int{}
	var ambiguous []string
	for _, t := range tables {
		qualifier, owner := c.qualify(t, owners[t], synonyms)
		if len(owners[t]) > 1 && c.cfg.Oracle.Schema == "" {
			ambiguous = append(ambiguous, fmt.Sprintf("%s (%s)", t, strings.Join(owners[t], ",")))
		}
		var one sql.NullString
		err := c.query(ctx, "SELECT 1 FROM "+qualifier+" WHERE ROWNUM < 1", func(rows source.Rows) error { return rows.Scan(&one) })
		if err != nil {
			c.r.TablesMissing = append(c.r.TablesMissing, t)
			c.log.Debug("table probe failed", "table", t, "qualifier", qualifier, "error", oraerr.Humanize(err))
			continue
		}
		c.r.TablesChecked++
		ownerCount[owner]++
	}
	if len(ambiguous) > 0 {
		c.r.add("tabelas", Warn, "mais de um owner para: %s (defina oracle.schema)", strings.Join(ambiguous, "; "))
	}
	var ownerParts []string
	for o, n := range ownerCount {
		ownerParts = append(ownerParts, fmt.Sprintf("%s=%d", o, n))
	}
	if len(c.r.TablesMissing) > 0 {
		c.r.add("tabelas", Fail, "%d/%d acessíveis; faltam ou sem leitura: %s", c.r.TablesChecked, len(tables), strings.Join(c.r.TablesMissing, ", "))
		return
	}
	c.r.add("tabelas", OK, "%d/%d acessíveis com SELECT (owner: %s)", c.r.TablesChecked, len(tables), strings.Join(ownerParts, ", "))
}

// qualify decides how to reference a table and reports the owner used.
func (c *checker) qualify(table string, owners []string, synonyms map[string]string) (string, string) {
	if s := strings.ToUpper(strings.TrimSpace(c.cfg.Oracle.Schema)); s != "" && catalog.ValidIdentifier(s) {
		return `"` + s + `"."` + table + `"`, s
	}
	if o, ok := synonyms[table]; ok {
		return `"` + table + `"`, o + " (sinônimo)"
	}
	if len(owners) >= 1 {
		return `"` + owners[0] + `"."` + table + `"`, owners[0]
	}
	return `"` + table + `"`, "?"
}

// catalogFiles warns about tables without an approved <TABLE>.yaml (E1.2).
func (c *checker) catalogFiles(tables []string) {
	var missing, notApproved, invalid int
	for _, t := range tables {
		spec, err := catalog.LoadTableSpec(c.cfg.SourcesDir, t)
		switch {
		case errors.Is(err, os.ErrNotExist):
			missing++
		case err != nil:
			invalid++
			c.log.Warn("invalid table spec", "table", t, "error", err.Error())
		case spec.Status != catalog.StatusApproved:
			notApproved++
		}
	}
	if missing+notApproved+invalid == 0 {
		c.r.add("arquivos de tabela", OK, "%d tabelas com arquivo aprovado", len(tables))
		return
	}
	c.r.add("arquivos de tabela", Warn, "%d sem arquivo, %d propostos (não aprovados), %d inválidos, de %d tabelas (E1.2)", missing, notApproved, invalid, len(tables))
}
