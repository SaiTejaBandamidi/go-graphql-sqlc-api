package internal

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB(dataSourceName string) *sql.DB {

	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		panic(err)

	}
	return db

}
