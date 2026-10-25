// Package oradict reads table metadata from the Oracle data dictionary
// (ALL_* views) through a read-only Source. It never reads data rows.
package oradict

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/the12company/mint/internal/catalog"
	"github.com/the12company/mint/internal/source"
)

// ErrNotFound means the table is not visible to the user.
var ErrNotFound = errors.New("table not found or not readable")

// Column is one row of ALL_TAB_COLUMNS plus its comment.
type Column struct {
	Name      string
	DataType  string
	Length    int
	Precision *int
	Scale     *int
	Nullable  bool
	CharUsed  string // B (bytes) or C (chars) for character columns
	ID        int
	Comment   string
}

// TypeString renders the Oracle type as in DDL (NUMBER(10,2), VARCHAR2(40 CHAR)).
func (c Column) TypeString() string {
	switch c.DataType {
	case "NUMBER":
		switch {
		case c.Precision == nil && c.Scale == nil:
			return "NUMBER"
		case c.Precision == nil:
			return fmt.Sprintf("NUMBER(*,%d)", *c.Scale)
		case c.Scale == nil || *c.Scale == 0:
			return fmt.Sprintf("NUMBER(%d)", *c.Precision)
		default:
			return fmt.Sprintf("NUMBER(%d,%d)", *c.Precision, *c.Scale)
		}
	case "VARCHAR2", "CHAR", "NVARCHAR2", "NCHAR", "RAW":
		unit := ""
		if c.CharUsed == "C" {
			unit = " CHAR"
		}
		return fmt.Sprintf("%s(%d%s)", c.DataType, c.Length, unit)
	default:
		return c.DataType
	}
}

// Constraint is a primary (P) or unique (U) key.
type Constraint struct {
	Name    string
	Type    string
	Columns []string
}

// Index is one index with its ordered columns.
type Index struct {
	Name    string
	Unique  bool
	Columns []string
}

// Stats come from ALL_TABLES (optimizer statistics; may be stale or absent).
type Stats struct {
	NumRows      *int64
	Blocks       *int64
	AvgRowLen    *int64
	LastAnalyzed *time.Time
}

// Table is the dictionary view of one table.
type Table struct {
	Owner       string
	Name        string
	Columns     []Column
	Constraints []Constraint
	Indexes     []Index
	Stats       Stats
}

// Column finds a column by name (case-insensitive).
func (t *Table) Column(name string) *Column {
	for i := range t.Columns {
		if strings.EqualFold(t.Columns[i].Name, name) {
			return &t.Columns[i]
		}
	}
	return nil
}

// PrimaryKey returns the enabled primary key constraint, if any.
func (t *Table) PrimaryKey() *Constraint {
	for i := range t.Constraints {
		if t.Constraints[i].Type == "P" {
			return &t.Constraints[i]
		}
	}
	return nil
}

// UniqueConstraints returns enabled unique constraints.
func (t *Table) UniqueConstraints() []Constraint {
	var out []Constraint
	for _, c := range t.Constraints {
		if c.Type == "U" {
			out = append(out, c)
		}
	}
	return out
}

// UniqueIndexes returns unique indexes.
func (t *Table) UniqueIndexes() []Index {
	var out []Index
	for _, i := range t.Indexes {
		if i.Unique {
			out = append(out, i)
		}
	}
	return out
}

// IndexLeading returns an index whose first column is col, if any.
func (t *Table) IndexLeading(col string) *Index {
	for i := range t.Indexes {
		if len(t.Indexes[i].Columns) > 0 && strings.EqualFold(t.Indexes[i].Columns[0], col) {
			return &t.Indexes[i]
		}
	}
	return nil
}

// Qualified returns "OWNER"."TABLE" after validating both identifiers.
func Qualified(owner, table string) (string, error) {
	if !catalog.ValidIdentifier(owner) || !catalog.ValidIdentifier(table) {
		return "", fmt.Errorf("invalid identifier in %s.%s", owner, table)
	}
	return `"` + owner + `"."` + table + `"`, nil
}

// query is a small helper around Source.Query.
func query(ctx context.Context, src source.Source, stmt string, args []any, scan func(source.Rows) error) error {
	rows, err := src.Query(ctx, stmt, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := scan(rows); err != nil {
			return err
		}
	}
	return rows.Err()
}

// ResolveOwner finds the schema that owns table: the configured schema, a
// private or public synonym, or the single owner in ALL_TABLES.
func ResolveOwner(ctx context.Context, src source.Source, table, schema string) (string, error) {
	if !catalog.ValidIdentifier(table) {
		return "", fmt.Errorf("invalid table name %q", table)
	}
	if schema = strings.ToUpper(strings.TrimSpace(schema)); schema != "" {
		if !catalog.ValidIdentifier(schema) {
			return "", fmt.Errorf("invalid schema %q", schema)
		}
		return schema, nil
	}
	var owner sql.NullString
	err := query(ctx, src, "SELECT table_owner FROM ALL_SYNONYMS WHERE synonym_name = :1 AND owner IN ('PUBLIC', USER) ORDER BY CASE owner WHEN 'PUBLIC' THEN 2 ELSE 1 END",
		[]any{table}, func(rows source.Rows) error {
			var o sql.NullString
			if err := rows.Scan(&o); err != nil {
				return err
			}
			if !owner.Valid {
				owner = o
			}
			return nil
		})
	if err != nil {
		return "", err
	}
	if owner.Valid && owner.String != "" {
		return owner.String, nil
	}
	var owners []string
	err = query(ctx, src, "SELECT owner FROM ALL_TABLES WHERE table_name = :1 ORDER BY owner", []any{table}, func(rows source.Rows) error {
		var o string
		if err := rows.Scan(&o); err != nil {
			return err
		}
		owners = append(owners, o)
		return nil
	})
	if err != nil {
		return "", err
	}
	switch len(owners) {
	case 0:
		return "", fmt.Errorf("%s: %w", table, ErrNotFound)
	case 1:
		return owners[0], nil
	default:
		return "", fmt.Errorf("%s: more than one owner (%s); set oracle.schema", table, strings.Join(owners, ", "))
	}
}
