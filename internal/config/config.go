// Package config loads the gateway/worker configuration from YAML files and
// environment variables. Later files overlay earlier ones; environment
// variables overlay everything.
package config

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/the12company/mint/internal/secrets"
)

// Config is the root configuration.
type Config struct {
	TenantID   string  `yaml:"tenant_id"`
	Oracle     Oracle  `yaml:"oracle"`
	SourcesDir string  `yaml:"sources_dir"`
	Logging    Logging `yaml:"logging"`
}

// Oracle holds the connection to the customer's Winthor database.
type Oracle struct {
	DSN            string        `yaml:"dsn"` // EZCONNECT: host[:port]/service
	User           string        `yaml:"user"`
	Password       string        `yaml:"password"` // literal, env:NAME or dpapi:<ciphertext>
	Schema         string        `yaml:"schema"`   // owner of PC* tables; empty = resolve
	ConnectTimeout time.Duration `yaml:"connect_timeout"`
	QueryTimeout   time.Duration `yaml:"query_timeout"`
	MaxSessions    int           `yaml:"max_sessions"`
}

// Logging configures the structured logger.
type Logging struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	File   string `yaml:"file"`
}

// DefaultFiles are loaded, in order, when no --config flag is given.
var DefaultFiles = []string{"config.yaml", "config.local.yaml"}

// Default returns the built-in defaults.
func Default() Config {
	return Config{
		TenantID:   "dev",
		SourcesDir: "sources/winthor",
		Oracle: Oracle{
			ConnectTimeout: 15 * time.Second,
			QueryTimeout:   5 * time.Minute,
			MaxSessions:    2,
		},
		Logging: Logging{Level: "info", Format: "text"},
	}
}

// Load reads the given files (missing files are skipped), overlays them in
// order on top of the defaults, applies environment overrides and validates.
// It returns the configuration and the list of files actually read.
func Load(paths []string) (Config, []string, error) {
	cfg := Default()
	var loaded []string
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return cfg, loaded, fmt.Errorf("read %s: %w", p, err)
		}
		dec := yaml.NewDecoder(bytes.NewReader(data))
		dec.KnownFields(true)
		if err := dec.Decode(&cfg); err != nil && !errors.Is(err, io.EOF) {
			return cfg, loaded, fmt.Errorf("parse %s: %w", p, err)
		}
		loaded = append(loaded, p)
	}
	cfg.ApplyEnv(os.LookupEnv)
	return cfg, loaded, cfg.Validate()
}

// ApplyEnv overlays well-known environment variables.
func (c *Config) ApplyEnv(lookup func(string) (string, bool)) {
	set := func(dst *string, name string) {
		if v, ok := lookup(name); ok && v != "" {
			*dst = v
		}
	}
	set(&c.Oracle.DSN, "ORACLE_DSN")
	set(&c.Oracle.User, "ORACLE_USER")
	// ORACLE_PASSWORD always wins over the file so operators can override a
	// stored secret without editing YAML (the verification scripts rely on it).
	// Only the reference is stored; the value is resolved when connecting.
	if v, ok := lookup("ORACLE_PASSWORD"); ok && v != "" {
		c.Oracle.Password = "env:ORACLE_PASSWORD"
	}
	set(&c.Oracle.Schema, "ORACLE_SCHEMA")
	set(&c.TenantID, "MINT_TENANT_ID")
	set(&c.Logging.Level, "MINT_LOG_LEVEL")
	set(&c.Logging.Format, "MINT_LOG_FORMAT")
	set(&c.SourcesDir, "MINT_SOURCES_DIR")
}

// Validate checks the parts that every command needs.
func (c Config) Validate() error {
	var errs []error
	if strings.TrimSpace(c.TenantID) == "" {
		errs = append(errs, errors.New("tenant_id is required"))
	}
	if c.Oracle.ConnectTimeout <= 0 {
		errs = append(errs, errors.New("oracle.connect_timeout must be positive"))
	}
	if c.Oracle.QueryTimeout <= 0 {
		errs = append(errs, errors.New("oracle.query_timeout must be positive"))
	}
	if c.Oracle.MaxSessions <= 0 {
		errs = append(errs, errors.New("oracle.max_sessions must be at least 1"))
	}
	return errors.Join(errs...)
}

// ValidateOracle checks that a connection can be attempted.
func (c Config) ValidateOracle() error {
	var errs []error
	if strings.TrimSpace(c.Oracle.DSN) == "" {
		errs = append(errs, errors.New("oracle.dsn is required (host[:port]/service, or ORACLE_DSN)"))
	}
	if strings.TrimSpace(c.Oracle.User) == "" {
		errs = append(errs, errors.New("oracle.user is required (or ORACLE_USER)"))
	}
	if strings.TrimSpace(c.Oracle.Password) == "" {
		errs = append(errs, errors.New("oracle.password is required (env:ORACLE_PASSWORD recommended)"))
	}
	return errors.Join(errs...)
}

// ResolvedPassword resolves the password reference (env:, dpapi: or literal).
func (o Oracle) ResolvedPassword() (string, secrets.Kind, error) {
	return secrets.Resolve(o.Password)
}
