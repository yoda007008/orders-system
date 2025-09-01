package migrator

import (
	"database/sql"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

func RunMigrations(connString string, migrationsPath string) error {
	abcPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		fmt.Errorf("Please get absolute PATH")
	}

	slog.Info("Apply migrations on these path: ", abcPath)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return fmt.Errorf("Error openning DB")
	}

	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Errorf("Error driver Postgres")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+abcPath,
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration up failed: %w", err)
	}

	slog.Info("Migrations was succesfully")
	return nil
}
