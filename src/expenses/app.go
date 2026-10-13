// Package expenses will be the core of the expenses splitter webapp's functionalities
package expenses

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/ESilva15/expenses/config"
	"github.com/ESilva15/expenses/expenses/repo"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	lua "github.com/yuin/gopher-lua"
)

// ExpApp is the data structure for the main runtime of this application.
type ExpApp struct {
	Lua          *lua.LState
	ExpRepo      repo.ExpenseRepository
	CategoryRepo repo.CategoryRepository
	UserRepo     repo.UserRepository
	StoreRepo    repo.StoreRepository
	TypeRepo     repo.TypeRepository
}

// App provides access to the expenses main service.
var App *ExpApp

// NewExpenseApp returns a new instance of our core App.
func NewExpenseApp(dbPool *pgxpool.Pool, luaVM *lua.LState) *ExpApp {
	return &ExpApp{
		ExpRepo:      repo.NewPgExpRepo(dbPool),
		CategoryRepo: repo.NewPgCatRepo(dbPool),
		UserRepo:     repo.NewPgUserRepo(dbPool),
		StoreRepo:    repo.NewPgStoreRepo(dbPool),
		TypeRepo:     repo.NewPgTypeRepo(dbPool),
		Lua:          luaVM,
	}
}

// Close will close this ExpApp instance.
func (a *ExpApp) Close() {
	// TODO - Close the repos?
	a.Lua.Close()
}

// startLuaVM creates the LuaVM we require to run our custom scripts.
func startLuaVM() (*lua.LState, error) {
	L := lua.NewState()
	App.registerLuaBinds(L)

	return L, nil
}

func getPgConnString(cfg *config.Configuration) string {
	return fmt.Sprintf("postgres://%s:%s@/%s?host=%s&port=%s&sslmode=disable",
		cfg.PgCfg.User, cfg.PgCfg.Pass, cfg.PgCfg.DB, cfg.PgCfg.Host, cfg.PgCfg.Port)
}

func createTargetDatabase(cfg *config.Configuration) error {
	// Connect to the default 'postgres' database instead of cfg.PgCfg.DB
	adminConnStr := fmt.Sprintf("postgres://%s:%s@/%s?host=%s&port=%s&sslmode=disable",
		cfg.PgCfg.User, cfg.PgCfg.Pass, "postgres", cfg.PgCfg.Host, cfg.PgCfg.Port)

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, adminConnStr)
	if err != nil {
		return fmt.Errorf("failed to connect to default postgres db: %w", err)
	}
	defer pool.Close()

	// Execute CREATE DATABASE (Note: database identifiers cannot be parameterized with $1)
	query := fmt.Sprintf("CREATE DATABASE %s", cfg.PgCfg.DB)
	_, err = pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	log.Printf("Database %q created successfully!", cfg.PgCfg.DB)
	return nil
}

func runMigrations(luaVM *lua.LState) error {
	cfg := config.GetInstance()
	pgStr := getPgConnString(cfg)

	migrator, err := repo.NewPgMigrator(pgStr)
	if err != nil {
		// Check if error is due to database not existing
		if isDatabaseNotFoundError(err) {
			log.Printf("Database %q does not exist. Creating it...", cfg.PgCfg.DB)
			if createErr := createTargetDatabase(cfg); createErr != nil {
				return fmt.Errorf("failed to create missing database: %w", createErr)
			}

			// Retry creating the migrator after the database is created
			migrator, err = repo.NewPgMigrator(pgStr)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	err = migrator.RunMigrations(luaVM)
	if err != nil {
		migrator.Close()
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	migrator.Close()
	return nil
}

// Helper to check for PostgreSQL code "3D000" or text fallback
func isDatabaseNotFoundError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "3D000" {
		return true
	}

	// Fallback string check if wrapped differently by the repo layer
	return strings.Contains(err.Error(), "does not exist")
}

// GoToMigration provides a way to force moving to the specified migration.
func (a *ExpApp) GoToMigration(id uint) error {
	cfg := config.GetInstance()
	pgStr := getPgConnString(cfg)

	migrator, err := repo.NewPgMigrator(pgStr)
	if err != nil {
		return err
	}

	return migrator.Goto(id)
}

func createPGXPoolConn(ctx context.Context) (*pgxpool.Pool, error) {
	cfg := config.GetInstance()
	pgStr := getPgConnString(cfg)

	conn, err := pgxpool.New(ctx, pgStr)
	if err != nil {
		return nil, err
	}

	var result int
	err = conn.QueryRow(ctx, "SELECT 1").Scan(&result)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func PrepareDB(ctx context.Context, luaVM *lua.LState) error {
	// Attempt to prepare the database
	err := runMigrations(luaVM)
	if err != nil {
		log.Fatalf("Unable to run migrations: %v", err)
	}

	return nil
}

func CheckDatabase(ctx context.Context, luaVM *lua.LState) error {
	cfg := config.GetInstance()
	pgStr := getPgConnString(cfg)

	conn, err := pgxpool.New(ctx, pgStr)
	if err != nil {
		// Unable to create the database
		return err
	}

	var result int
	err = conn.QueryRow(ctx, "SELECT status,version FROM dbstat").Scan(&result)
	if err != nil {
		if err := PrepareDB(ctx, luaVM); err != nil {
			return err
		}
	}

	return nil
}

// StartApp will configure and open the required connections to run this app.
func StartApp() error {
	config.SetConfig("./config.yaml")
	ctx := context.Background()

	// Create the LuaVM
	luaVM, err := startLuaVM()
	if err != nil {
		return err
	}

	// Run migrations (will automatically create the database if it doesn't exist)
	err = runMigrations(luaVM)
	if err != nil {
		log.Fatalf("Unable to run migrations: %v", err)
	}

	// Create our final app pool connection
	conn, err := createPGXPoolConn(ctx)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	App = NewExpenseApp(conn, luaVM)

	return nil
}
