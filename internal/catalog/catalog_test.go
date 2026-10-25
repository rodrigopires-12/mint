package catalog

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadRepositoryWaves(t *testing.T) {
	w, err := LoadWaves(filepath.Join("..", "..", "sources", "winthor", WavesFile))
	if err != nil {
		t.Fatal(err)
	}
	if got := len(w.AllTables()); got != 46 {
		t.Fatalf("expected 46 tables, got %d", got)
	}
	if len(w.Waves) != 3 {
		t.Fatalf("expected 3 waves, got %d", len(w.Waves))
	}
	if w.WaveOf("PCPREST") != 2 || w.WaveOf("PCFILIAL") != 1 || w.WaveOf("NOPE") != 0 {
		t.Fatal("WaveOf mismatch")
	}
}

func TestLoadWavesRejectsBadInput(t *testing.T) {
	dir := t.TempDir()
	cases := map[string]string{
		"dup":      "version: 1\nwaves:\n  - id: 1\n    name: a\n    tables: [PCA]\n  - id: 2\n    name: b\n    tables: [PCA]\n",
		"lower":    "version: 1\nwaves:\n  - id: 1\n    name: a\n    tables: [pca]\n",
		"inject":   "version: 1\nwaves:\n  - id: 1\n    name: a\n    tables: ['PCA; DROP TABLE X']\n",
		"version":  "version: 2\nwaves:\n  - id: 1\n    name: a\n    tables: [PCA]\n",
		"unknown":  "version: 1\nwavez: []\n",
		"emptywav": "version: 1\nwaves:\n  - id: 1\n    name: a\n    tables: []\n",
	}
	for name, body := range cases {
		p := filepath.Join(dir, name+".yaml")
		if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
		if _, err := LoadWaves(p); err == nil {
			t.Errorf("%s: expected error", name)
		}
	}
}

func TestLoadTableSpec(t *testing.T) {
	dir := t.TempDir()
	if _, err := LoadTableSpec(dir, "PCMOV"); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected ErrNotExist, got %v", err)
	}
	write := func(name, body string) {
		if err := os.WriteFile(filepath.Join(dir, name+".yaml"), []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	write("PCMOV", "table: PCMOV\nstatus: proposed\nwave: 1\ncolumns:\n  - name: NUMTRANSITEM\n")
	spec, err := LoadTableSpec(dir, "PCMOV")
	if err != nil || spec.Status != StatusProposed || spec.Wave != 1 {
		t.Fatalf("spec=%+v err=%v", spec, err)
	}
	write("PCFILIAL", "table: PCOTHER\nstatus: approved\n")
	if _, err := LoadTableSpec(dir, "PCFILIAL"); err == nil {
		t.Fatal("expected table name mismatch error")
	}
	write("PCEST", "status: draft\n")
	if _, err := LoadTableSpec(dir, "PCEST"); err == nil {
		t.Fatal("expected invalid status error")
	}
	if _, err := LoadTableSpec(dir, "../etc/passwd"); err == nil {
		t.Fatal("expected invalid identifier error")
	}
}
