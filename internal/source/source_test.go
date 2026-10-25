package source

import "testing"

func TestParseEZConnect(t *testing.T) {
	cases := map[string]string{
		"db.local/WINTHOR":          "db.local:1521/WINTHOR",
		"10.0.0.5:1522/orcl.dom":    "10.0.0.5:1522/orcl.dom",
		"host-a:1521/SVC_1":         "host-a:1521/SVC_1",
		"user:pass@host:1521/SVC":   "",
		"host":                      "",
		"host:abc/SVC":              "",
		"host:70000/SVC":            "",
		"oracle://host:1521/SVC":    "",
		"host:1521/SVC; drop table": "",
	}
	for in, want := range cases {
		got, err := ParseEZConnect(in)
		if want == "" {
			if err == nil {
				t.Errorf("%q: expected error, got %v", in, got)
			}
			continue
		}
		if err != nil || got.String() != want {
			t.Errorf("%q: got %v (%v), want %s", in, got, err, want)
		}
	}
}
