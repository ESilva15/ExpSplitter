package models

import (
	"database/sql"
	"expenses/config"
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

	err = m.Migrate(id)
	// err = m.Force(int(id))
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
			log.Printf("migration %d is dirty:", err)

			forceErr := m.Force(int(dirty.Version - 1))
			if forceErr != nil {
				log.Fatalln("force reset failed: %w", forceErr)
			}

			migErr := m.Migrate(uint(dirty.Version))
			log.Fatalln(migErr)
		}
	}

	ver, _, _ := m.Version()
	log.Printf("Successfuly jumped to migration %d", ver)

	return nil
}
