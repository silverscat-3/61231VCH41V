package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(dbPath string) (*sqlx.DB, error) {
	if _, err := os.Stat(dbPath); nil != err {
		return nil, err
	}

	dbPath = "file:" + dbPath
	db, err := sqlx.Connect("sqlite3", dbPath)
	if nil != err {
		return nil, err
	}

	return db, nil
}
