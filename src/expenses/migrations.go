package expenses

import (
	"database/sql"
	"expenses/config"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	mig "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	lua "github.com/yuin/gopher-lua"
)

const LAST_GO_MIGRATION = 3

func prepareMigrator(db *sql.DB) (*mig.Migrate, error) {
	cfg := config.GetInstance()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	m, err := mig.NewWithDatabaseInstance(cfg.MigrationsPath, "sqlite3", driver)

	return m, err
}

func (a *ExpensesApp) Goto(id uint) error {
	m, err := prepareMigrator(a.DB)
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

func listCustomScripts() (map[uint]string, error) {
	// List the available scripts
	cfg := config.GetInstance()

	dir, err := os.Open(cfg.MigCustomScript)
	if err != nil {
		return make(map[uint]string), err
	}

	files, err := dir.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	customScripts := make(map[uint]string)
	for _, f := range files {
		parts := strings.SplitN(f.Name(), "_", 2)
		numStr := parts[0]
		num, err := strconv.ParseUint(numStr, 10, 16)
		if err != nil {
			return customScripts, err
		}

		customScripts[uint(num)] = filepath.Join(cfg.MigCustomScript, f.Name())
	}

	return customScripts, nil
}

func runCustomScript(ver uint, lua *lua.LState) error {
	customScripts, err := listCustomScripts()
	if err != nil {
		return err
	}

	if val, ok := customScripts[ver]; ok {
		if err := lua.DoFile(val); err != nil {
			log.Printf("Failed to run lua script: %+v", err)
		}
	}

	return nil
}

func runCustomMigrationLogic(ver uint) {
	switch ver {
	case 4:
		// Get all expenses, normalize the shares, update them
	}
}

func RunMigrations(db *sql.DB, lua *lua.LState) error {
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

	// One day we will have custom stuff being ran in Lua instead of here
	if ver > LAST_GO_MIGRATION {
		err = runCustomScript(ver, lua)
		if err != nil {
			return err
		}
	} else {
		runCustomMigrationLogic(ver)
	}

	return nil
}
