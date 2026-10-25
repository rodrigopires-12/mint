// Package secrets resolves secret references found in configuration values.
//
// Supported forms:
//
//	env:NAME        value of the environment variable NAME
//	dpapi:<base64>  ciphertext produced by `gateway secret set` (Windows DPAPI, E1.3)
//	anything else   literal value
package secrets

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Kind tells how a value was resolved.
type Kind int

const (
	Literal Kind = iota
	Env
	DPAPI
)

func (k Kind) String() string {
	switch k {
	case Env:
		return "env"
	case DPAPI:
		return "dpapi"
	default:
		return "literal"
	}
}

// ErrDPAPIUnavailable is returned until DPAPI support lands in E1.3.
var ErrDPAPIUnavailable = errors.New("dpapi: secrets are not supported yet (planned for E1.3)")

// Resolve interprets a configuration value and returns the secret it denotes.
func Resolve(value string) (string, Kind, error) {
	return resolve(value, os.LookupEnv)
}

func resolve(value string, lookup func(string) (string, bool)) (string, Kind, error) {
	switch {
	case strings.HasPrefix(value, "env:"):
		name := strings.TrimSpace(strings.TrimPrefix(value, "env:"))
		if name == "" {
			return "", Env, errors.New("env: reference without a variable name")
		}
		v, ok := lookup(name)
		if !ok || v == "" {
			return "", Env, fmt.Errorf("environment variable %s is not set", name)
		}
		return v, Env, nil
	case strings.HasPrefix(value, "dpapi:"):
		return "", DPAPI, ErrDPAPIUnavailable
	default:
		return value, Literal, nil
	}
}
