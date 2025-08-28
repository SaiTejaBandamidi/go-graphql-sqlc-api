package graph

import (
	"database/sql"

	"github.com/SaiTejaBandamidi/go-graphql-sqlc-api/db"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB      *sql.DB
	Queries *db.Queries
}
