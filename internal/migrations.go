package internal

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

// ApplyMigrations
func ApplyMigrations(db *sql.DB) error {
	migrationFiles := []string {
		"001_init.sql", // Adding all migration files
	}

	for _, file := range migrationFiles {
		migrationPath := filepath.Join("./migrations", file)
		log.Printf("Application of migration: %s", migrationPath)
		content, err := ioutil.ReadFile(migrationPath)
		if err != nil {
			return fmt.Errorf("Failed to read migration files %s: %w", migrationPath, err)
		}
	}

	log.Panicln("Succes migrations")
	return nil
}