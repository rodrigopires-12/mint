package inventory

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/the12company/mint/internal/catalog"
	"github.com/the12company/mint/internal/oradict"
)

// analyzer builds the TableSpec of one table.
type analyzer struct {
	o    Options
	q    *querier
	name string
	seed catalog.ScopeTable
}

func (a *analyzer) run(ctx context.Context) *TableResult {
	spec := &catalog.TableSpec{
		Table: a.name, Wave: a.o.Waves.WaveOf(a.name), Status: catalog.StatusProposed,
		Description: a.seed.Description, Strategy: a.seed.Strategy,
		Watermark: a.seed.Watermark, Partition: a.seed.Partition, ReprocessDays: a.seed.ReprocessDays,
		OpenWhenNull: a.seed.OpenWhenNull, Questions: append([]string(nil), a.seed.Questions...),
		Stats: &catalog.TableStats{InventoriedAt: a.o.Now().Format("2006-01-02T15:04:05Z07:00")},
	}
	tr := &TableResult{Spec: spec}
	if spec.Strategy == "" {
		spec.Strategy = catalog.StrategyFull
	}

	owner, err := oradict.ResolveOwner(ctx, a.o.Src, a.name, a.o.Schema)
	if err == nil {
		tr.Dict, err = oradict.Describe(ctx, a.o.Src, owner, a.name)
	}
	if err != nil {
		tr.Err = err
		spec.Risks = append(spec.Risks, "not_found: "+err.Error())
		spec.PrimaryKey = catalog.KeySpec{Columns: a.seed.Key, Source: "proposed"}
		for _, c := range a.seed.Columns {
			spec.MissingColumns = append(spec.MissingColumns, c.Name)
		}
		return tr
	}
	dict := tr.Dict
	spec.Owner = owner
	spec.Stats.DictionaryColumns = len(dict.Columns)
	spec.Stats.NumRowsDictionary = dict.Stats.NumRows
	spec.Stats.AvgRowLen = dict.Stats.AvgRowLen
	if dict.Stats.LastAnalyzed != nil {
		spec.Stats.LastAnalyzed = dict.Stats.LastAnalyzed.Format("2006-01-02")
	}
	if dict.Stats.Blocks != nil {
		mb := float64(*dict.Stats.Blocks) * 8192 / (1024 * 1024)
		spec.Stats.EstimatedMB = &mb
	}

	a.selectColumns(spec, dict)
	a.resolveKey(spec, dict)
	a.checkWatermark(spec, dict)
	a.dictionaryRisks(spec, dict)
	if a.o.Count {
		a.measure(ctx, tr)
	}
	return tr
}

// selectColumns maps the seed columns onto the dictionary.
func (a *analyzer) selectColumns(spec *catalog.TableSpec, dict *oradict.Table) {
	for _, sc := range a.seed.Columns {
		col := dict.Column(sc.Name)
		if col == nil {
			spec.MissingColumns = append(spec.MissingColumns, sc.Name)
			if sc.Use == catalog.UseKey {
				spec.Risks = append(spec.Risks, "key_column_missing: "+sc.Name)
			}
			continue
		}
		spec.Columns = append(spec.Columns, a.columnSpec(*col, sc.Use, sc.Ref, sc.Note))
	}
}

func (a *analyzer) columnSpec(col oradict.Column, use, ref, note string) catalog.ColumnSpec {
	chType, risk := ClickHouseType(col)
	cs := catalog.ColumnSpec{Name: col.Name, OracleType: col.TypeString(), ClickHouseType: chType, Nullable: col.Nullable, Use: use, Ref: ref, Note: note, Comment: strings.TrimSpace(col.Comment)}
	if risk != "" {
		cs.Note = strings.Trim(strings.Join([]string{cs.Note, "risk: " + risk}, "; "), "; ")
	}
	return cs
}

// ensureColumn adds a dictionary column to the spec when the seed omitted it.
func (a *analyzer) ensureColumn(spec *catalog.TableSpec, dict *oradict.Table, name, use string) {
	for i := range spec.Columns {
		if strings.EqualFold(spec.Columns[i].Name, name) {
			if use == catalog.UseKey && spec.Columns[i].Use != catalog.UseKey {
				spec.Columns[i].Use = catalog.UseKey
			}
			return
		}
	}
	if col := dict.Column(name); col != nil {
		spec.Columns = append(spec.Columns, a.columnSpec(*col, use, "", "added by inventory: required for "+use))
	}
}

