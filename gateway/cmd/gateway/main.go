// Command gateway is the Winthor extraction gateway. E1.1 ships `doctor`,
// `sqlcheck` and `version`; extraction, service and secret commands arrive in
// later deliverables.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"github.com/the12company/mint/gateway/internal/doctor"
	"github.com/the12company/mint/internal/config"
	"github.com/the12company/mint/internal/logging"
	"github.com/the12company/mint/internal/source/goora"
	"github.com/the12company/mint/internal/source/readonly"
)

// version is set at build time with -ldflags "-X main.version=...".
var version = "dev"

// exitError carries a process exit code through cobra.
type exitError struct {
	code int
	msg  string
}

func (e *exitError) Error() string { return e.msg }

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	restoreConsole := useUTF8Console()
	defer restoreConsole()
	if err := newRoot().ExecuteContext(ctx); err != nil {
		var ee *exitError
		if errors.As(err, &ee) {
			if ee.msg != "" {
				fmt.Fprintln(os.Stderr, ee.msg)
			}
			restoreConsole()
			os.Exit(ee.code)
		}
		fmt.Fprintln(os.Stderr, "erro:", err)
		restoreConsole()
		os.Exit(2)
	}
}

func newRoot() *cobra.Command {
	var configFiles []string
	var logLevel, logFormat string

	root := &cobra.Command{
		Use:           "gateway",
		Short:         "Winthor extraction gateway (read-only)",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.PersistentFlags().StringSliceVar(&configFiles, "config", config.DefaultFiles, "configuration files, overlaid in order (missing files are skipped)")
	root.PersistentFlags().StringVar(&logLevel, "log-level", "", "override logging.level (debug|info|warn|error)")
	root.PersistentFlags().StringVar(&logFormat, "log-format", "", "override logging.format (text|json)")

	loadConfig := func() (config.Config, []string, error) {
		cfg, loaded, err := config.Load(configFiles)
		if err != nil {
			return cfg, loaded, err
		}
		if logLevel != "" {
			cfg.Logging.Level = logLevel
		}
		if logFormat != "" {
			cfg.Logging.Format = logFormat
		}
		return cfg, loaded, nil
	}

	root.AddCommand(newDoctorCmd(loadConfig), newInventoryCmd(loadConfig), newSQLCheckCmd(), newVersionCmd())
	return root
}

func newDoctorCmd(loadConfig func() (config.Config, []string, error)) *cobra.Command {
	var asJSON bool
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Diagnose the Oracle connection, privileges and catalog tables",
		RunE: func(cmd *cobra.Command, _ []string) error {
			cfg, loaded, err := loadConfig()
			if err != nil {
				return &exitError{2, "configuração inválida: " + err.Error()}
			}
			log := logging.New(cmd.ErrOrStderr(), logging.Options{Level: cfg.Logging.Level, Format: cfg.Logging.Format})
			log.Debug("configuration loaded", "files", strings.Join(loaded, ","))
			report := doctor.Run(cmd.Context(), doctor.Options{Config: cfg, Connect: goora.Connect, Logger: log})
			if asJSON {
				if err := doctor.WriteJSON(cmd.OutOrStdout(), report); err != nil {
					return err
				}
			} else {
				doctor.WriteText(cmd.OutOrStdout(), report)
			}
			if report.HasFail() {
				return &exitError{code: 1}
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&asJSON, "json", false, "print the report as JSON")
	return cmd
}

func newSQLCheckCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sqlcheck <file.sql>",
		Short: "Run the read-only guard over every statement of a SQL file (no database access)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := os.ReadFile(args[0])
			if err != nil {
				return &exitError{2, err.Error()}
			}
			refused := 0
			stmts := readonly.SplitStatements(string(data))
			for _, stmt := range stmts {
				label := strings.Join(strings.Fields(stmt), " ")
				if len(label) > 70 {
					label = label[:70] + "…"
				}
				if err := readonly.Check(stmt); err != nil {
					refused++
					fmt.Fprintf(cmd.OutOrStdout(), "[RECUSADO] %s\n           motivo: %v\n", label, err)
				} else {
					fmt.Fprintf(cmd.OutOrStdout(), "[OK]       %s\n", label)
				}
			}
			fmt.Fprintf(cmd.OutOrStdout(), "\n%d instruções, %d recusadas\n", len(stmts), refused)
			if refused > 0 {
				return &exitError{code: 1}
			}
			return nil
		},
	}
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "gateway %s (%s %s/%s)\n", version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
		},
	}
}
