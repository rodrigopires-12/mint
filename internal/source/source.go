// Package source defines the read-only interface to a customer's Oracle
// database. Unit tests use sourcetest.Mock; production uses goora (default)
// or godror (build tag "godror", from E1.5).
package source

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// Rows is a forward-only cursor, deliberately shaped like database/sql.Rows.
type Rows interface {
	Columns() ([]string, error)
	Next() bool
	Scan(dest ...any) error
	Err() error
	Close() error
}

// Source is a connection (or pool) to one Oracle database.
// Implementations must only ever execute SELECT/WITH statements; the readonly
// package enforces that for every implementation.
type Source interface {
	Ping(ctx context.Context) error
	Query(ctx context.Context, stmt string, args ...any) (Rows, error)
	Close() error
}

// ConnectParams are the inputs every driver needs.
type ConnectParams struct {
	DSN            string // EZCONNECT host[:port]/service
	User           string
	Password       string
	ConnectTimeout time.Duration
	MaxSessions    int
}

// Connector opens a Source. The goora and godror packages provide one each.
type Connector func(ctx context.Context, p ConnectParams) (Source, error)

// EZConnect is a parsed host[:port]/service string.
type EZConnect struct {
	Host    string
	Port    int
	Service string
}

var ezRe = regexp.MustCompile(`^([A-Za-z0-9][A-Za-z0-9.\-]*)(?::([0-9]{1,5}))?/([A-Za-z0-9][A-Za-z0-9_.\-]*)$`)

// ParseEZConnect parses "host[:port]/service"; the port defaults to 1521.
func ParseEZConnect(dsn string) (EZConnect, error) {
	m := ezRe.FindStringSubmatch(dsn)
	if m == nil {
		return EZConnect{}, fmt.Errorf("invalid dsn %q: expected host[:port]/service", dsn)
	}
	port := 1521
	if m[2] != "" {
		p, err := strconv.Atoi(m[2])
		if err != nil || p <= 0 || p > 65535 {
			return EZConnect{}, fmt.Errorf("invalid port in dsn %q", dsn)
		}
		port = p
	}
	return EZConnect{Host: m[1], Port: port, Service: m[3]}, nil
}

// String renders the EZConnect form (never includes credentials).
func (e EZConnect) String() string { return fmt.Sprintf("%s:%d/%s", e.Host, e.Port, e.Service) }
