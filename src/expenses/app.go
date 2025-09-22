// Package expenses will be the core of the expenses splitter webapp's functionalities
package expenses

import (
	"context"
	"fmt"
	"log"

	"github.com/ESilva15/expenses/config"
	"github.com/ESilva15/expenses/expenses/repo"

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

var (
	// App provides access to the expenses main service.
	App *ExpApp
)

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

func runMigrations(luaVM *lua.LState) error {
	cfg := config.GetInstance()
	pgStr := getPgConnString(cfg)

	migrator, err := repo.NewPgMigrator(pgStr)
	if err != nil {
		return err
	}

	err = migrator.RunMigrations(luaVM)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	migrator.Close()

	return nil
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
		panic("connection test query failed: " + err.Error())
	}

	return conn, nil
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

	// Run the migrations
	err = runMigrations(luaVM)
	if err != nil {
		log.Printf("Unable to run migrations: %v", err)
	}

	// Create our final app thing
	conn, err := createPGXPoolConn(ctx)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	App = NewExpenseApp(conn, luaVM)

	return nil
}
