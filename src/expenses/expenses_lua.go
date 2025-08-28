package expenses

import (
	"encoding/json"
	mod "expenses/expenses/models"

	lua "github.com/yuin/gopher-lua"
)

func (a *ExpensesApp) luaGetAllExpenses(L *lua.LState) int {
	tx, err := a.DB.Begin()
	if err != nil {
		return returnWithError(L, err.Error())
	}
	defer tx.Rollback()

	expenses, err := mod.GetAllExpenses(tx)

	tbl := L.NewTable()
	for _, e := range expenses {
		err := a.LoadExpensePayments(&e)
		if err != nil {
			return returnWithError(L, err.Error())
		}

		err = a.LoadExpenseShares(&e)
		if err != nil {
			return returnWithError(L, err.Error())
		}

		jsonData, err := json.Marshal(e)

		et := L.NewTable()
		et.RawSetString("expense", lua.LString(jsonData))
		tbl.Append(et)
	}

	L.Push(tbl)
	return 1
}

func (a *ExpensesApp) luaNormalizeShares(L *lua.LState) int {
	expenseJson := L.CheckString(1)

	expense, err := mod.ExpenseFromJSON([]byte(expenseJson))
	if err != nil {
		return returnWithError(L, err.Error())
	}

	err = a.NormalizeShares(expense)

	jsonData, err := json.Marshal(expense)
	if err != nil {
		return returnWithError(L, err.Error())
	}

	resultExpense := L.NewTable()
	resultExpense.RawSetString("expense", lua.LString(jsonData))

	L.Push(resultExpense)
	return 1
}
