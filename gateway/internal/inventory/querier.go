package inventory

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/the12company/mint/internal/source"
)

// querier runs the aggregate queries of the inventory with a timeout.
type querier struct {
	src     source.Source
	timeout time.Duration
}

func (q *querier) each(ctx context.Context, stmt string, scan func(source.Rows) error) error {
	qctx, cancel := context.WithTimeout(ctx, q.timeout)
	defer cancel()
	rows, err := q.src.Query(qctx, stmt)
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

func (q *querier) dbVersion(ctx context.Context) string {
	var product, version string
	_ = q.each(ctx, "SELECT product, version FROM PRODUCT_COMPONENT_VERSION WHERE UPPER(product) LIKE '%ORACLE DATABASE%' AND ROWNUM < 2",
		func(rows source.Rows) error { return rows.Scan(&product, &version) })
	return strings.TrimSpace(product + " " + version)
}

// count returns COUNT(*) of the qualified table.
func (q *querier) count(ctx context.Context, qualified string) (int64, error) {
	var n int64
	err := q.each(ctx, "SELECT COUNT(*) FROM "+qualified, func(rows source.Rows) error { return rows.Scan(&n) })
	return n, err
}

// duplicateGroups counts key values that appear more than once, ignoring
// rows with a NULL key column (Oracle unique constraints ignore them too),
// and returns how many rows have a NULL in any key column.
func (q *querier) duplicateGroups(ctx context.Context, qualified string, key []string) (dups, nullRows int64, err error) {
	cols := quoteList(key)
	var notNull, anyNull []string
	for _, k := range key {
		notNull = append(notNull, `"`+k+`" IS NOT NULL`)
		anyNull = append(anyNull, `"`+k+`" IS NULL`)
	}
	err = q.each(ctx, fmt.Sprintf("SELECT COUNT(*) FROM (SELECT %s FROM %s WHERE %s GROUP BY %s HAVING COUNT(*) > 1)", cols, qualified, strings.Join(notNull, " AND "), cols),
		func(rows source.Rows) error { return rows.Scan(&dups) })
	if err != nil {
		return
	}
	err = q.each(ctx, fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s", qualified, strings.Join(anyNull, " OR ")),
		func(rows source.Rows) error { return rows.Scan(&nullRows) })
	return
}

// rangeStats returns COUNT(*), MIN, MAX and NULL count of a column in one scan.
type rangeStats struct {
	Count    int64
	Min, Max string
	Nulls    int64
}

func (q *querier) rangeStats(ctx context.Context, qualified, col, format string) (rangeStats, error) {
	var rs rangeStats
	var mn, mx sql.NullString
	var nulls sql.NullInt64
	expr := `"` + col + `"`
	minExpr, maxExpr := "TO_CHAR(MIN("+expr+"))", "TO_CHAR(MAX("+expr+"))"
	if format != "" {
		minExpr, maxExpr = "TO_CHAR(MIN("+expr+"), '"+format+"')", "TO_CHAR(MAX("+expr+"), '"+format+"')"
	}
	stmt := fmt.Sprintf("SELECT COUNT(*), %s, %s, SUM(CASE WHEN %s IS NULL THEN 1 ELSE 0 END) FROM %s", minExpr, maxExpr, expr, qualified)
	err := q.each(ctx, stmt, func(rows source.Rows) error { return rows.Scan(&rs.Count, &mn, &mx, &nulls) })
	rs.Min, rs.Max, rs.Nulls = mn.String, mx.String, nulls.Int64
	return rs, err
}

// byMonth returns the row distribution of a DATE column for the last months.
func (q *querier) byMonth(ctx context.Context, qualified, col string, months int) ([]MonthCount, error) {
	expr := `"` + col + `"`
	stmt := fmt.Sprintf("SELECT TO_CHAR(TRUNC(%s, 'MM'), 'YYYY-MM'), COUNT(*) FROM %s WHERE %s >= ADD_MONTHS(TRUNC(SYSDATE, 'MM'), -%d) GROUP BY TRUNC(%s, 'MM') ORDER BY 1",
		expr, qualified, expr, months, expr)
	var out []MonthCount
	err := q.each(ctx, stmt, func(rows source.Rows) error {
		var m MonthCount
		var month sql.NullString
		if err := rows.Scan(&month, &m.Rows); err != nil {
			return err
		}
		m.Month = month.String
		out = append(out, m)
		return nil
	})
	return out, err
}

// fillRate returns COUNT(*) and COUNT(col).
func (q *querier) fillRate(ctx context.Context, qualified, col string) (total, filled int64, err error) {
	err = q.each(ctx, fmt.Sprintf(`SELECT COUNT(*), COUNT("%s") FROM %s`, col, qualified), func(rows source.Rows) error { return rows.Scan(&total, &filled) })
	return
}

// distinctCount returns COUNT(*) and COUNT(DISTINCT col).
func (q *querier) distinctCount(ctx context.Context, qualified, col string) (total, distinct int64, err error) {
	err = q.each(ctx, fmt.Sprintf(`SELECT COUNT(*), COUNT(DISTINCT "%s") FROM %s`, col, qualified), func(rows source.Rows) error { return rows.Scan(&total, &distinct) })
	return
}

func quoteList(cols []string) string {
	quoted := make([]string, len(cols))
	for i, c := range cols {
		quoted[i] = `"` + c + `"`
	}
	return strings.Join(quoted, ", ")
}
