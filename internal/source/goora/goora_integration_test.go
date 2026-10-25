//go:build integration

package goora

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/the12company/mint/internal/source"
)

// TestIntegrationConnect runs only against the company's test Oracle, when
// ORACLE_DSN, ORACLE_USER and ORACLE_PASSWORD are set. It never reads data
// rows, only DUAL.
func TestIntegrationConnect(t *testing.T) {
	dsn := os.Getenv("ORACLE_DSN")
	if dsn == "" {
		t.Skip("ORACLE_DSN not set")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	src, err := Connect(ctx, source.ConnectParams{
		DSN: dsn, User: os.Getenv("ORACLE_USER"), Password: os.Getenv("ORACLE_PASSWORD"),
		ConnectTimeout: 15 * time.Second, MaxSessions: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer src.Close()
	if err := src.Ping(ctx); err != nil {
		t.Fatal(err)
	}
	rows, err := src.Query(ctx, "SELECT 1 FROM DUAL")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	if !rows.Next() {
		t.Fatal("no row from DUAL")
	}
	if _, err := src.Query(ctx, "UPDATE DUAL SET DUMMY = 'Y'"); err == nil {
		t.Fatal("guard let an UPDATE through")
	}
}
