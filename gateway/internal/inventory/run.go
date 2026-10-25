package inventory

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/the12company/mint/internal/catalog"
	"github.com/the12company/mint/internal/oradict"
	"github.com/the12company/mint/internal/source"
)

// Options configures an inventory run.
type Options struct {
	Src          source.Source
	Waves        *catalog.Waves
	Scope        *catalog.Scope
	Schema       string   // oracle.schema (optional)
	Tables       []string // subset; empty = all catalog tables
	Count        bool     // run COUNT(*), uniqueness and monthly distribution queries (full scans)
	Months       int      // months of history for the monthly distribution
	QueryTimeout time.Duration
	Now          func() time.Time
	Log          *slog.Logger
}

// MonthCount is one bucket of the watermark distribution.
type MonthCount struct {
	Month string `json:"month"`
	Rows  int64  `json:"rows"`
}

// TableResult is the inventory of one table.
type TableResult struct {
	Spec    *catalog.TableSpec
	Dict    *oradict.Table
	ByMonth []MonthCount
	Err     error
}

// Finding answers one of the dictionary-only questions of docs/escopo-v1.md.
type Finding struct {
	ID       string
	Title    string
	Evidence []string
	Answer   string
}

// Result of a run.
type Result struct {
	InventoriedAt time.Time
	DBVersion     string
	Tables        []*TableResult
	Findings      []Finding
}

// Table returns the result for a table name, or nil.
func (r *Result) Table(name string) *TableResult {
	for _, t := range r.Tables {
		if t.Spec.Table == name {
			return t
		}
	}
	return nil
}

// Run inventories every selected table and computes the findings.
func Run(ctx context.Context, o Options) *Result {
	if o.Now == nil {
		o.Now = time.Now
	}
	if o.Log == nil {
		o.Log = slog.New(slog.DiscardHandler)
	}
	if o.QueryTimeout <= 0 {
		o.QueryTimeout = 10 * time.Minute
	}
	if o.Months <= 0 {
		o.Months = 36
	}
	q := &querier{src: o.Src, timeout: o.QueryTimeout}
	res := &Result{InventoriedAt: o.Now()}
	res.DBVersion = q.dbVersion(ctx)

	selected := o.Waves.AllTables()
	if len(o.Tables) > 0 {
		want := map[string]bool{}
		for _, t := range o.Tables {
			want[strings.ToUpper(t)] = true
		}
		var sub []string
		for _, t := range selected {
			if want[t] {
				sub = append(sub, t)
			}
		}
		selected = sub
	}
	for _, name := range selected {
		o.Log.Info("inventory", "table", name)
		a := &analyzer{o: o, q: q, name: name, seed: o.Scope.Tables[name]}
		res.Tables = append(res.Tables, a.run(ctx))
	}
	res.Findings = findings(ctx, q, res)
	return res
}
