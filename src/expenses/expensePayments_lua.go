package expenses

import (
	mod "github.com/ESilva15/expenses/expenses/models"
	lua "github.com/yuin/gopher-lua"
)

func (a *ExpApp) luaInsertPayment(L *lua.LState) int {
	paymentJSON := L.CheckString(1)
	expenseID := L.CheckInt(2)

	payment, err := mod.PaymentFromJSON([]byte(paymentJSON))
	if err != nil {
		return returnWithError(L, err.Error())
	}

	err = a.insertPayment(*payment, int32(expenseID))
	if err != nil {
		return returnWithError(L, err.Error())
	}

	L.Push(lua.LBool(true))
	L.Push(lua.LString(""))
	return 2
}
