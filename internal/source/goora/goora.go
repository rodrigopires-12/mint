// Package goora implements source.Source with the pure-Go driver
// github.com/sijms/go-ora (no Instant Client, no cgo). It is the default
// driver; see docs/adr/0003-driver-oracle.md.
package goora

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strconv"
	"time"

	go_ora "github.com/sijms/go-ora/v2"

	"github.com/the12company/mint/internal/source"
	"github.com/the12company/mint/internal/source/readonly"
)

// Connect opens a pooled connection. The returned Source is wrapped by the
// read-only guard; there is no way to obtain an unguarded one.
func Connect(_ context.Context, p source.ConnectParams) (source.Source, error) {
	ez, err := source.ParseEZConnect(p.DSN)
	if err != nil {
		return nil, err
	}
	opts := map[string]string{}
	if p.ConnectTimeout > 0 {
		opts["CONNECTION TIMEOUT"] = strconv.Itoa(int(math.Ceil(p.ConnectTimeout.Seconds())))
	}
	url := go_ora.BuildUrl(ez.Host, ez.Port, ez.Service, p.User, p.Password, opts)
	db, err := sql.Open("oracle", url)
	if err != nil {
		return nil, fmt.Errorf("open oracle driver: %w", err)
	}
	if p.MaxSessions > 0 {
		db.SetMaxOpenConns(p.MaxSessions)
		db.SetMaxIdleConns(min(p.MaxSessions, 2))
	}
	db.SetConnMaxLifetime(30 * time.Minute)
	return readonly.Guard(&sqlSource{db: db}), nil
}

type sqlSource struct{ db *sql.DB }

func (s *sqlSource) Ping(ctx context.Context) error { return s.db.PingContext(ctx) }
func (s *sqlSource) Close() error                   { return s.db.Close() }

func (s *sqlSource) Query(ctx context.Context, stmt string, args ...any) (source.Rows, error) {
	rows, err := s.db.QueryContext(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
