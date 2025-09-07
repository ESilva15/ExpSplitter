package expenses

import (
	"context"
	"encoding/json"
	mod "expenses/expenses/models"
	lua "github.com/yuin/gopher-lua"
)

func (a *ExpensesApp) prepareExpense(exp *mod.Expense) error {
	err := a.LoadExpensePayments(exp)
	if err != nil {
		return err
	}

	err = a.LoadExpenseShares(exp)
	if err != nil {
		return err
	}

	return nil
}

func (a *ExpensesApp) luaGetAllExpenses(L *lua.LState) int {
	ctx := context.Background()
	expenses, err := a.ExpRepo.GetAll(ctx)

	tbl := L.NewTable()
	for _, e := range expenses {
		err = a.prepareExpense(&e)
		if err != nil {
			return returnWithError(L, err.Error())
		}

		jsonData, err := json.Marshal(&e)
		if err != nil {
			return returnWithError(L, err.Error())
		}

		et := L.NewTable()
		et.RawSetString("expense", lua.LString(jsonData))
		tbl.Append(et)
	}

	L.Push(lua.LBool(true))
	L.Push(tbl)
	return 2
}

func (a *ExpensesApp) luaGetExpense(L *lua.LState) int {
	expId := L.CheckInt(1)

	// TODO
	// This also doesnt make sense
	ctx := context.Background()

	expense, err := a.ExpRepo.Get(ctx, int32(expId))
	if err != nil {
		return returnWithError(L, err.Error())
	}

	err = a.prepareExpense(&expense)
	if err != nil {
		return returnWithError(L, err.Error())
	}

	jsonData, err := json.Marshal(&expense)
	if err != nil {
		return returnWithError(L, err.Error())
	}

	L.Push(lua.LBool(true))
	L.Push(lua.LString(jsonData))
	return 2
}

func (a *ExpensesApp) luaUpdateExpense(L *lua.LState) int {
	expJson := L.CheckString(1)

	exp, err := mod.ExpenseFromJSON([]byte(expJson))
	if err != nil {
		return returnWithError(L, err.Error())
	}

	err = a.UpdateExpense(*exp)
	if err != nil {
		return returnWithError(L, err.Error())
	}

	L.Push(lua.LBool(true))
	L.Push(lua.LString(""))
	return 2
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

	L.Push(lua.LBool(true))
	L.Push(lua.LString(jsonData))
	return 2
}
