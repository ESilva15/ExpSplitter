package expenses

import (
	"database/sql"
	"expenses/config"
	"log"

	lua "github.com/yuin/gopher-lua"
)

type ExpensesApp struct {
	DB  *sql.DB
	Lua *lua.LState
}

var (
	App *ExpensesApp
)

func NewExpenseApp(db *sql.DB, luaVM *lua.LState) *ExpensesApp {
	return &ExpensesApp{
		DB:  db,
		Lua: luaVM,
	}
}

func (a *ExpensesApp) Close() {
	a.DB.Close()
	a.Lua.Close()
}

func openDB(sys string, path string, extra string) (*sql.DB, error) {
	db, err := sql.Open(sys, "file:"+path+extra)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func StartLuaVM() (*lua.LState, error) {
	L := lua.NewState()
	App.registerLuaBinds(L)
	return L, nil
}

func StartApp() error {
	config.SetConfig("./config.yaml")
	cfg := config.GetInstance()

	migDB, err := openDB(cfg.DBSys, cfg.DBPath, "")
	if err != nil {
		log.Fatalf("Failed to open migration DB: %v", err)
	}

	App = NewExpenseApp(migDB, nil)

	luaVM, err := StartLuaVM()
	if err != nil {
		return err
	}
	App.Lua = luaVM

	err = RunMigrations(migDB, luaVM)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	App.Close()

	db, err := openDB(cfg.DBSys, cfg.DBPath, "?_foreign_keys=on")
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	App = NewExpenseApp(db, luaVM)

	return nil
}
