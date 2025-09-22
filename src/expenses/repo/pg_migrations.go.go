package repo

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ESilva15/expenses/config"

	mig "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	lua "github.com/yuin/gopher-lua"
)

// LastGoMigration defines the last migration execute with go code. After
// that the lua API should be used.
const LastGoMigration = 3

// PgMigrator implements the Migrator interface for the PG connection.
type PgMigrator struct {
	Mig *mig.Migrate
}

// NewPgMigrator creates and returns a new migrator.
func NewPgMigrator(connStr string) (PgMigrator, error) {
	cfg := config.GetInstance()

	m, err := mig.New(cfg.MigrationsPath, connStr)
	if err != nil {
		return PgMigrator{}, err
	}

	return PgMigrator{
		Mig: m,
	}, nil
}

// Close closes what is necessary to close on a PgMigrator.
func (p PgMigrator) Close() {
	p.Mig.Close()
}

// Goto forces the migrations to a given migrations of id: `id`.
func (p PgMigrator) Goto(id uint) error {
	err := p.Mig.Migrate(id)
	if err != nil && !errors.Is(err, mig.ErrNoChange) {
		return err
	}

	log.Println("migrated to", id)

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

func (p PgMigrator) RunMigrations(lua *lua.LState) error {
	err := p.Mig.Up()
	if err != nil {
		if err == mig.ErrNoChange {
			return nil
		}

		if dirty, ok := err.(mig.ErrDirty); ok {
			log.Printf("migration %d is dirty:", err)

			forceErr := p.Mig.Force(int(dirty.Version - 1))
			if forceErr != nil {
				log.Fatalln("force reset failed: %w", forceErr)
			}

			migErr := p.Mig.Migrate(uint(dirty.Version))
			log.Fatalln(migErr)
		}
	}

	ver, _, _ := p.Mig.Version()
	log.Printf("Successfuly jumped to migration %d", ver)

	err = runCustomScript(ver, lua)
	if err != nil {
		return err
	}

	return nil
}
