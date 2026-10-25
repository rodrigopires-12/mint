package catalog

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Load strategies (docs/escopo-v1.md, "Carga: janela e tabelas de origem").
const (
	StrategyFull      = "full"       // small tables: reload everything
	StrategyWatermark = "watermark"  // rows with watermark > last value
	StrategyWindow    = "window"     // watermark plus a moving reprocess window
	StrategyOpenItems = "open_items" // window plus every open item regardless of date
)

// Column uses inside a TableSpec.
const (
	UseKey       = "key"
	UseScope     = "scope"
	UseProposed  = "proposed"
	UseWatermark = "watermark"
	UseControl   = "control"
)

var validStrategies = map[string]bool{StrategyFull: true, StrategyWatermark: true, StrategyWindow: true, StrategyOpenItems: true}
var validUses = map[string]bool{UseKey: true, UseScope: true, UseProposed: true, UseWatermark: true, UseControl: true}
var validKinds = map[string]bool{"date": true, "number": true, "rowid": true}

// ColumnRef names a column used for watermark or partitioning.
type ColumnRef struct {
	Column string `yaml:"column"`
	Kind   string `yaml:"kind"` // date | number | rowid
}

// TableSpec is the per-table contract (sources/winthor/<TABLE>.yaml).
type TableSpec struct {
	Table          string       `yaml:"table"`
	Wave           int          `yaml:"wave"`
	Status         string       `yaml:"status"`
	Description    string       `yaml:"description,omitempty"`
	Owner          string       `yaml:"owner,omitempty"`
	Strategy       string       `yaml:"strategy,omitempty"`
	Watermark      *ColumnRef   `yaml:"watermark,omitempty"`
	Partition      *ColumnRef   `yaml:"partition,omitempty"`
	ReprocessDays  int          `yaml:"reprocess_days,omitempty"`
	OpenWhenNull   []string     `yaml:"open_when_null,omitempty"`
	PrimaryKey     KeySpec      `yaml:"primary_key"`
	Columns        []ColumnSpec `yaml:"columns"`
	Stats          *TableStats  `yaml:"stats,omitempty"`
	Risks          []string     `yaml:"risks,omitempty"`
	MissingColumns []string     `yaml:"missing_columns,omitempty"`
	Questions      []string     `yaml:"questions,omitempty"`
}

// KeySpec is the row identity used for dedup, deletes and reconciliation.
type KeySpec struct {
	Columns         []string `yaml:"columns"`
	Source          string   `yaml:"source,omitempty"` // constraint:<name> | unique_index:<name> | proposed | none
	VerifiedUnique  *bool    `yaml:"verified_unique,omitempty"`
	DuplicateGroups *int64   `yaml:"duplicate_groups,omitempty"`
	NullKeyRows     *int64   `yaml:"null_key_rows,omitempty"`
}

// ColumnSpec is one extracted column.
type ColumnSpec struct {
	Name           string `yaml:"name"`
	OracleType     string `yaml:"oracle_type,omitempty"`
	ClickHouseType string `yaml:"clickhouse_type,omitempty"`
	Nullable       bool   `yaml:"nullable"`
	Use            string `yaml:"use,omitempty"`
	Ref            string `yaml:"ref,omitempty"`
	Note           string `yaml:"note,omitempty"`
	Comment        string `yaml:"comment,omitempty"`
}

// TableStats are dictionary and count figures gathered by the inventory.
type TableStats struct {
	InventoriedAt     string      `yaml:"inventoried_at"`
	DictionaryColumns int         `yaml:"dictionary_columns"`
	NumRowsDictionary *int64      `yaml:"num_rows_dictionary,omitempty"`
	LastAnalyzed      string      `yaml:"last_analyzed,omitempty"`
	AvgRowLen         *int64      `yaml:"avg_row_len,omitempty"`
	EstimatedMB       *float64    `yaml:"estimated_mb,omitempty"`
	Count             *int64      `yaml:"count,omitempty"`
	WatermarkMin      string      `yaml:"watermark_min,omitempty"`
	WatermarkMax      string      `yaml:"watermark_max,omitempty"`
	WatermarkNulls    *int64      `yaml:"watermark_nulls,omitempty"`
	ByMonth           []MonthRows `yaml:"by_month,omitempty"`
}

// MonthRows is one bucket of the watermark distribution.
type MonthRows struct {
	Month string `yaml:"month"`
	Rows  int64  `yaml:"rows"`
}

