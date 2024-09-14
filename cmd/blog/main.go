package main

import (
	"database/sql"
	"log"
)

func main() {
	// Loading the configuration
	cfg, err := config.LoodConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Connecting to the database
	db, err := sql.Open("sqlite3", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}