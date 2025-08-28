package expenses

import (
	lua "github.com/yuin/gopher-lua"
)

func (a *ExpensesApp) luaInsertShare(L *lua.LState) int {
	// shareJson := L.CheckString(1)

	// share, err := mod.ShareFromJSON([]byte(shareJson))
	// if err != nil {
	// 	return ReturnWithError(L, err.Error())
	// }

	return 1
}
