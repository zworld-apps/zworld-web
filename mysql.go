package main

import (
	"database/sql"
	"errors"
	"os"
)

func NewMySQLDatabase() (db *sql.DB, err error) {
	dbURL := os.Getenv("CLEARDB_DATABASE_URL")
	if dbURL == "" {
		return nil, errors.New("Not using ClearDB for MySQL")
	}

	db, err = sql.Open("mysql", dbURL)
	return
}
