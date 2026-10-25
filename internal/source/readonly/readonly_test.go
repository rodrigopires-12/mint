package readonly

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckAccepts(t *testing.T) {
	for _, stmt := range []string{
		"SELECT 1 FROM DUAL",
		"  select codigo from pcfilial where rownum < 10 ",
		"WITH t AS (SELECT 1 n FROM DUAL) SELECT n FROM t",
		"SELECT 'a;b' FROM DUAL",
		"SELECT 'it''s; fine' FROM DUAL",
		"SELECT \"WEIRD;COL\" FROM T",
		"-- UPDATE in a comment\nSELECT 1 FROM DUAL",
		"/* DELETE */ SELECT 1 FROM DUAL -- trailing; comment",
		"SELECT 'for update' FROM DUAL",
		"SELECT * FROM PCMOV AS OF SCN 12345 WHERE ROWNUM < 1",
	} {
		if err := Check(stmt); err != nil {
			t.Errorf("%q: unexpected refusal: %v", stmt, err)
		}
	}
}

func TestCheckRefuses(t *testing.T) {
	for _, stmt := range []string{
		"",
		"   ",
		"-- only a comment",
		"UPDATE PCFILIAL SET X = 1",
		"delete from pcmov",
		"INSERT INTO T VALUES (1)",
		"MERGE INTO T USING DUAL ON (1=0) WHEN NOT MATCHED THEN INSERT (A) VALUES (1)",
		"CREATE TABLE T (A NUMBER)",
		"DROP TABLE T",
		"TRUNCATE TABLE T",
		"ALTER SESSION SET NLS_DATE_FORMAT = 'YYYY'",
		"BEGIN NULL; END;",
		"DECLARE x NUMBER; BEGIN NULL; END;",
		"CALL p()",
		"EXEC p",
		"GRANT SELECT ON T TO U",
		"SELECT 1 FROM DUAL; SELECT 2 FROM DUAL",
		"SELECT 1 FROM DUAL;",
		"SELECT * FROM T FOR UPDATE",
		"SELECT * FROM T FOR   UPDATE NOWAIT",
		"WITH FUNCTION f RETURN NUMBER IS BEGIN RETURN 1; END; SELECT f FROM DUAL",
		"(SELECT 1 FROM DUAL)",
	} {
		err := Check(stmt)
		if err == nil {
			t.Errorf("%q: expected refusal", stmt)
			continue
		}
		if !errors.Is(err, ErrNotReadOnly) {
			t.Errorf("%q: error does not wrap ErrNotReadOnly: %v", stmt, err)
		}
	}
}

// The SQL files used by `gateway sqlcheck` and the verification script must
// agree with the guard.
func TestVerificationFiles(t *testing.T) {
	root := filepath.Join("..", "..", "..", "scripts", "verify", "testdata")
	check := func(name string, wantOK bool) {
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		n := 0
		for _, stmt := range SplitStatements(string(data)) {
			n++
			err := Check(stmt)
			if wantOK && err != nil {
				t.Errorf("%s: %q refused: %v", name, stmt, err)
			}
			if !wantOK && err == nil {
				t.Errorf("%s: %q accepted", name, stmt)
			}
		}
		if n == 0 {
			t.Errorf("%s: no statements found", name)
		}
	}
	check("forbidden.sql", false)
	check("allowed.sql", true)
}

func TestSplitStatements(t *testing.T) {
	got := SplitStatements("-- c; not a split\nSELECT 1 FROM DUAL;\n\nSELECT 'a;b' FROM DUAL;\nBEGIN NULL; END;")
	want := []string{"-- c; not a split\nSELECT 1 FROM DUAL", "SELECT 'a;b' FROM DUAL", "BEGIN NULL", "END"}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestMaskKeepsLength(t *testing.T) {
	in := "SELECT 'ação' /* çomment */ FROM T -- fim"
	if len([]rune(mask(in))) != len([]rune(in)) {
		t.Fatalf("mask changed rune length")
	}
}
