// Package inventory builds the per-table specifications (E1.2) from the
// scope seed and the Oracle dictionary, and writes docs/tables-inventory.md.
package inventory

import (
	"fmt"
	"strings"

	"github.com/the12company/mint/internal/oradict"
)

// Risk labels attached to a column type mapping.
const (
	RiskLOB               = "lob"
	RiskNumberNoPrecision = "number_without_precision"
	RiskTimeZone          = "timestamp_with_time_zone"
	RiskUnknownType       = "unknown_type"
)

// ClickHouseType proposes the raw-layer ClickHouse type for an Oracle column.
// The raw layer keeps source fidelity: integers stay integers, decimals keep
// precision and scale, DATE keeps its time part, strings stay strings.
// Nullability is handled separately by the caller.
func ClickHouseType(c oradict.Column) (string, string) {
	dt := strings.ToUpper(c.DataType)
	switch {
	case dt == "NUMBER" || dt == "INTEGER" || dt == "DECIMAL" || dt == "NUMERIC":
		return numberType(c)
	case dt == "FLOAT" || dt == "BINARY_DOUBLE" || dt == "DOUBLE PRECISION" || dt == "REAL":
		return "Float64", ""
	case dt == "BINARY_FLOAT":
		return "Float32", ""
	case dt == "DATE":
		return "DateTime64(0)", ""
	case strings.HasPrefix(dt, "TIMESTAMP"):
		scale := 6
		if c.Scale != nil && *c.Scale >= 0 && *c.Scale <= 9 {
			scale = *c.Scale
		}
		if strings.Contains(dt, "TIME ZONE") {
			return fmt.Sprintf("DateTime64(%d)", scale), RiskTimeZone
		}
		return fmt.Sprintf("DateTime64(%d)", scale), ""
	case dt == "VARCHAR2" || dt == "VARCHAR" || dt == "CHAR" || dt == "NVARCHAR2" || dt == "NCHAR" || dt == "ROWID" || dt == "UROWID":
		return "String", ""
	case dt == "CLOB" || dt == "NCLOB" || dt == "BLOB" || dt == "LONG" || dt == "LONG RAW" || dt == "RAW" || dt == "BFILE" || strings.Contains(dt, "XMLTYPE"):
		return "String", RiskLOB
	default:
		return "String", RiskUnknownType
	}
}

func numberType(c oradict.Column) (string, string) {
	if c.Precision == nil {
		if c.Scale != nil && *c.Scale == 0 {
			// NUMBER(*,0): integer of unbounded precision.
			return "Int64", RiskNumberNoPrecision
		}
		return "Decimal(38,10)", RiskNumberNoPrecision
	}
	p := *c.Precision
	s := 0
	if c.Scale != nil {
		s = *c.Scale
	}
	switch {
	case s < 0:
		return "Decimal(38,0)", ""
	case s == 0 && p <= 18:
		return "Int64", ""
	case s == 0:
		return fmt.Sprintf("Decimal(%d,0)", p), ""
	default:
		if p < s {
			p = s
		}
		return fmt.Sprintf("Decimal(%d,%d)", p, s), ""
	}
}

// IsLOB reports whether the Oracle type is a LOB/LONG that must stay out of raw.
func IsLOB(c oradict.Column) bool {
	_, risk := ClickHouseType(c)
	return risk == RiskLOB
}
