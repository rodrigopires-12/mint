package oraerr

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestHumanize(t *testing.T) {
	err := errors.New("ORA-01017: invalid username/password; logon denied\nstack frame 1\nstack frame 2")
	got := Humanize(err)
	if !strings.Contains(got, "ORA-01017") || !strings.Contains(got, "usuário ou senha") || strings.Contains(got, "stack") {
		t.Fatalf("got %q", got)
	}
	if Code(err) != "ORA-01017" {
		t.Fatalf("code %q", Code(err))
	}
	if got := Humanize(errors.New("ORA-99999: strange")); !strings.HasPrefix(got, "erro Oracle ORA-99999") {
		t.Fatalf("got %q", got)
	}
	if got := Humanize(context.DeadlineExceeded); got != "tempo esgotado" {
		t.Fatalf("got %q", got)
	}
	if Humanize(nil) != "" || Code(nil) != "" {
		t.Fatal("nil handling")
	}
}
