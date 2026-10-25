package doctor

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/the12company/mint/internal/config"
	"github.com/the12company/mint/internal/source"
	"github.com/the12company/mint/internal/source/sourcetest"
)

func TestAssessSystemPrivileges(t *testing.T) {
	if r := AssessSystemPrivileges([]string{"CREATE SESSION", "ALTER SESSION"}); r.Status != OK {
		t.Fatalf("%+v", r)
	}
	if r := AssessSystemPrivileges([]string{"CREATE SESSION", "SELECT ANY TABLE"}); r.Status != Warn {
		t.Fatalf("%+v", r)
	}
	// Own-schema creation privileges (common in Winthor databases) cannot
	// change the PC* tables: WARN with a recommendation to revoke.
	r := AssessSystemPrivileges([]string{"CREATE SESSION", "CREATE JOB", "CREATE CREDENTIAL", "UNLIMITED TABLESPACE", "CREATE TABLE"})
	if r.Status != Warn || !strings.Contains(r.Detail, "CREATE JOB") || !strings.Contains(r.Detail, "revogar") {
		t.Fatalf("%+v", r)
	}
	for _, p := range []string{"UPDATE ANY TABLE", "CREATE ANY TABLE", "CREATE EXTERNAL JOB", "CREATE PUBLIC SYNONYM", "ALTER SYSTEM", "EXECUTE ANY PROCEDURE", "DROP ANY TABLE"} {
		if r := AssessSystemPrivileges([]string{"CREATE SESSION", p}); r.Status != Fail || !strings.Contains(r.Detail, p) {
			t.Fatalf("%s: %+v", p, r)
		}
	}
}

func TestAssessRoles(t *testing.T) {
	if r := AssessRoles([]string{"CONNECT"}); r.Status != OK {
		t.Fatalf("%+v", r)
	}
	if r := AssessRoles([]string{"CONNECT", "APP_READ"}); r.Status != Warn {
		t.Fatalf("%+v", r)
	}
	if r := AssessRoles([]string{"CONNECT", "DBA"}); r.Status != Fail {
		t.Fatalf("%+v", r)
	}
}

func TestAssessTablePrivileges(t *testing.T) {
	ok := []TablePrivilege{{"WINT", "PCMOV", "SELECT"}, {"WINT", "PCPREST", "READ"}}
	if r := AssessTablePrivileges(ok); r.Status != OK {
		t.Fatalf("%+v", r)
	}
	bad := append(ok, TablePrivilege{"WINT", "PCMOV", "UPDATE"})
	if r := AssessTablePrivileges(bad); r.Status != Fail || !strings.Contains(r.Detail, "UPDATE em WINT.PCMOV") {
		t.Fatalf("%+v", r)
	}
}

