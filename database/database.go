package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func StartDatabase(connStr string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		log.Panicln(err)
	}

	Db = db

	return nil
}
