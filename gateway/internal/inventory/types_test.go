package inventory

import (
	"testing"

	"github.com/the12company/mint/internal/oradict"
)

func ip(v int) *int { return &v }

func TestClickHouseType(t *testing.T) {
	cases := []struct {
		col      oradict.Column
		wantType string
		wantRisk string
	}{
		{oradict.Column{DataType: "NUMBER", Precision: ip(10), Scale: ip(0)}, "Int64", ""},
		{oradict.Column{DataType: "NUMBER", Precision: ip(10)}, "Int64", ""},
		{oradict.Column{DataType: "NUMBER", Precision: ip(22), Scale: ip(0)}, "Decimal(22,0)", ""},
		{oradict.Column{DataType: "NUMBER", Precision: ip(18), Scale: ip(6)}, "Decimal(18,6)", ""},
		{oradict.Column{DataType: "NUMBER"}, "Decimal(38,10)", RiskNumberNoPrecision},
		{oradict.Column{DataType: "NUMBER", Scale: ip(0)}, "Int64", RiskNumberNoPrecision},
		{oradict.Column{DataType: "FLOAT", Precision: ip(126)}, "Float64", ""},
		{oradict.Column{DataType: "DATE"}, "DateTime64(0)", ""},
		{oradict.Column{DataType: "TIMESTAMP(6)", Scale: ip(6)}, "DateTime64(6)", ""},
		{oradict.Column{DataType: "TIMESTAMP(6) WITH TIME ZONE", Scale: ip(6)}, "DateTime64(6)", RiskTimeZone},
		{oradict.Column{DataType: "VARCHAR2", Length: 40}, "String", ""},
		{oradict.Column{DataType: "CHAR", Length: 1}, "String", ""},
		{oradict.Column{DataType: "CLOB"}, "String", RiskLOB},
		{oradict.Column{DataType: "LONG"}, "String", RiskLOB},
		{oradict.Column{DataType: "SDO_GEOMETRY"}, "String", RiskUnknownType},
	}
	for _, c := range cases {
		gotType, gotRisk := ClickHouseType(c.col)
		if gotType != c.wantType || gotRisk != c.wantRisk {
			t.Errorf("%s: got (%s, %s) want (%s, %s)", c.col.TypeString(), gotType, gotRisk, c.wantType, c.wantRisk)
		}
	}
}

func TestOracleTypeString(t *testing.T) {
	cases := map[string]oradict.Column{
		"NUMBER(10,2)":      {DataType: "NUMBER", Precision: ip(10), Scale: ip(2)},
		"NUMBER(10)":        {DataType: "NUMBER", Precision: ip(10), Scale: ip(0)},
		"NUMBER":            {DataType: "NUMBER"},
		"VARCHAR2(40 CHAR)": {DataType: "VARCHAR2", Length: 40, CharUsed: "C"},
		"VARCHAR2(40)":      {DataType: "VARCHAR2", Length: 40, CharUsed: "B"},
		"DATE":              {DataType: "DATE"},
	}
	for want, col := range cases {
		if got := col.TypeString(); got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
}
