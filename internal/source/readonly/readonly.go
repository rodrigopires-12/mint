// Package readonly is the second line of defence against writes to the
// customer's Oracle database (the first is the database user having only
// SELECT). Every statement must start with SELECT or WITH, be a single
// statement, and must not contain FOR UPDATE or inline PL/SQL.
package readonly

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/the12company/mint/internal/source"
)

// ErrNotReadOnly wraps every refusal so callers can test for it.
var ErrNotReadOnly = errors.New("statement refused by read-only guard")

var (
	firstWordRe = regexp.MustCompile(`^[A-Za-z]+`)
	forUpdateRe = regexp.MustCompile(`\bFOR\s+UPDATE\b`)
	plsqlWithRe = regexp.MustCompile(`^WITH\s+(FUNCTION|PROCEDURE)\b`)
)

// Check returns nil when stmt is a single SELECT/WITH query, or an error
// wrapping ErrNotReadOnly otherwise.
func Check(stmt string) error {
	trimmed := strings.TrimSpace(mask(stmt))
	if trimmed == "" {
		return fmt.Errorf("%w: empty statement", ErrNotReadOnly)
	}
	if strings.Contains(trimmed, ";") {
		return fmt.Errorf("%w: ';' is not allowed (one statement per query)", ErrNotReadOnly)
	}
	upper := strings.ToUpper(trimmed)
	first := firstWordRe.FindString(upper)
	switch first {
	case "SELECT":
	case "WITH":
		if plsqlWithRe.MatchString(upper) {
			return fmt.Errorf("%w: inline PL/SQL (WITH FUNCTION/PROCEDURE) is not allowed", ErrNotReadOnly)
		}
	default:
		if first == "" {
			first = string([]rune(trimmed)[0])
		}
		return fmt.Errorf("%w: statement starts with %q; only SELECT and WITH are allowed", ErrNotReadOnly, first)
	}
	if forUpdateRe.MatchString(upper) {
		return fmt.Errorf("%w: FOR UPDATE is not allowed", ErrNotReadOnly)
	}
	return nil
}

// mask returns stmt with every character inside comments and string literals
// replaced by a space (newlines are kept), so keyword checks cannot be fooled
// by text inside quotes or comments. The result has the same rune length and
// offsets as the input, which SplitStatements relies on.
func mask(s string) string {
	r := []rune(s)
	out := make([]rune, len(r))
	copy(out, r)
	blank := func(i int) {
		if out[i] != '\n' {
			out[i] = ' '
		}
	}
	for i := 0; i < len(r); i++ {
		switch {
		case r[i] == '-' && i+1 < len(r) && r[i+1] == '-':
			for i < len(r) && r[i] != '\n' {
				blank(i)
				i++
			}
		case r[i] == '/' && i+1 < len(r) && r[i+1] == '*':
			blank(i)
			blank(i + 1)
			i += 2
			for i < len(r) && !(r[i] == '*' && i+1 < len(r) && r[i+1] == '/') {
				blank(i)
				i++
			}
			if i < len(r) {
				blank(i)
				blank(i + 1)
				i++
			}
		case r[i] == '\'':
			i++
			for i < len(r) {
				if r[i] == '\'' {
					if i+1 < len(r) && r[i+1] == '\'' { // escaped quote
						blank(i)
						blank(i + 1)
						i += 2
						continue
					}
					break
				}
				blank(i)
				i++
			}
		case r[i] == '"':
			i++
			for i < len(r) && r[i] != '"' {
				blank(i)
				i++
			}
		}
	}
	return string(out)
}

// SplitStatements splits a SQL script on ';' outside string literals and
// comments, returning the original text of each non-empty statement. It
// serves `gateway sqlcheck` and the verification scripts; the guard itself
// never accepts more than one statement.
func SplitStatements(script string) []string {
	orig := []rune(script)
	masked := []rune(mask(script))
	var out []string
	start := 0
	for i := 0; i <= len(masked); i++ {
		if i == len(masked) || masked[i] == ';' {
			if strings.TrimSpace(string(masked[start:i])) != "" {
				out = append(out, strings.TrimSpace(string(orig[start:i])))
			}
			start = i + 1
		}
	}
	return out
}

// Guard wraps a Source so every Query passes through Check.
func Guard(inner source.Source) source.Source { return &guarded{inner: inner} }

type guarded struct{ inner source.Source }

func (g *guarded) Ping(ctx context.Context) error { return g.inner.Ping(ctx) }
func (g *guarded) Close() error                   { return g.inner.Close() }

func (g *guarded) Query(ctx context.Context, stmt string, args ...any) (source.Rows, error) {
	if err := Check(stmt); err != nil {
		return nil, err
	}
	return g.inner.Query(ctx, stmt, args...)
}