// Validate checks identifiers and enumerations. Gaps (no key, no watermark)
// are allowed while the spec is proposed; they are reported as risks.
func (s *TableSpec) Validate() error {
	if !ValidIdentifier(s.Table) {
		return fmt.Errorf("invalid table name %q", s.Table)
	}
	if s.Status != StatusProposed && s.Status != StatusApproved {
		return fmt.Errorf("%s: status must be %q or %q, got %q", s.Table, StatusProposed, StatusApproved, s.Status)
	}
	if s.Strategy != "" && !validStrategies[s.Strategy] {
		return fmt.Errorf("%s: unknown strategy %q", s.Table, s.Strategy)
	}
	for _, ref := range []*ColumnRef{s.Watermark, s.Partition} {
		if ref != nil && (!ValidIdentifier(ref.Column) || !validKinds[ref.Kind]) {
			return fmt.Errorf("%s: invalid column reference %+v", s.Table, *ref)
		}
	}
	for _, c := range s.Columns {
		if !ValidIdentifier(c.Name) {
			return fmt.Errorf("%s: invalid column name %q", s.Table, c.Name)
		}
		if c.Use != "" && !validUses[c.Use] {
			return fmt.Errorf("%s.%s: unknown use %q", s.Table, c.Name, c.Use)
		}
	}
	return nil
}

// SaveTableSpec writes <dir>/<TABLE>.yaml.
func SaveTableSpec(dir string, spec *TableSpec) error {
	if err := spec.Validate(); err != nil {
		return err
	}
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "# %s — generated by `gateway inventory`; review and set status: approved.\n", spec.Table)
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(spec); err != nil {
		return err
	}
	if err := enc.Close(); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, spec.Table+".yaml"), buf.Bytes(), 0o644)
}

// ScopeFile is the hand-written seed consumed by the inventory.
const ScopeFile = "_scope.yaml"

// Scope lists, per table, what docs/escopo-v1.md requires.
type Scope struct {
	Version int                   `yaml:"version"`
	Tables  map[string]ScopeTable `yaml:"tables"`
}

// ScopeTable is the seed for one table.
type ScopeTable struct {
	Description   string        `yaml:"description"`
	Strategy      string        `yaml:"strategy"`
	Key           []string      `yaml:"key"`
	AlternateKeys [][]string    `yaml:"alternate_keys"` // tested for uniqueness when the key fails
	Watermark     *ColumnRef    `yaml:"watermark"`
	Partition     *ColumnRef    `yaml:"partition"`
	ReprocessDays int           `yaml:"reprocess_days"`
	OpenWhenNull  []string      `yaml:"open_when_null"`
	Columns       []ScopeColumn `yaml:"columns"`
	Questions     []string      `yaml:"questions"`
}

// ScopeColumn is one required or proposed column.
type ScopeColumn struct {
	Name string `yaml:"name"`
	Use  string `yaml:"use"`
	Ref  string `yaml:"ref"`
	Note string `yaml:"note"`
}

// LoadScope reads and validates _scope.yaml.
func LoadScope(path string) (*Scope, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	var s Scope
	if err := dec.Decode(&s); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	if s.Version != 1 {
		return nil, fmt.Errorf("%s: unsupported version %d", path, s.Version)
	}
	for name, t := range s.Tables {
		if !ValidIdentifier(name) {
			return nil, fmt.Errorf("%s: invalid table name %q", path, name)
		}
		if !validStrategies[t.Strategy] {
			return nil, fmt.Errorf("%s: %s: unknown strategy %q", path, name, t.Strategy)
		}
		for _, k := range t.Key {
			if !ValidIdentifier(k) {
				return nil, fmt.Errorf("%s: %s: invalid key column %q", path, name, k)
			}
		}
		for _, ref := range []*ColumnRef{t.Watermark, t.Partition} {
			if ref != nil && (!ValidIdentifier(ref.Column) || !validKinds[ref.Kind]) {
				return nil, fmt.Errorf("%s: %s: invalid column reference %+v", path, name, *ref)
			}
		}
		for _, c := range t.Columns {
			if !ValidIdentifier(c.Name) || !validUses[c.Use] {
				return nil, fmt.Errorf("%s: %s: invalid column %+v", path, name, c)
			}
		}
	}
	return &s, nil
}
