package graph

import (
	"database/sql"
	"testing"

	"github.com/SaiTejaBandamidi/go-graphql-sqlc-api/db"
)

// A minimal fake that satisfies db.DBTX used by db.New; we only need *db.Queries pointer
type fakeDB struct{}

func TestResolver_Wiring(t *testing.T) {
	// Build a Resolver with nil DB and a nil Queries pointer (allowed)
	r := &Resolver{DB: &sql.DB{}, Queries: &db.Queries{}}

	if r.Query() == nil {
		t.Fatalf("expected Query() to return non-nil resolver")
	}
	if r.Mutation() == nil {
		t.Fatalf("expected Mutation() to return non-nil resolver")
	}
}
