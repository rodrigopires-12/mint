// Package catalog reads sources/winthor/_waves.yaml (tables per wave) and the
// per-table specification files sources/winthor/<TABLE>.yaml.
package catalog

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// Waves is the content of _waves.yaml.
type Waves struct {
	Version int    `yaml:"version"`
	Waves   []Wave `yaml:"waves"`
}

// Wave groups the tables of one delivery wave.
type Wave struct {
	ID     int      `yaml:"id"`
	Name   string   `yaml:"name"`
	Tables []string `yaml:"tables"`
}

// Table statuses of a <TABLE>.yaml file.
const (
	StatusProposed = "proposed"
	StatusApproved = "approved"
)

// WavesFile is the file name inside the sources directory.
const WavesFile = "_waves.yaml"

var identRe = regexp.MustCompile(`^[A-Z][A-Z0-9_$#]{0,127}$`)

// ValidIdentifier reports whether s is a safe, upper-case Oracle identifier.
// Table names from the catalog are interpolated into SQL text, so this is a
// security boundary, not just a style check.
func ValidIdentifier(s string) bool { return identRe.MatchString(s) }

// LoadWaves reads and validates _waves.yaml.
func LoadWaves(path string) (*Waves, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	var w Waves
	if err := dec.Decode(&w); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	if err := w.validate(); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return &w, nil
}

func (w *Waves) validate() error {
	if w.Version != 1 {
		return fmt.Errorf("unsupported version %d (expected 1)", w.Version)
	}
	if len(w.Waves) == 0 {
		return errors.New("no waves defined")
	}
	seenID := map[int]bool{}
	seenTable := map[string]int{}
	for _, wave := range w.Waves {
		if wave.ID <= 0 || seenID[wave.ID] {
			return fmt.Errorf("wave id %d is missing or duplicated", wave.ID)
		}
		seenID[wave.ID] = true
		if len(wave.Tables) == 0 {
			return fmt.Errorf("wave %d has no tables", wave.ID)
		}
		for _, t := range wave.Tables {
			if !ValidIdentifier(t) {
				return fmt.Errorf("wave %d: invalid table name %q (upper-case Oracle identifier expected)", wave.ID, t)
			}
			if prev, dup := seenTable[t]; dup {
				return fmt.Errorf("table %s appears in waves %d and %d", t, prev, wave.ID)
			}
			seenTable[t] = wave.ID
		}
	}
	return nil
}

// AllTables returns every table in wave order.
func (w *Waves) AllTables() []string {
	var out []string
	for _, wave := range w.Waves {
		out = append(out, wave.Tables...)
	}
	return out
}

// WaveOf returns the wave id of a table, or 0.
func (w *Waves) WaveOf(table string) int {
	for _, wave := range w.Waves {
		for _, t := range wave.Tables {
			if t == table {
				return wave.ID
			}
		}
	}
	return 0
}

// LoadTableSpec reads <dir>/<TABLE>.yaml. A missing file returns an error
// wrapping os.ErrNotExist.
func LoadTableSpec(dir, table string) (*TableSpec, error) {
	if !ValidIdentifier(table) {
		return nil, fmt.Errorf("invalid table name %q", table)
	}
	path := filepath.Join(dir, table+".yaml")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var spec TableSpec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	if spec.Table == "" {
		spec.Table = table
	}
	if !strings.EqualFold(spec.Table, table) {
		return nil, fmt.Errorf("%s: field table is %q, expected %q", path, spec.Table, table)
	}
	if err := spec.Validate(); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return &spec, nil
}
