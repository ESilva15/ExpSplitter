package expenses

import (
	"github.com/ESilva15/expenses/luadec"

	ljson "github.com/layeh/gopher-json"
	lua "github.com/yuin/gopher-lua"
)

func (a *ExpApp) registerLuaBinds(L *lua.LState) {
	ljson.Preload(L)
	// Add the decimal.Decimal operations
	L.SetGlobal("AddDecimal", L.NewFunction(luadec.AddDecimal))

	// Add our API
	L.SetGlobal("GetAllExpenses", L.NewFunction(a.luaGetAllExpenses))
	L.SetGlobal("GetExpense", L.NewFunction(a.luaGetExpense))
	L.SetGlobal("UpdateExpense", L.NewFunction(a.luaUpdateExpense))
	L.SetGlobal("NormalizeShare", L.NewFunction(a.luaNormalizeShares))

	L.SetGlobal("InsertShare", L.NewFunction(a.luaInsertShare))

	L.SetGlobal("InsertPayment", L.NewFunction(a.luaInsertPayment))
}

func returnWithError(L *lua.LState, strErr string) int {
	L.Push(lua.LBool(false))
	L.Push(lua.LString(strErr))
	return 2
}
