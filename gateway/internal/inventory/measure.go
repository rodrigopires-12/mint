package inventory

import (
	"context"
	"fmt"
	"strings"

	"github.com/the12company/mint/internal/catalog"
	"github.com/the12company/mint/internal/oradict"
)

// measure runs the full-scan queries: COUNT(*), key uniqueness and the
// watermark distribution. Only aggregates are read, never data rows.
func (a *analyzer) measure(ctx context.Context, tr *TableResult) {
	spec, dict := tr.Spec, tr.Dict
	qualified, err := oradict.Qualified(dict.Owner, dict.Name)
	if err != nil {
		spec.Risks = append(spec.Risks, "measure_failed: "+err.Error())
		return
	}

	wm := spec.Watermark
	if wm != nil && dict.Column(wm.Column) != nil {
		format := ""
		if wm.Kind == "date" {
			format = "YYYY-MM-DD"
		}
		rs, err := a.q.rangeStats(ctx, qualified, wm.Column, format)
		if err != nil {
			spec.Risks = append(spec.Risks, "measure_failed: range: "+err.Error())
		} else {
			spec.Stats.Count = &rs.Count
			spec.Stats.WatermarkMin, spec.Stats.WatermarkMax = rs.Min, rs.Max
			nulls := rs.Nulls
			spec.Stats.WatermarkNulls = &nulls
			if nulls > 0 {
				spec.Risks = append(spec.Risks, fmt.Sprintf("watermark_has_nulls: %d rows", nulls))
			}
		}
		if wm.Kind == "date" {
			tr.ByMonth, err = a.q.byMonth(ctx, qualified, wm.Column, a.o.Months)
			if err != nil {
				spec.Risks = append(spec.Risks, "measure_failed: by_month: "+err.Error())
			}
		}
	} else {
		n, err := a.q.count(ctx, qualified)
		if err != nil {
			spec.Risks = append(spec.Risks, "measure_failed: count: "+err.Error())
		} else {
			spec.Stats.Count = &n
		}
	}

	for _, m := range tr.ByMonth {
		spec.Stats.ByMonth = append(spec.Stats.ByMonth, catalog.MonthRows{Month: m.Month, Rows: m.Rows})
	}

	if len(spec.PrimaryKey.Columns) > 0 {
		dups, nullRows, err := a.q.duplicateGroups(ctx, qualified, spec.PrimaryKey.Columns)
		if err != nil {
			spec.Risks = append(spec.Risks, "measure_failed: uniqueness: "+err.Error())
		} else {
			unique := dups == 0
			spec.PrimaryKey.VerifiedUnique = &unique
			spec.PrimaryKey.DuplicateGroups = &dups
			spec.PrimaryKey.NullKeyRows = &nullRows
			if !unique {
				spec.Risks = append(spec.Risks, fmt.Sprintf("key_not_unique: %d duplicate groups", dups))
			}
			if nullRows > 0 {
				spec.Risks = append(spec.Risks, fmt.Sprintf("key_has_nulls: %d rows with NULL in a key column", nullRows))
			}
			if !unique || nullRows > 0 {
				a.tryAlternateKeys(ctx, qualified, tr)
			}
		}
	}
}

// tryAlternateKeys tests the seed's alternate keys and records the outcome.
func (a *analyzer) tryAlternateKeys(ctx context.Context, qualified string, tr *TableResult) {
	for _, alt := range a.seed.AlternateKeys {
		missing := false
		for _, c := range alt {
			if tr.Dict.Column(c) == nil {
				missing = true
			}
		}
		if missing {
			continue
		}
		dups, nullRows, err := a.q.duplicateGroups(ctx, qualified, alt)
		if err != nil {
			tr.Spec.Risks = append(tr.Spec.Risks, "measure_failed: alternate key: "+err.Error())
			continue
		}
		tr.Spec.Questions = append(tr.Spec.Questions, fmt.Sprintf("Chave alternativa (%s): %d grupos duplicados, %d linhas com nulo na chave.", strings.Join(alt, ", "), dups, nullRows))
	}
}
