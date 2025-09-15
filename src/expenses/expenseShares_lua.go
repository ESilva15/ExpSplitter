package expenses

import (
	mod "expenses/expenses/models"
	lua "github.com/yuin/gopher-lua"
)

func (a *ExpApp) luaInsertShare(L *lua.LState) int {
	shareJson := L.CheckString(1)
	expenseID := L.CheckInt(2)

	share, err := mod.ShareFromJSON([]byte(shareJson))
	if err != nil {
		return returnWithError(L, err.Error())
	}

	err = a.insertShare(*share, int32(expenseID))
	if err != nil {
		return returnWithError(L, err.Error())
	}

	L.Push(lua.LBool(true))
	L.Push(lua.LString(""))
	return 2
}
