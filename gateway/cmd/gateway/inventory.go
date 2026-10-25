package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/the12company/mint/gateway/internal/inventory"
	"github.com/the12company/mint/internal/catalog"
	"github.com/the12company/mint/internal/config"
	"github.com/the12company/mint/internal/logging"
	"github.com/the12company/mint/internal/source"
	"github.com/the12company/mint/internal/source/goora"
	"github.com/the12company/mint/internal/source/oraerr"
)

func newInventoryCmd(loadConfig func() (config.Config, []string, error)) *cobra.Command {
	var tables []string
	var noCount, force bool
	var months int
	var report, dictDir string
	cmd := &cobra.Command{
		Use:   "inventory",
		Short: "Inventory the catalog tables from the Oracle dictionary and write <TABLE>.yaml proposals (E1.2)",
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, _, err := loadConfig()
			if err != nil {
				return &exitError{2, "configuração inválida: " + err.Error()}
			}
			if err := cfg.ValidateOracle(); err != nil {
				return &exitError{2, "configuração inválida: " + err.Error()}
			}
			log := logging.New(cmd.ErrOrStderr(), logging.Options{Level: cfg.Logging.Level, Format: cfg.Logging.Format})
			waves, err := catalog.LoadWaves(filepath.Join(cfg.SourcesDir, catalog.WavesFile))
			if err != nil {
				return &exitError{2, err.Error()}
			}
			scope, err := catalog.LoadScope(filepath.Join(cfg.SourcesDir, catalog.ScopeFile))
			if err != nil {
				return &exitError{2, err.Error()}
			}
			password, _, err := cfg.Oracle.ResolvedPassword()
			if err != nil {
				return &exitError{2, "senha: " + err.Error()}
			}
			ctx := cmd.Context()
			src, err := goora.Connect(ctx, source.ConnectParams{DSN: cfg.Oracle.DSN, User: cfg.Oracle.User, Password: password, ConnectTimeout: cfg.Oracle.ConnectTimeout, MaxSessions: 1})
			if err == nil {
				err = src.Ping(ctx)
			}
			if err != nil {
				return &exitError{1, "conexão: " + logging.Scrub(oraerr.Humanize(err), []string{password})}
			}
			defer src.Close()

			started := time.Now()
			res := inventory.Run(ctx, inventory.Options{
				Src: src, Waves: waves, Scope: scope, Schema: cfg.Oracle.Schema, Tables: tables,
				Count: !noCount, Months: months, QueryTimeout: cfg.Oracle.QueryTimeout, Log: log,
			})
			written, skipped, err := inventory.WriteSpecs(cfg.SourcesDir, res, force)
			if err != nil {
				return err
			}
			if err := inventory.WriteDictionaries(dictDir, res); err != nil {
				return err
			}
			if err := inventory.WriteReport(report, res, dictDir); err != nil {
				return err
			}
			out := cmd.OutOrStdout()
			failed := 0
			for _, t := range res.Tables {
				if t.Err != nil {
					failed++
					fmt.Fprintf(out, "[FAIL] %-22s %s\n", t.Spec.Table, oraerr.Humanize(t.Err))
					continue
				}
				fmt.Fprintf(out, "[OK]   %-22s %d colunas extraídas, chave %s, riscos: %d\n", t.Spec.Table, len(t.Spec.Columns), strings.Join(t.Spec.PrimaryKey.Columns, "+"), len(t.Spec.Risks))
			}
			fmt.Fprintf(out, "\n%d tabelas em %s; %d arquivos escritos, %d aprovados preservados, %d não inventariadas. Relatório: %s\n",
				len(res.Tables), time.Since(started).Round(time.Second), len(written), len(skipped), failed, report)
			if failed > 0 {
				return &exitError{code: 1}
			}
			return nil
		},
	}
	cmd.Flags().StringSliceVar(&tables, "tables", nil, "subset of tables (default: all tables of _waves.yaml)")
	cmd.Flags().BoolVar(&noCount, "no-count", false, "skip COUNT(*), uniqueness and monthly distribution (no full scans)")
	cmd.Flags().BoolVar(&force, "force", false, "overwrite <TABLE>.yaml files even when approved")
	cmd.Flags().IntVar(&months, "months", 36, "months of history for the monthly distribution")
	cmd.Flags().StringVar(&report, "report", filepath.Join("docs", "tables-inventory.md"), "Markdown report path")
	cmd.Flags().StringVar(&dictDir, "dict-dir", filepath.Join("docs", "inventory"), "directory for per-table dictionary dumps")
	return cmd
}
