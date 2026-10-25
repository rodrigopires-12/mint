package inventory

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/the12company/mint/internal/catalog"
	"github.com/the12company/mint/internal/source/sourcetest"
)

func mockDictionary() *sourcetest.Mock {
	analyzed := time.Date(2026, 10, 1, 3, 0, 0, 0, time.UTC)
	return &sourcetest.Mock{Handlers: []sourcetest.Handler{
		{Match: "PRODUCT_COMPONENT_VERSION", Rows: [][]any{{"Oracle Database 19c", "19.0.0.0.0"}}},
		{Match: "FROM ALL_SYNONYMS", Rows: [][]any{}},
		{Match: "SELECT owner FROM ALL_TABLES", Rows: [][]any{{"WINT"}}},
		{Match: "FROM ALL_TAB_COLUMNS", Rows: [][]any{
			{"NUMTRANSITEM", "NUMBER", int64(22), int64(12), int64(0), "N", nil, int64(1)},
			{"DTMOV", "DATE", int64(7), nil, nil, "Y", nil, int64(2)},
			{"CODFILIAL", "VARCHAR2", int64(2), nil, nil, "Y", "B", int64(3)},
			{"QT", "NUMBER", int64(22), int64(18), int64(6), "Y", nil, int64(4)},
			{"OBS", "CLOB", int64(4000), nil, nil, "Y", nil, int64(5)},
		}},
		{Match: "FROM ALL_COL_COMMENTS", Rows: [][]any{{"DTMOV", "Data do movimento"}}},
		{Match: "FROM ALL_CONSTRAINTS", Rows: [][]any{{"PK_PCMOV", "P", "NUMTRANSITEM"}}},
		{Match: "FROM ALL_INDEXES", Rows: [][]any{{"PK_PCMOV", "UNIQUE", "NUMTRANSITEM"}, {"I_PCMOV_DTMOV", "NONUNIQUE", "DTMOV"}}},
		{Match: "SELECT num_rows", Rows: [][]any{{int64(1000), int64(128), int64(210), analyzed}}},
		{Match: "SELECT COUNT(*), TO_CHAR(MIN(", Rows: [][]any{{int64(1000), "2024-01-01", "2026-10-01", int64(0)}}},
		{Match: "GROUP BY TRUNC(", Rows: [][]any{{"2026-09", int64(400)}, {"2026-10", int64(600)}}},
		{Match: "HAVING COUNT(*) > 1", Rows: [][]any{{int64(0)}}},
		{Match: `WHERE "NUMTRANSITEM" IS NULL`, Rows: [][]any{{int64(0)}}},
		{Match: `SELECT COUNT(*), COUNT("NUMPED")`, Rows: [][]any{{int64(1000), int64(0)}}},
	}}
}

func testOptions(mock *sourcetest.Mock) Options {
	waves := &catalog.Waves{Version: 1, Waves: []catalog.Wave{{ID: 1, Name: "t", Tables: []string{"PCMOV"}}}}
	scope := &catalog.Scope{Version: 1, Tables: map[string]catalog.ScopeTable{"PCMOV": {
		Description: "Itens", Strategy: catalog.StrategyWindow, Key: []string{"NUMTRANSITEM"},
		Watermark: &catalog.ColumnRef{Column: "DTMOV", Kind: "date"}, Partition: &catalog.ColumnRef{Column: "DTMOV", Kind: "date"}, ReprocessDays: 45,
		Columns: []catalog.ScopeColumn{
			{Name: "NUMTRANSITEM", Use: catalog.UseKey, Ref: "fato_faturamento.num_transacao_item"},
			{Name: "DTMOV", Use: catalog.UseWatermark},
			{Name: "CODFILIAL", Use: catalog.UseScope, Ref: "fato_faturamento.cod_filial"},
			{Name: "QT", Use: catalog.UseProposed},
			{Name: "MISSINGCOL", Use: catalog.UseProposed, Note: "does not exist"},
		},
	}}}
	return Options{Src: mock, Waves: waves, Scope: scope, Count: true, Months: 12, Now: func() time.Time { return time.Date(2026, 10, 25, 12, 0, 0, 0, time.UTC) }}
}

