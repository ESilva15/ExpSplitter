package models

import (
	"database/sql"
	"expenses/config"
	"fmt"
	"log"

	mig "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func prepareMigrator(db *sql.DB) (*mig.Migrate, error) {
	cfg := config.GetInstance()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	m, err := mig.NewWithDatabaseInstance(cfg.MigrationsPath, "sqlite3", driver)

	return m, err
}

func Goto(db *sql.DB, id uint) error {
	m, err := prepareMigrator(db)
	if err != nil {
		return err
	}

	err = m.Force(int(id))
	if err != nil && err != mig.ErrNoChange {
		return err
	}

	return nil
}

func RunMigrations(db *sql.DB) error {
	m, err := prepareMigrator(db)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if err == mig.ErrNoChange {
			return nil
		}

		if dirty, ok := err.(mig.ErrDirty); ok {
			log.Printf("Database is dirty at version %d, cleaning", dirty.Version)

			if ferr := m.Force(dirty.Version); ferr != nil {
				return fmt.Errorf("failed to force clean state: %w", ferr)
			}

			if uerr := m.Up(); uerr != nil && uerr != mig.ErrNoChange {
				return fmt.Errorf("retry migration failed: %w", uerr)
			}

			return nil
		}
	}

	return nil
}
