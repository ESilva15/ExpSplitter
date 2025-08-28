package expenses

import (
	"expenses/luadec"

	ljson "github.com/layeh/gopher-json"
	lua "github.com/yuin/gopher-lua"
)

func (a *ExpensesApp) registerLuaBinds(L *lua.LState) {
	ljson.Preload(L)
	L.SetGlobal("GetAllExpenses", L.NewFunction(a.luaGetAllExpenses))
	L.SetGlobal("AddDecimal", L.NewFunction(luadec.AddDecimal))
	L.SetGlobal("NormalizeShare", L.NewFunction(a.luaNormalizeShares))
}

func returnWithError(L *lua.LState, strErr string) int {
	L.Push(lua.LNil)
	L.Push(lua.LString(strErr))
	return 2
}
