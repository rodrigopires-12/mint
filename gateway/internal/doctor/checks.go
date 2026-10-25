package doctor

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/the12company/mint/internal/config"
	"github.com/the12company/mint/internal/source"
	"github.com/the12company/mint/internal/source/oraerr"
)

type checker struct {
	src source.Source
	r   *Report
	cfg config.Config
	now func() time.Time
	log *slog.Logger
}

// query runs stmt with the configured timeout and feeds each row to scan.
func (c *checker) query(ctx context.Context, stmt string, scan func(source.Rows) error) error {
	qctx, cancel := context.WithTimeout(ctx, c.cfg.Oracle.QueryTimeout)
	defer cancel()
	rows, err := c.src.Query(qctx, stmt)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := scan(rows); err != nil {
			return err
		}
	}
	return rows.Err()
}

// addResult appends a pre-built Result (from the assess functions).
func (c *checker) addResult(res Result) {
	c.r.Results = append(c.r.Results, res)
	c.r.Summary[res.Status]++
}

func scanStrings(dst *[]string) func(source.Rows) error {
	return func(rows source.Rows) error {
		var s string
		if err := rows.Scan(&s); err != nil {
			return err
		}
		*dst = append(*dst, s)
		return nil
	}
}

func (c *checker) identity(ctx context.Context) {
	var user, db, host, inst sql.NullString
	err := c.query(ctx, "SELECT USER, SYS_CONTEXT('USERENV','DB_NAME'), SYS_CONTEXT('USERENV','SERVER_HOST'), SYS_CONTEXT('USERENV','INSTANCE_NAME') FROM DUAL",
		func(rows source.Rows) error { return rows.Scan(&user, &db, &host, &inst) })
	if err != nil {
		c.r.add("identidade", Fail, "%s", oraerr.Humanize(err))
		return
	}
	c.r.add("identidade", OK, "usuário=%s banco=%s instância=%s host=%s", user.String, db.String, inst.String, host.String)
}

// version returns the numeric database version (e.g. "19.0.0.0.0") or "".
func (c *checker) version(ctx context.Context) string {
	var product, version string
	err := c.query(ctx, "SELECT product, version FROM PRODUCT_COMPONENT_VERSION WHERE UPPER(product) LIKE '%ORACLE DATABASE%' AND ROWNUM < 2",
		func(rows source.Rows) error { return rows.Scan(&product, &version) })
	if err != nil || version == "" {
		var banner string
		err2 := c.query(ctx, "SELECT banner FROM V$VERSION WHERE ROWNUM < 2", func(rows source.Rows) error { return rows.Scan(&banner) })
		if err2 != nil {
			c.r.add("versão", Warn, "não foi possível ler a versão: %s", oraerr.Humanize(errors.Join(err, err2)))
			return ""
		}
		c.r.add("versão", OK, "%s", strings.TrimSpace(banner))
		return ""
	}
	c.r.add("versão", OK, "%s %s", strings.TrimSpace(product), version)
	return version
}

func (c *checker) charset(ctx context.Context) {
	params := map[string]string{}
	err := c.query(ctx, "SELECT parameter, value FROM NLS_DATABASE_PARAMETERS WHERE parameter IN ('NLS_CHARACTERSET','NLS_NCHAR_CHARACTERSET','NLS_LENGTH_SEMANTICS')",
		func(rows source.Rows) error {
			var k, v string
			if err := rows.Scan(&k, &v); err != nil {
				return err
			}
			params[k] = v
			return nil
		})
	if err != nil {
		c.r.add("charset", Fail, "%s", oraerr.Humanize(err))
		return
	}
	cs := params["NLS_CHARACTERSET"]
	detail := fmt.Sprintf("NLS_CHARACTERSET=%s NLS_NCHAR_CHARACTERSET=%s NLS_LENGTH_SEMANTICS=%s", cs, params["NLS_NCHAR_CHARACTERSET"], params["NLS_LENGTH_SEMANTICS"])
	switch cs {
	case "AL32UTF8", "UTF8":
		c.r.add("charset", OK, "%s", detail)
	case "":
		c.r.add("charset", Warn, "NLS_CHARACTERSET não encontrado")
	default:
		c.r.add("charset", Warn, "%s (não é UTF-8: a conversão para UTF-8 fica ativa na extração)", detail)
	}
}

func (c *checker) clock(ctx context.Context) {
	var dbUTC time.Time
	var dbTZ, sessTZ string
	before := c.now()
	err := c.query(ctx, "SELECT SYS_EXTRACT_UTC(SYSTIMESTAMP), DBTIMEZONE, SESSIONTIMEZONE FROM DUAL",
		func(rows source.Rows) error { return rows.Scan(&dbUTC, &dbTZ, &sessTZ) })
	if err != nil {
		c.r.add("relógio", Fail, "%s", oraerr.Humanize(err))
		return
	}
	// The driver may attach a local zone to a zone-less TIMESTAMP; only the
	// wall-clock fields matter here.
	wall := time.Date(dbUTC.Year(), dbUTC.Month(), dbUTC.Day(), dbUTC.Hour(), dbUTC.Minute(), dbUTC.Second(), dbUTC.Nanosecond(), time.UTC)
	skew := wall.Sub(before.UTC()).Round(time.Second)
	if skew < 0 {
		skew = -skew
	}
	st := OK
	if skew > 60*time.Second {
		st = Warn
	}
	c.r.add("relógio", st, "banco (UTC)=%s local (UTC)=%s desvio=%s DBTIMEZONE=%s SESSIONTIMEZONE=%s",
		wall.Format(time.RFC3339), before.UTC().Format(time.RFC3339), skew, dbTZ, sessTZ)
}

// privileges reports system privileges, roles and object privileges; it
// returns the system privileges for the capabilities check.
func (c *checker) privileges(ctx context.Context) []string {
	var sys []string
	if err := c.query(ctx, "SELECT privilege FROM SESSION_PRIVS ORDER BY privilege", scanStrings(&sys)); err != nil {
		c.r.add("privilégios de sistema", Fail, "%s", oraerr.Humanize(err))
	} else {
		c.addResult(AssessSystemPrivileges(sys))
	}

	var roles []string
	if err := c.query(ctx, "SELECT granted_role FROM USER_ROLE_PRIVS ORDER BY granted_role", scanStrings(&roles)); err != nil {
		c.r.add("papéis", Fail, "%s", oraerr.Humanize(err))
	} else {
		c.addResult(AssessRoles(roles))
	}

	var tab []TablePrivilege
	scanTab := func(rows source.Rows) error {
		var p TablePrivilege
		if err := rows.Scan(&p.Owner, &p.Table, &p.Privilege); err != nil {
			return err
		}
		tab = append(tab, p)
		return nil
	}
	err := c.query(ctx, "SELECT owner, table_name, privilege FROM USER_TAB_PRIVS WHERE grantee = USER ORDER BY owner, table_name, privilege", scanTab)
	if err == nil {
		err = c.query(ctx, "SELECT owner, table_name, privilege FROM ROLE_TAB_PRIVS ORDER BY owner, table_name, privilege", scanTab)
	}
	if err != nil {
		c.r.add("privilégios de objeto", Fail, "%s", oraerr.Humanize(err))
	} else {
		c.addResult(AssessTablePrivileges(tab))
	}
	return sys
}
