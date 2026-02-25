package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "data/app.db")
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		uuid TEXT PRIMARY KEY,
		email TEXT UNIQUE,
		password TEXT,
		name TEXT,
		surname TEXT
	);`

	_, err = db.Exec(query)
	return db, err
}
