package internal

import (
    "testing"
)

func TestConnectDB_ReturnsDB(t *testing.T) {
    // Provide an intentionally invalid DSN; ConnectDB should return a *sql.DB
    // (it currently calls sql.Open and panics on error only on Open failure).
    dsn := "user=invalid password=invalid host=localhost port=5432 dbname=invalid sslmode=disable"
    db := ConnectDB(dsn)
    if db == nil {
        t.Fatalf("expected non-nil *sql.DB, got nil")
    }
    // Close the DB to avoid leaks
    _ = db.Close()
}
