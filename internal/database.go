package internal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Database initialization
func InitDB(filepath) (*sql.DB, error) {
	db, err := sql.Open("*sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	// Create a table if it doesn't exist
	err = createTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
 // Create a table of posts
func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
