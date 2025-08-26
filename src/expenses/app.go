package expenses

import (
	"database/sql"
	"expenses/config"
	mod "expenses/expenses/models"
	"expenses/luadec"
	"log"

	lua "github.com/yuin/gopher-lua"
)

var (
	Serv *Service
)

type Service struct {
	DB  *sql.DB
	Lua *lua.LState
}

func NewExpenseService(db *sql.DB, luaVM *lua.LState) *Service {
	return &Service{
		DB:  db,
		Lua: luaVM,
	}
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

	L.SetGlobal("GetAllExpenses", L.NewFunction(Serv.LuaGetAllExpenses))
	L.SetGlobal("AddDecimal", L.NewFunction(luadec.AddDecimal))

	return L, nil
}

func StartApp() error {
	config.SetConfig("./config.yaml")
	cfg := config.GetInstance()
	luaVM, err := StartLuaVM()
	if err != nil {
		return err
	}

	migDB, err := openDB(cfg.DBSys, cfg.DBPath, "")
	if err != nil {
		log.Fatalf("Failed to open migration DB: %v", err)
	}

	err = mod.RunMigrations(migDB, luaVM)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	migDB.Close()

	db, err := openDB(cfg.DBSys, cfg.DBPath, "?_foreign_keys=on")
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	Serv = NewExpenseService(db, luaVM)

	return nil
}
