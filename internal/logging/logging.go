// Package logging builds the structured logger used by every component.
// It redacts attributes whose key looks like a secret and scrubs known secret
// values from string attributes, so credentials never reach a log line.
package logging

import (
	"io"
	"log/slog"
	"strings"
)

// Options configures the logger.
type Options struct {
	Level  string   // debug | info | warn | error
	Format string   // text | json
	Redact []string // literal secret values to scrub from every string attribute
}

var secretKeyFragments = []string{"password", "passwd", "secret", "token", "authorization", "credential"}

// New returns a slog.Logger writing to w.
func New(w io.Writer, opts Options) *slog.Logger {
	h := &slog.HandlerOptions{Level: ParseLevel(opts.Level), ReplaceAttr: redactor(opts.Redact)}
	var handler slog.Handler
	if strings.EqualFold(opts.Format, "json") {
		handler = slog.NewJSONHandler(w, h)
	} else {
		handler = slog.NewTextHandler(w, h)
	}
	return slog.New(handler)
}

// ParseLevel converts a level name to slog.Level; unknown values map to Info.
func ParseLevel(s string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func redactor(values []string) func([]string, slog.Attr) slog.Attr {
	return func(_ []string, a slog.Attr) slog.Attr {
		key := strings.ToLower(a.Key)
		for _, frag := range secretKeyFragments {
			if strings.Contains(key, frag) {
				return slog.String(a.Key, "***")
			}
		}
		if a.Value.Kind() == slog.KindString {
			return slog.String(a.Key, Scrub(a.Value.String(), values))
		}
		return a
	}
}

// Scrub replaces every occurrence of the given secret values in s with "***".
// Values shorter than three characters are skipped: replacing them would
// mangle ordinary words, and such a secret offers no protection anyway.
func Scrub(s string, values []string) string {
	for _, v := range values {
		if len(v) >= 3 {
			s = strings.ReplaceAll(s, v, "***")
		}
	}
	return s
}
