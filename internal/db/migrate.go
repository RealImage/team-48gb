package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations runs all pending migrations
func RunMigrations(mongoURI, dbName, migrationsPath string) error {
	// Construct the MongoDB URI for migrations
	// directConnection=true bypasses replica set discovery for local MongoDB
	databaseURI := fmt.Sprintf("mongodb://127.0.0.1:27017/%s?directConnection=true", dbName)

	// Create migration instance
	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		databaseURI,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	// Run all pending migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	if err == migrate.ErrNoChange {
		log.Println("No new migrations to run")
	} else {
		log.Println("Migrations completed successfully")
	}

	return nil
}

// MigrateDown rolls back the last migration
func MigrateDown(mongoURI, dbName, migrationsPath string) error {
	databaseURI := fmt.Sprintf("mongodb://127.0.0.1:27017/%s?directConnection=true", dbName)

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		databaseURI,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to rollback migration: %w", err)
	}

	log.Println("Migration rolled back successfully")
	return nil
}

// MigrateVersion gets the current migration version
func MigrateVersion(mongoURI, dbName, migrationsPath string) (uint, bool, error) {
	databaseURI := fmt.Sprintf("mongodb://127.0.0.1:27017/%s?directConnection=true", dbName)

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		databaseURI,
	)
	if err != nil {
		return 0, false, fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return 0, false, fmt.Errorf("failed to get migration version: %w", err)
	}

	return version, dirty, nil
}
