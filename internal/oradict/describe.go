package oradict

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/the12company/mint/internal/source"
)

// Describe reads columns, comments, key constraints, indexes and optimizer
// statistics of owner.table.
func Describe(ctx context.Context, src source.Source, owner, table string) (*Table, error) {
	if _, err := Qualified(owner, table); err != nil {
		return nil, err
	}
	t := &Table{Owner: owner, Name: table}
	args := []any{owner, table}

	err := query(ctx, src, `SELECT column_name, data_type, data_length, data_precision, data_scale, nullable, char_used, column_id
FROM ALL_TAB_COLUMNS WHERE owner = :1 AND table_name = :2 ORDER BY column_id`, args, func(rows source.Rows) error {
		var c Column
		var prec, scale sql.NullInt64
		var nullable string
		var charUsed sql.NullString
		var length, id int64
		if err := rows.Scan(&c.Name, &c.DataType, &length, &prec, &scale, &nullable, &charUsed, &id); err != nil {
			return err
		}
		c.Length, c.ID = int(length), int(id)
		if prec.Valid {
			v := int(prec.Int64)
			c.Precision = &v
		}
		if scale.Valid {
			v := int(scale.Int64)
			c.Scale = &v
		}
		c.Nullable = nullable == "Y"
		c.CharUsed = charUsed.String
		t.Columns = append(t.Columns, c)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("ALL_TAB_COLUMNS: %w", err)
	}
	if len(t.Columns) == 0 {
		return nil, fmt.Errorf("%s.%s: %w", owner, table, ErrNotFound)
	}

	err = query(ctx, src, `SELECT column_name, comments FROM ALL_COL_COMMENTS WHERE owner = :1 AND table_name = :2 AND comments IS NOT NULL`, args,
		func(rows source.Rows) error {
			var name string
			var comment sql.NullString
			if err := rows.Scan(&name, &comment); err != nil {
				return err
			}
			if c := t.Column(name); c != nil {
				c.Comment = comment.String
			}
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("ALL_COL_COMMENTS: %w", err)
	}

	byName := map[string]*Constraint{}
	err = query(ctx, src, `SELECT c.constraint_name, c.constraint_type, cc.column_name
FROM ALL_CONSTRAINTS c JOIN ALL_CONS_COLUMNS cc ON cc.owner = c.owner AND cc.constraint_name = c.constraint_name
WHERE c.owner = :1 AND c.table_name = :2 AND c.constraint_type IN ('P','U') AND c.status = 'ENABLED'
ORDER BY c.constraint_type, c.constraint_name, cc.position`, args, func(rows source.Rows) error {
		var name, typ, col string
		if err := rows.Scan(&name, &typ, &col); err != nil {
			return err
		}
		k, ok := byName[name]
		if !ok {
			t.Constraints = append(t.Constraints, Constraint{Name: name, Type: typ})
			k = &t.Constraints[len(t.Constraints)-1]
			byName[name] = k
		}
		k.Columns = append(k.Columns, col)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("ALL_CONSTRAINTS: %w", err)
	}

	idxByName := map[string]int{}
	err = query(ctx, src, `SELECT i.index_name, i.uniqueness, ic.column_name
FROM ALL_INDEXES i JOIN ALL_IND_COLUMNS ic ON ic.index_owner = i.owner AND ic.index_name = i.index_name
WHERE i.table_owner = :1 AND i.table_name = :2 ORDER BY i.index_name, ic.column_position`, args, func(rows source.Rows) error {
		var name, uniq, col string
		if err := rows.Scan(&name, &uniq, &col); err != nil {
			return err
		}
		pos, ok := idxByName[name]
		if !ok {
			t.Indexes = append(t.Indexes, Index{Name: name, Unique: uniq == "UNIQUE"})
			pos = len(t.Indexes) - 1
			idxByName[name] = pos
		}
		t.Indexes[pos].Columns = append(t.Indexes[pos].Columns, col)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("ALL_INDEXES: %w", err)
	}

	err = query(ctx, src, `SELECT num_rows, blocks, avg_row_len, last_analyzed FROM ALL_TABLES WHERE owner = :1 AND table_name = :2`, args,
		func(rows source.Rows) error {
			var numRows, blocks, avg sql.NullInt64
			var analyzed sql.NullTime
			if err := rows.Scan(&numRows, &blocks, &avg, &analyzed); err != nil {
				return err
			}
			if numRows.Valid {
				t.Stats.NumRows = &numRows.Int64
			}
			if blocks.Valid {
				t.Stats.Blocks = &blocks.Int64
			}
			if avg.Valid {
				t.Stats.AvgRowLen = &avg.Int64
			}
			if analyzed.Valid {
				v := analyzed.Time
				t.Stats.LastAnalyzed = &v
			}
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("ALL_TABLES: %w", err)
	}
	return t, nil
}