func TestRunBuildsSpec(t *testing.T) {
	res := Run(context.Background(), testOptions(mockDictionary()))
	if len(res.Tables) != 1 || res.Tables[0].Err != nil {
		t.Fatalf("tables: %+v", res.Tables)
	}
	s := res.Tables[0].Spec
	if s.Owner != "WINT" || s.Wave != 1 || s.Status != catalog.StatusProposed || s.Strategy != catalog.StrategyWindow {
		t.Fatalf("spec header: %+v", s)
	}
	if len(s.Columns) != 4 || len(s.MissingColumns) != 1 || s.MissingColumns[0] != "MISSINGCOL" {
		t.Fatalf("columns: %+v missing: %v", s.Columns, s.MissingColumns)
	}
	byName := map[string]catalog.ColumnSpec{}
	for _, c := range s.Columns {
		byName[c.Name] = c
	}
	if byName["CODFILIAL"].ClickHouseType != "String" || byName["QT"].ClickHouseType != "Decimal(18,6)" || byName["NUMTRANSITEM"].ClickHouseType != "Int64" {
		t.Fatalf("types: %+v", byName)
	}
	if byName["DTMOV"].Comment != "Data do movimento" || byName["NUMTRANSITEM"].Nullable {
		t.Fatalf("comment/nullable: %+v", byName)
	}
	if s.PrimaryKey.Source != "constraint:PK_PCMOV" || s.PrimaryKey.VerifiedUnique == nil || !*s.PrimaryKey.VerifiedUnique {
		t.Fatalf("key: %+v", s.PrimaryKey)
	}
	if s.Stats.Count == nil || *s.Stats.Count != 1000 || s.Stats.WatermarkMin != "2024-01-01" || s.Stats.LastAnalyzed != "2026-10-01" {
		t.Fatalf("stats: %+v", s.Stats)
	}
	if len(res.Tables[0].ByMonth) != 2 {
		t.Fatalf("by month: %+v", res.Tables[0].ByMonth)
	}
	joined := strings.Join(s.Risks, " | ")
	if !strings.Contains(joined, "dictionary_has_lob_or_long") || strings.Contains(joined, "watermark_not_indexed") {
		t.Fatalf("risks: %v", s.Risks)
	}
	var f2, f12 bool
	for _, f := range res.Findings {
		switch f.ID {
		case "2":
			f2 = strings.Contains(strings.Join(f.Evidence, "\n"), "PK PK_PCMOV (NUMTRANSITEM)")
		case "12":
			f12 = strings.Contains(strings.Join(f.Evidence, "\n"), "CODFILIAL: VARCHAR2(2) em 1 tabelas (PCMOV)")
		}
	}
	if !f2 || !f12 {
		t.Fatalf("findings: %+v", res.Findings)
	}
}

func TestWriteSpecsAndReport(t *testing.T) {
	res := Run(context.Background(), testOptions(mockDictionary()))
	dir := t.TempDir()
	written, skipped, err := WriteSpecs(dir, res, false)
	if err != nil || len(written) != 1 || len(skipped) != 0 {
		t.Fatalf("write: %v %v %v", written, skipped, err)
	}
	spec, err := catalog.LoadTableSpec(dir, "PCMOV")
	if err != nil || spec.PrimaryKey.Columns[0] != "NUMTRANSITEM" || len(spec.Columns) != 4 {
		t.Fatalf("roundtrip: %+v %v", spec, err)
	}
	// Approved files are preserved.
	spec.Status = catalog.StatusApproved
	if err := catalog.SaveTableSpec(dir, spec); err != nil {
		t.Fatal(err)
	}
	if _, skipped, err = WriteSpecs(dir, res, false); err != nil || len(skipped) != 1 {
		t.Fatalf("approved not preserved: %v %v", skipped, err)
	}
	report := filepath.Join(dir, "docs", "tables-inventory.md")
	if err := WriteReport(report, res, filepath.Join(dir, "docs", "inventory")); err != nil {
		t.Fatal(err)
	}
	if err := WriteDictionaries(filepath.Join(dir, "docs", "inventory"), res); err != nil {
		t.Fatal(err)
	}
	body, _ := os.ReadFile(report)
	for _, want := range []string{"| PCMOV | 1 | WINT | 1000 | 1000 |", "### #2", "### #12", "MISSINGCOL", "2026-09=400"} {
		if !strings.Contains(string(body), want) {
			t.Errorf("report lacks %q", want)
		}
	}
	if _, err := os.Stat(filepath.Join(dir, "docs", "inventory", "PCMOV.md")); err != nil {
		t.Fatal(err)
	}
}

func TestRunTableNotFound(t *testing.T) {
	mock := &sourcetest.Mock{Handlers: []sourcetest.Handler{
		{Match: "PRODUCT_COMPONENT_VERSION", Rows: [][]any{{"Oracle Database 19c", "19.0.0.0.0"}}},
		{Match: "FROM ALL_SYNONYMS", Rows: [][]any{}},
		{Match: "SELECT owner FROM ALL_TABLES", Rows: [][]any{}},
	}}
	res := Run(context.Background(), testOptions(mock))
	if res.Tables[0].Err == nil || len(res.Tables[0].Spec.MissingColumns) != 5 {
		t.Fatalf("expected not-found result: %+v", res.Tables[0])
	}
}
