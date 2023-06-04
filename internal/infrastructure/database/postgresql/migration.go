package postgresql

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"log"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate runs the database migrations for the PostgreSQL database.
func Migrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"../../internal/infrastructure/database/postgresql/migrations",
		"users", driver)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply database migrations: %v", err)
	}

	return nil
}
