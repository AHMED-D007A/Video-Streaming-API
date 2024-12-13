package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDBConnection(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
