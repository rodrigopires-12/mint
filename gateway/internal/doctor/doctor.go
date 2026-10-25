// Package doctor diagnoses the connection to a customer's Oracle: identity,
// charset, clock, privileges (read-only or not), optional capabilities and
// access to every table of the catalog. It never reads data rows.
package doctor

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"time"

	"github.com/the12company/mint/internal/catalog"
	"github.com/the12company/mint/internal/config"
	"github.com/the12company/mint/internal/logging"
	"github.com/the12company/mint/internal/secrets"
	"github.com/the12company/mint/internal/source"
	"github.com/the12company/mint/internal/source/oraerr"
)

// Status of one check.
type Status string

const (
	OK   Status = "OK"
	Warn Status = "WARN"
	Fail Status = "FAIL"
	Info Status = "INFO"
)

// Result of one check.
type Result struct {
	Check  string `json:"check"`
	Status Status `json:"status"`
	Detail string `json:"detail"`
}

// Report is the outcome of a doctor run.
type Report struct {
	StartedAt      time.Time      `json:"started_at"`
	Results        []Result       `json:"results"`
	TablesExpected int            `json:"tables_expected"`
	TablesChecked  int            `json:"tables_checked"`
	TablesMissing  []string       `json:"tables_missing"`
	Summary        map[Status]int `json:"summary"`

	redact []string
}

// Options configures a run.
type Options struct {
	Config  config.Config
	Connect source.Connector
	Now     func() time.Time
	Logger  *slog.Logger
}

func (r *Report) add(check string, st Status, format string, args ...any) {
	detail := logging.Scrub(fmt.Sprintf(format, args...), r.redact)
	r.Results = append(r.Results, Result{Check: check, Status: st, Detail: detail})
	r.Summary[st]++
}

// HasFail reports whether any check failed.
func (r *Report) HasFail() bool { return r.Summary[Fail] > 0 }

// Run executes every check and returns the report. It stops early when the
// catalog cannot be read or the connection fails, since nothing else can be
// diagnosed then.
func Run(ctx context.Context, o Options) *Report {
	if o.Now == nil {
		o.Now = time.Now
	}
	if o.Logger == nil {
		o.Logger = slog.New(slog.DiscardHandler)
	}
	r := &Report{StartedAt: o.Now(), Summary: map[Status]int{}}
	cfg := o.Config

	waves, err := catalog.LoadWaves(filepath.Join(cfg.SourcesDir, catalog.WavesFile))
	if err != nil {
		r.add("catálogo", Fail, "não foi possível ler %s: %v", catalog.WavesFile, err)
		return r
	}
	tables := waves.AllTables()
	r.TablesExpected = len(tables)
	r.add("catálogo", Info, "%d tabelas em %d ondas (%s)", len(tables), len(waves.Waves), filepath.Join(cfg.SourcesDir, catalog.WavesFile))

	if err := cfg.ValidateOracle(); err != nil {
		r.add("configuração", Fail, "%v", err)
		return r
	}
	password, kind, err := cfg.Oracle.ResolvedPassword()
	if err != nil {
		r.add("configuração", Fail, "senha: %v", err)
		return r
	}
	r.redact = []string{password}
	r.add("configuração", Info, "tenant=%s dsn=%s user=%s schema=%q senha=%s", cfg.TenantID, cfg.Oracle.DSN, cfg.Oracle.User, cfg.Oracle.Schema, kind)
	if kind == secrets.Literal {
		r.add("configuração", Warn, "a senha está em texto puro na configuração; use env:NOME_DA_VARIAVEL ou dpapi: (E1.3)")
	}

	connectCtx, cancel := context.WithTimeout(ctx, cfg.Oracle.ConnectTimeout)
	defer cancel()
	src, err := o.Connect(connectCtx, source.ConnectParams{
		DSN: cfg.Oracle.DSN, User: cfg.Oracle.User, Password: password,
		ConnectTimeout: cfg.Oracle.ConnectTimeout, MaxSessions: cfg.Oracle.MaxSessions,
	})
	if err == nil {
		err = src.Ping(connectCtx)
	}
	if err != nil {
		r.add("conexão", Fail, "%s", oraerr.Humanize(err))
		o.Logger.Debug("connection failed", "error", logging.Scrub(err.Error(), r.redact))
		if src != nil {
			_ = src.Close()
		}
		return r
	}
	defer src.Close()
	r.add("conexão", OK, "conectado a %s em %s", cfg.Oracle.DSN, o.Now().Sub(r.StartedAt).Round(time.Millisecond))

	c := &checker{src: src, r: r, cfg: cfg, now: o.Now, log: o.Logger}
	c.identity(ctx)
	version := c.version(ctx)
	c.charset(ctx)
	c.clock(ctx)
	sysPrivs := c.privileges(ctx)
	c.capabilities(ctx, sysPrivs, version)
	c.tables(ctx, tables)
	c.catalogFiles(tables)
	return r
}
