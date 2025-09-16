// Package expenses will be the core of the expenses splitter webapp's functionalities
package expenses

import (
	"context"
	"log"
	"strings"

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
	var s strings.Builder

	s.WriteString("port=" + cfg.PgCfg.Port + " ")
	s.WriteString("host=" + cfg.PgCfg.Host + " ")
	s.WriteString("user=" + cfg.PgCfg.User + " ")
	s.WriteString("dbname=" + cfg.PgCfg.DB + " ")
	s.WriteString("password=" + cfg.PgCfg.Pass)

	return s.String()
}

func runMigrations(luaVM *lua.LState) error {
	migrator, err := repo.NewPgMigrator()
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
	migrator, err := repo.NewPgMigrator()
	if err != nil {
		return err
	}

	return migrator.Goto(id)
}

// StartApp will configure and open the required connections to run this app.
func StartApp() error {
	config.SetConfig("./config.yaml")
	cfg := config.GetInstance()
	ctx := context.Background()
	pgStr := getPgConnString(cfg)

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
	conn, err := pgxpool.New(ctx, pgStr)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	App = NewExpenseApp(conn, luaVM)

	return nil
}