func writeWaves(t *testing.T, dir string, tables ...string) config.Config {
	t.Helper()
	body := "version: 1\nwaves:\n  - id: 1\n    name: t\n    tables: [" + strings.Join(tables, ", ") + "]\n"
	if err := os.WriteFile(filepath.Join(dir, "_waves.yaml"), []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	cfg := config.Default()
	cfg.SourcesDir = dir
	cfg.Oracle.DSN, cfg.Oracle.User, cfg.Oracle.Password = "db:1521/SVC", "ro", "env:DOCTOR_TEST_PW"
	return cfg
}

func TestRunConnectionFailureIsCleanAndRedacted(t *testing.T) {
	t.Setenv("DOCTOR_TEST_PW", "topsecret")
	cfg := writeWaves(t, t.TempDir(), "PCFILIAL")
	connect := func(context.Context, source.ConnectParams) (source.Source, error) {
		return nil, errors.New("ORA-01017: invalid username/password; logon denied (password was topsecret)\ngoroutine 1 [running]:")
	}
	r := Run(context.Background(), Options{Config: cfg, Connect: connect})
	if !r.HasFail() {
		t.Fatal("expected FAIL")
	}
	var buf bytes.Buffer
	WriteText(&buf, r)
	out := buf.String()
	if strings.Contains(out, "topsecret") || strings.Contains(out, "goroutine") {
		t.Fatalf("output leaks secret or stack: %s", out)
	}
	if !strings.Contains(out, "ORA-01017") || !strings.Contains(out, "usuário ou senha") {
		t.Fatalf("expected humanized message: %s", out)
	}
}

func TestRunHappyPath(t *testing.T) {
	t.Setenv("DOCTOR_TEST_PW", "pw")
	dir := t.TempDir()
	cfg := writeWaves(t, dir, "PCFILIAL", "PCMOV", "PCNOPE")
	if err := os.WriteFile(filepath.Join(dir, "PCFILIAL.yaml"), []byte("table: PCFILIAL\nstatus: approved\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 10, 25, 12, 0, 0, 0, time.UTC)
	mock := &sourcetest.Mock{Handlers: []sourcetest.Handler{
		{Match: "SYS_CONTEXT('USERENV','DB_NAME')", Rows: [][]any{{"MINT_RO", "WINTDB", "srv1", "wint"}}},
		{Match: "PRODUCT_COMPONENT_VERSION", Rows: [][]any{{"Oracle Database 19c Enterprise Edition", "19.0.0.0.0"}}},
		{Match: "NLS_DATABASE_PARAMETERS", Rows: [][]any{{"NLS_CHARACTERSET", "WE8MSWIN1252"}, {"NLS_NCHAR_CHARACTERSET", "AL16UTF16"}, {"NLS_LENGTH_SEMANTICS", "BYTE"}}},
		{Match: "SYS_EXTRACT_UTC", Rows: [][]any{{now.Add(5 * time.Second), "-03:00", "-03:00"}}},
		{Match: "SESSION_PRIVS", Rows: [][]any{{"CREATE SESSION"}}},
		{Match: "USER_ROLE_PRIVS", Rows: [][]any{{"CONNECT"}}},
		{Match: "USER_TAB_PRIVS", Rows: [][]any{{"WINT", "PCFILIAL", "SELECT"}, {"WINT", "PCMOV", "SELECT"}}},
		{Match: "ROLE_TAB_PRIVS", Rows: [][]any{}},
		{Match: "DBMS_FLASHBACK", Err: errors.New("ORA-00904: invalid identifier")},
		{Match: "DBA_EXTENTS", Err: errors.New("ORA-00942: table or view does not exist")},
		{Match: "V$SESSION", Rows: [][]any{{"0"}}},
		{Match: "ALL_TABLES", Rows: [][]any{{"PCFILIAL", "WINT"}, {"PCMOV", "WINT"}}},
		{Match: "ALL_SYNONYMS", Rows: [][]any{{"PCFILIAL", "WINT"}, {"PCMOV", "WINT"}}},
		{Match: `FROM "PCFILIAL" WHERE ROWNUM < 1`, Rows: [][]any{}},
		{Match: `FROM "PCMOV" WHERE ROWNUM < 1`, Rows: [][]any{}},
		{Match: `FROM "PCNOPE" WHERE ROWNUM < 1`, Err: errors.New("ORA-00942: table or view does not exist")},
	}}
	connect := func(context.Context, source.ConnectParams) (source.Source, error) { return mock, nil }
	r := Run(context.Background(), Options{Config: cfg, Connect: connect, Now: func() time.Time { return now }})

	if r.TablesExpected != 3 || r.TablesChecked != 2 || len(r.TablesMissing) != 1 || r.TablesMissing[0] != "PCNOPE" {
		t.Fatalf("tables: expected=%d checked=%d missing=%v", r.TablesExpected, r.TablesChecked, r.TablesMissing)
	}
	if !r.HasFail() {
		t.Fatal("missing table must FAIL")
	}
	find := func(check string) Result {
		for _, res := range r.Results {
			if res.Check == check {
				return res
			}
		}
		t.Fatalf("check %q not found in %+v", check, r.Results)
		return Result{}
	}
	if find("charset").Status != Warn {
		t.Fatalf("charset: %+v", find("charset"))
	}
	if find("relógio").Status != OK {
		t.Fatalf("clock: %+v", find("relógio"))
	}
	if find("privilégios de sistema").Status != OK || find("papéis").Status != OK || find("privilégios de objeto").Status != OK {
		t.Fatal("privileges should be OK")
	}
	if find("arquivos de tabela").Status != Warn {
		t.Fatalf("catalog files: %+v", find("arquivos de tabela"))
	}
	if !mock.Closed {
		t.Fatal("source not closed")
	}
	var buf bytes.Buffer
	if err := WriteJSON(&buf, r); err != nil || !strings.Contains(buf.String(), `"tables_checked": 2`) {
		t.Fatalf("json: %v %s", err, buf.String())
	}
}
