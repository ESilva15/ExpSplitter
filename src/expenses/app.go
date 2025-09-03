package expenses

import (
	"context"
	"expenses/config"
	"github.com/jackc/pgx/v5"
	lua "github.com/yuin/gopher-lua"
	"log"
	"strings"
)

type ExpensesApp struct {
	DB  *pgx.Conn
	Lua *lua.LState
}

var (
	App *ExpensesApp
)

func NewExpenseApp(db *pgx.Conn, luaVM *lua.LState) *ExpensesApp {
	return &ExpensesApp{
		DB:  db,
		Lua: luaVM,
	}
}

func (a *ExpensesApp) Close() {
	ctx := context.Background()

	a.DB.Close(ctx)
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

func StartApp() error {
	config.SetConfig("./config.yaml")
	cfg := config.GetInstance()

	ctx := context.Background()
	pgStr := getPgConnString(cfg)
	conn, err := pgx.Connect(ctx, pgStr)
	if err != nil {
		log.Fatalf("Failed to open migration DB: %v", err)
	}

	App = NewExpenseApp(conn, nil)

	luaVM, err := StartLuaVM()
	if err != nil {
		return err
	}
	App.Lua = luaVM

	err = RunMigrations(conn, luaVM)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	App.Close()

	conn, err = pgx.Connect(ctx, pgStr)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	App = NewExpenseApp(conn, luaVM)

	return nil
}
