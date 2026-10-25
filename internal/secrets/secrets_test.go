package secrets

import (
	"errors"
	"testing"
)

func TestResolve(t *testing.T) {
	env := map[string]string{"ORACLE_PASSWORD": "s3cret"}
	lookup := func(k string) (string, bool) { v, ok := env[k]; return v, ok }

	v, kind, err := resolve("env:ORACLE_PASSWORD", lookup)
	if err != nil || v != "s3cret" || kind != Env {
		t.Fatalf("env resolve: v=%q kind=%v err=%v", v, kind, err)
	}
	if _, _, err := resolve("env:MISSING", lookup); err == nil {
		t.Fatal("expected error for missing variable")
	}
	if _, _, err := resolve("env:", lookup); err == nil {
		t.Fatal("expected error for empty variable name")
	}
	v, kind, err = resolve("plain", lookup)
	if err != nil || v != "plain" || kind != Literal {
		t.Fatalf("literal resolve: v=%q kind=%v err=%v", v, kind, err)
	}
	if _, kind, err := resolve("dpapi:AQAAANCM", lookup); !errors.Is(err, ErrDPAPIUnavailable) || kind != DPAPI {
		t.Fatalf("dpapi: kind=%v err=%v", kind, err)
	}
}
