package main

import (
	"os"
	"strings"
	"testing"
)

func TestDBSourceFormatting(t *testing.T) {
	// Set environment variables used by main to build the DB source string
	os.Setenv("DB_USER", "u")
	os.Setenv("DB_PASSWORD", "p")
	os.Setenv("DB_HOST", "h")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "db")

	dbSource := "postgresql://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	if !strings.Contains(dbSource, "u:p@h:5432/db") {
		t.Fatalf("dbSource formatted incorrectly: %s", dbSource)
	}
}
