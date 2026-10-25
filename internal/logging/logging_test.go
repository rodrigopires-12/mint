package logging

import (
	"bytes"
	"strings"
	"testing"
)

func TestRedaction(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf, Options{Level: "debug", Format: "json", Redact: []string{"s3cret"}})
	log.Info("connect", "password", "s3cret", "user", "ro", "dsn", "host/svc", "msgtext", "pw is s3cret here")
	out := buf.String()
	if strings.Contains(out, "s3cret") {
		t.Fatalf("secret leaked: %s", out)
	}
	if !strings.Contains(out, `"password":"***"`) || !strings.Contains(out, `"user":"ro"`) {
		t.Fatalf("unexpected output: %s", out)
	}
}
