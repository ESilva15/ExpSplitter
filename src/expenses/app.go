package expenses

import (
	"context"
	"expenses/config"
	"expenses/expenses/repo"
	"log"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	lua "github.com/yuin/gopher-lua"
)

type ExpensesApp struct {
	Lua          *lua.LState
	ExpRepo      repo.ExpenseRepository
	CategoryRepo repo.CategoryRepository
	UserRepo     repo.UserRepository
	StoreRepo    repo.StoreRepository
	TypeRepo     repo.TypeRepository
}

var (
	App *ExpensesApp
)

func NewExpenseApp(db *pgxpool.Pool, luaVM *lua.LState) *ExpensesApp {
	return &ExpensesApp{
		ExpRepo:      repo.NewPgExpRepo(db),
		CategoryRepo: repo.NewPgCatRepo(db),
		UserRepo:     repo.NewPgUserRepo(db),
		StoreRepo:    repo.NewPgStoreRepo(db),
		TypeRepo:     repo.NewPgTypeRepo(db),
		Lua:          luaVM,
	}
}

func (a *ExpensesApp) Close() {
	// TODO
	// Close the repos?
	a.Lua.Close()
}

func StartLuaVM() (*lua.LState, error) {
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

func (a *ExpensesApp) GoToMigration(id uint) error {
	migrator, err := repo.NewPgMigrator()
	if err != nil {
		return err
	}

	return migrator.Goto(id)
}

func StartApp() error {
	config.SetConfig("./config.yaml")
	cfg := config.GetInstance()
	ctx := context.Background()
	pgStr := getPgConnString(cfg)

	// Create the LuaVM
	luaVM, err := StartLuaVM()
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
