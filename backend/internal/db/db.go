package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
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
