// Package sourcetest provides an in-memory source.Source for unit tests.
package sourcetest

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/the12company/mint/internal/source"
)

// Handler answers queries whose text contains Match (case-insensitive).
type Handler struct {
	Match   string
	Columns []string
	Rows    [][]any
	Err     error
}

// Mock is a scripted Source. Queries with no matching handler fail loudly.
type Mock struct {
	PingErr  error
	Handlers []Handler

	mu      sync.Mutex
	Queries []string
	Closed  bool
}

func (m *Mock) Ping(context.Context) error { return m.PingErr }

func (m *Mock) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Closed = true
	return nil
}

func (m *Mock) Query(_ context.Context, stmt string, _ ...any) (source.Rows, error) {
	m.mu.Lock()
	m.Queries = append(m.Queries, stmt)
	m.mu.Unlock()
	up := strings.ToUpper(stmt)
	for _, h := range m.Handlers {
		if strings.Contains(up, strings.ToUpper(h.Match)) {
			if h.Err != nil {
				return nil, h.Err
			}
			return &rows{cols: h.Columns, data: h.Rows, i: -1}, nil
		}
	}
	return nil, fmt.Errorf("sourcetest: no handler for query %q", stmt)
}

type rows struct {
	cols []string
	data [][]any
	i    int
	err  error
}

func (r *rows) Columns() ([]string, error) { return r.cols, nil }
func (r *rows) Err() error                 { return r.err }
func (r *rows) Close() error               { return nil }

func (r *rows) Next() bool {
	r.i++
	return r.i < len(r.data)
}

func (r *rows) Scan(dest ...any) error {
	if r.i < 0 || r.i >= len(r.data) {
		return fmt.Errorf("sourcetest: Scan called without Next")
	}
	row := r.data[r.i]
	if len(row) != len(dest) {
		return fmt.Errorf("sourcetest: row has %d values, Scan got %d destinations", len(row), len(dest))
	}
	for j, d := range dest {
		if err := assign(d, row[j]); err != nil {
			return err
		}
	}
	return nil
}

func assign(dest, val any) error {
	switch d := dest.(type) {
	case *string:
		switch v := val.(type) {
		case nil:
			*d = ""
		case string:
			*d = v
		default:
			*d = fmt.Sprint(v)
		}
	case *sql.NullString:
		switch v := val.(type) {
		case nil:
			*d = sql.NullString{}
		case string:
			*d = sql.NullString{String: v, Valid: true}
		default:
			*d = sql.NullString{String: fmt.Sprint(v), Valid: true}
		}
	case *int64:
		switch v := val.(type) {
		case int:
			*d = int64(v)
		case int64:
			*d = v
		case float64:
			*d = int64(v)
		case nil:
			*d = 0
		default:
			return fmt.Errorf("sourcetest: cannot assign %T to *int64", val)
		}
	case *time.Time:
		switch v := val.(type) {
		case time.Time:
			*d = v
		case nil:
			*d = time.Time{}
		default:
			return fmt.Errorf("sourcetest: cannot assign %T to *time.Time", val)
		}
	case *any:
		*d = val
	default:
		return fmt.Errorf("sourcetest: unsupported scan destination %T", dest)
	}
	return nil
}
