package doctor

import (
	"fmt"
	"sort"
	"strings"
)

// Privilege rules are pure functions so they can be unit-tested without a
// database. They decide whether the Oracle user is read-only.

var allowedSystemPrivileges = map[string]bool{
	"CREATE SESSION":        true,
	"ALTER SESSION":         true,
	"SET CONTAINER":         true,
	"SELECT ANY DICTIONARY": true,
	"FLASHBACK ANY TABLE":   true,
}

var broadReadPrivileges = map[string]bool{
	"SELECT ANY TABLE": true,
	"READ ANY TABLE":   true,
}

var allowedRoles = map[string]bool{
	"CONNECT":              true,
	"SELECT_CATALOG_ROLE":  true,
	"HS_ADMIN_SELECT_ROLE": true,
}

var writeRoles = map[string]bool{
	"DBA":                        true,
	"RESOURCE":                   true,
	"IMP_FULL_DATABASE":          true,
	"EXP_FULL_DATABASE":          true,
	"DATAPUMP_EXP_FULL_DATABASE": true,
	"DATAPUMP_IMP_FULL_DATABASE": true,
	"EXECUTE_CATALOG_ROLE":       true,
	"DELETE_CATALOG_ROLE":        true,
	"SCHEDULER_ADMIN":            true,
	"AQ_ADMINISTRATOR_ROLE":      true,
}

// dangerousCreatePrivileges are CREATE privileges that reach beyond the
// user's own schema or the database (OS commands, public objects).
var dangerousCreatePrivileges = map[string]bool{
	"CREATE EXTERNAL JOB":         true,
	"CREATE PUBLIC SYNONYM":       true,
	"CREATE PUBLIC DATABASE LINK": true,
	"CREATE ROLE":                 true,
	"CREATE USER":                 true,
	"CREATE PROFILE":              true,
	"CREATE TABLESPACE":           true,
	"CREATE ROLLBACK SEGMENT":     true,
	"CREATE LIBRARY":              true,
}

// ownSchemaOnly reports whether a system privilege can only create objects
// inside the grantee's own schema (CREATE TABLE, CREATE JOB, CREATE
// CREDENTIAL, UNLIMITED TABLESPACE...). Such privileges cannot touch the
// Winthor tables, so they are reported as WARN rather than FAIL.
func ownSchemaOnly(p string) bool {
	if p == "UNLIMITED TABLESPACE" {
		return true
	}
	if !strings.HasPrefix(p, "CREATE ") || dangerousCreatePrivileges[p] {
		return false
	}
	return !strings.Contains(p, " ANY ") && !strings.Contains(p, "PUBLIC")
}

// AssessSystemPrivileges classifies system privileges: FAIL for anything that
// can change data or objects outside the user's schema, WARN for broad read
// access or own-schema object creation, OK otherwise.
func AssessSystemPrivileges(privs []string) Result {
	var bad, broad, own []string
	for _, p := range privs {
		p = strings.ToUpper(strings.TrimSpace(p))
		switch {
		case allowedSystemPrivileges[p]:
		case broadReadPrivileges[p]:
			broad = append(broad, p)
		case ownSchemaOnly(p):
			own = append(own, p)
		default:
			bad = append(bad, p)
		}
	}
	sort.Strings(bad)
	sort.Strings(broad)
	sort.Strings(own)
	switch {
	case len(bad) > 0:
		return Result{"privilégios de sistema", Fail, fmt.Sprintf("permitem escrita ou alcançam outros schemas: %s (o usuário do gateway deve ter só CREATE SESSION e SELECT nas tabelas)", strings.Join(bad, ", "))}
	case len(own)+len(broad) > 0:
		var parts []string
		if len(own) > 0 {
			parts = append(parts, fmt.Sprintf("criam objetos só no schema do próprio usuário e não alteram as tabelas do Winthor, mas recomendamos revogar: %s", strings.Join(own, ", ")))
		}
		if len(broad) > 0 {
			parts = append(parts, fmt.Sprintf("leitura ampla (funciona, mas prefira SELECT por tabela): %s", strings.Join(broad, ", ")))
		}
		return Result{"privilégios de sistema", Warn, strings.Join(parts, ". ")}
	default:
		return Result{"privilégios de sistema", OK, fmt.Sprintf("%d privilégios, todos de leitura ou sessão", len(privs))}
	}
}

// AssessRoles flags administrative or write-capable roles.
func AssessRoles(roles []string) Result {
	var bad, unknown []string
	for _, r := range roles {
		r = strings.ToUpper(strings.TrimSpace(r))
		switch {
		case allowedRoles[r]:
		case writeRoles[r]:
			bad = append(bad, r)
		default:
			unknown = append(unknown, r)
		}
	}
	sort.Strings(bad)
	sort.Strings(unknown)
	switch {
	case len(bad) > 0:
		return Result{"papéis", Fail, fmt.Sprintf("papéis com escrita ou administração: %s", strings.Join(bad, ", "))}
	case len(unknown) > 0:
		return Result{"papéis", Warn, fmt.Sprintf("papéis não reconhecidos: %s (os privilégios efetivos foram avaliados em SESSION_PRIVS e ROLE_TAB_PRIVS)", strings.Join(unknown, ", "))}
	default:
		return Result{"papéis", OK, fmt.Sprintf("%d papéis, todos de leitura", len(roles))}
	}
}

// TablePrivilege is one row of USER_TAB_PRIVS or ROLE_TAB_PRIVS.
type TablePrivilege struct {
	Owner, Table, Privilege string
}

// AssessTablePrivileges flags object privileges other than SELECT/READ/FLASHBACK.
func AssessTablePrivileges(privs []TablePrivilege) Result {
	var bad []string
	for _, p := range privs {
		switch strings.ToUpper(strings.TrimSpace(p.Privilege)) {
		case "SELECT", "READ", "FLASHBACK":
		default:
			bad = append(bad, fmt.Sprintf("%s em %s.%s", p.Privilege, p.Owner, p.Table))
		}
	}
	if len(bad) > 0 {
		sort.Strings(bad)
		shown := bad
		more := ""
		if len(shown) > 5 {
			shown = shown[:5]
			more = fmt.Sprintf(" e mais %d", len(bad)-5)
		}
		return Result{"privilégios de objeto", Fail, fmt.Sprintf("%d além de leitura: %s%s", len(bad), strings.Join(shown, "; "), more)}
	}
	return Result{"privilégios de objeto", OK, fmt.Sprintf("%d privilégios de objeto, todos de leitura", len(privs))}
}
