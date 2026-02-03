package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS expenses (
		id TEXT PRIMARY KEY,
		amount INTEGER NOT NULL,
		category TEXT NOT NULL,
		description TEXT,
		date TEXT NOT NULL,
		created_at TEXT NOT NULL
	);`

	if _, err := db.Exec(schema); err != nil {
		log.Fatal(err)
	}

	return db
}