// resolveKey prefers the dictionary's primary key, then a unique constraint or
// index, then the seed proposal.
func (a *analyzer) resolveKey(spec *catalog.TableSpec, dict *oradict.Table) {
	if pk := dict.PrimaryKey(); pk != nil {
		spec.PrimaryKey = catalog.KeySpec{Columns: pk.Columns, Source: "constraint:" + pk.Name}
		if len(a.seed.Key) > 0 && !sameSet(pk.Columns, a.seed.Key) {
			spec.Questions = append(spec.Questions, fmt.Sprintf("A chave primária do dicionário (%s: %s) difere da esperada pelo escopo (%s). O inventário adota a do dicionário.", pk.Name, strings.Join(pk.Columns, ", "), strings.Join(a.seed.Key, ", ")))
		}
	} else if u := bestUnique(dict, a.seed.Key); u != nil {
		spec.PrimaryKey = catalog.KeySpec{Columns: u.Columns, Source: u.Source}
		if len(a.seed.Key) > 0 && !sameSet(u.Columns, a.seed.Key) {
			spec.Questions = append(spec.Questions, fmt.Sprintf("Sem chave primária; a chave única %s (%s) difere da esperada (%s).", u.Source, strings.Join(u.Columns, ", "), strings.Join(a.seed.Key, ", ")))
		}
	} else if len(a.seed.Key) > 0 {
		spec.PrimaryKey = catalog.KeySpec{Columns: a.seed.Key, Source: "proposed"}
		spec.Risks = append(spec.Risks, "no_unique_key_in_dictionary")
	} else {
		spec.PrimaryKey = catalog.KeySpec{Source: "none"}
		spec.Risks = append(spec.Risks, "no_key")
	}
	for _, k := range spec.PrimaryKey.Columns {
		a.ensureColumn(spec, dict, k, catalog.UseKey)
		if c := dict.Column(k); c != nil && c.Nullable {
			spec.Risks = append(spec.Risks, "nullable_key_column: "+k)
		}
	}
}

type uniqueKey struct {
	Columns []string
	Source  string
}

// bestUnique picks the unique constraint or index matching the expected key,
// else the shortest one.
func bestUnique(dict *oradict.Table, expected []string) *uniqueKey {
	var cands []uniqueKey
	for _, c := range dict.UniqueConstraints() {
		cands = append(cands, uniqueKey{c.Columns, "constraint:" + c.Name})
	}
	for _, i := range dict.UniqueIndexes() {
		cands = append(cands, uniqueKey{i.Columns, "unique_index:" + i.Name})
	}
	if len(cands) == 0 {
		return nil
	}
	for i := range cands {
		if sameSet(cands[i].Columns, expected) {
			return &cands[i]
		}
	}
	sort.SliceStable(cands, func(i, j int) bool { return len(cands[i].Columns) < len(cands[j].Columns) })
	return &cands[0]
}

func sameSet(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := map[string]bool{}
	for _, x := range a {
		m[strings.ToUpper(x)] = true
	}
	for _, y := range b {
		if !m[strings.ToUpper(y)] {
			return false
		}
	}
	return true
}

// checkWatermark verifies watermark and partition columns exist and are indexed.
func (a *analyzer) checkWatermark(spec *catalog.TableSpec, dict *oradict.Table) {
	refs := []struct {
		label string
		ref   *catalog.ColumnRef
		use   string
	}{{"watermark", spec.Watermark, catalog.UseWatermark}, {"partition", spec.Partition, catalog.UseProposed}}
	for _, r := range refs {
		if r.ref == nil {
			continue
		}
		col := dict.Column(r.ref.Column)
		if col == nil {
			spec.Risks = append(spec.Risks, r.label+"_column_missing: "+r.ref.Column)
			spec.MissingColumns = appendUnique(spec.MissingColumns, r.ref.Column)
			continue
		}
		a.ensureColumn(spec, dict, r.ref.Column, r.use)
		if r.ref.Kind == "date" && col.DataType != "DATE" && !strings.HasPrefix(col.DataType, "TIMESTAMP") {
			spec.Risks = append(spec.Risks, fmt.Sprintf("%s_not_a_date: %s is %s", r.label, r.ref.Column, col.TypeString()))
		}
		if dict.IndexLeading(r.ref.Column) == nil {
			spec.Risks = append(spec.Risks, r.label+"_not_indexed: "+r.ref.Column)
		}
	}
	if spec.Strategy != catalog.StrategyFull && spec.Watermark == nil {
		spec.Risks = append(spec.Risks, "incremental_without_watermark")
	}
}

// dictionaryRisks flags LOB/LONG columns and very wide tables.
func (a *analyzer) dictionaryRisks(spec *catalog.TableSpec, dict *oradict.Table) {
	var lobs []string
	for _, c := range dict.Columns {
		if IsLOB(c) {
			lobs = append(lobs, c.Name+" "+c.DataType)
		}
	}
	for _, c := range spec.Columns {
		if strings.Contains(c.Note, "risk: "+RiskLOB) {
			spec.Risks = append(spec.Risks, "selected_lob_column: "+c.Name)
		}
	}
	if len(lobs) > 0 {
		spec.Risks = append(spec.Risks, fmt.Sprintf("dictionary_has_lob_or_long (%d, not extracted): %s", len(lobs), strings.Join(lobs, ", ")))
	}
	if len(dict.Columns) > 200 {
		spec.Risks = append(spec.Risks, fmt.Sprintf("wide_table: %d columns in dictionary", len(dict.Columns)))
	}
}

func appendUnique(list []string, v string) []string {
	for _, x := range list {
		if x == v {
			return list
		}
	}
	return append(list, v)
}
