package config

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func write(t *testing.T, dir, name, body string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	return p
}

func TestLoadOverlayAndEnv(t *testing.T) {
	dir := t.TempDir()
	base := write(t, dir, "config.yaml", "tenant_id: piloto\noracle:\n  dsn: db:1521/SVC\n  user: ro\n  password: env:ORACLE_PASSWORD\n  max_sessions: 4\n")
	local := write(t, dir, "config.local.yaml", "oracle:\n  user: local_user\nlogging:\n  level: debug\n")
	empty := write(t, dir, "empty.yaml", "")
	missing := filepath.Join(dir, "nope.yaml")

	cfg, loaded, err := Load([]string{base, local, empty, missing})
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded) != 3 {
		t.Fatalf("loaded = %v", loaded)
	}
	if cfg.TenantID != "piloto" || cfg.Oracle.DSN != "db:1521/SVC" || cfg.Oracle.User != "local_user" {
		t.Fatalf("overlay failed: %+v", cfg)
	}
	if cfg.Oracle.MaxSessions != 4 || cfg.Oracle.ConnectTimeout != 15*time.Second {
		t.Fatalf("defaults/overlay failed: %+v", cfg.Oracle)
	}
	if cfg.Logging.Level != "debug" || cfg.Logging.Format != "text" {
		t.Fatalf("logging: %+v", cfg.Logging)
	}

	env := map[string]string{"ORACLE_PASSWORD": "override", "ORACLE_SCHEMA": "WINT"}
	cfg.ApplyEnv(func(k string) (string, bool) { v, ok := env[k]; return v, ok })
	if cfg.Oracle.Password != "env:ORACLE_PASSWORD" || cfg.Oracle.Schema != "WINT" {
		t.Fatalf("env overlay failed: %+v", cfg.Oracle)
	}
}

func TestUnknownFieldIsRejected(t *testing.T) {
	dir := t.TempDir()
	p := write(t, dir, "config.yaml", "tenant_id: x\noracle:\n  dsnn: typo\n")
	if _, _, err := Load([]string{p}); err == nil {
		t.Fatal("expected error for unknown field")
	}
}

func TestValidateOracle(t *testing.T) {
	cfg := Default()
	if err := cfg.ValidateOracle(); err == nil {
		t.Fatal("expected missing connection fields to fail")
	}
	cfg.Oracle.DSN, cfg.Oracle.User, cfg.Oracle.Password = "h/s", "u", "p"
	if err := cfg.ValidateOracle(); err != nil {
		t.Fatal(err)
	}
}
