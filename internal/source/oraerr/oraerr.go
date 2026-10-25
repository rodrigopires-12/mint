// Package oraerr turns Oracle driver errors into short, credential-free
// messages for operators (pt-BR, since they are shown in the CLI).
package oraerr

import (
	"context"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

var codeRe = regexp.MustCompile(`ORA-(\d{5})`)

// Code extracts the first ORA-NNNNN code from err, or "" when there is none.
func Code(err error) string {
	if err == nil {
		return ""
	}
	return codeRe.FindString(err.Error())
}

var messages = map[string]string{
	"ORA-01017": "usuário ou senha inválidos",
	"ORA-28000": "conta bloqueada",
	"ORA-28001": "senha expirada",
	"ORA-28002": "senha vai expirar em breve",
	"ORA-01005": "senha não informada",
	"ORA-12154": "alias TNS não resolvido",
	"ORA-12514": "serviço desconhecido no listener (confira o service name do DSN)",
	"ORA-12505": "SID desconhecido no listener (confira o DSN)",
	"ORA-12541": "nenhum listener no host e porta informados",
	"ORA-12543": "host inacessível",
	"ORA-12170": "tempo esgotado ao conectar (rede ou firewall)",
	"ORA-12537": "conexão fechada pelo servidor",
	"ORA-00942": "tabela ou view não existe (ou sem privilégio de leitura)",
	"ORA-01031": "privilégios insuficientes",
	"ORA-00904": "identificador inválido",
	"ORA-01555": "snapshot too old (undo insuficiente)",
	"ORA-03113": "conexão perdida com o banco",
	"ORA-03114": "não conectado ao banco",
	"ORA-01000": "cursores abertos demais",
}

// Humanize returns a one-line explanation without stack traces or secrets.
func Humanize(err error) string {
	if err == nil {
		return ""
	}
	if code := Code(err); code != "" {
		if msg, ok := messages[code]; ok {
			return fmt.Sprintf("%s (%s)", msg, code)
		}
		return "erro Oracle " + code + ": " + firstLine(err.Error(), 160)
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return "tempo esgotado"
	}
	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return "tempo esgotado na rede: " + firstLine(err.Error(), 120)
		}
		return "falha de rede: " + firstLine(err.Error(), 120)
	}
	return firstLine(err.Error(), 200)
}

func firstLine(s string, max int) string {
	if i := strings.IndexAny(s, "\r\n"); i >= 0 {
		s = s[:i]
	}
	s = strings.TrimSpace(s)
	if len(s) > max {
		return s[:max] + "…"
	}
	return s
}
