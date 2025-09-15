package expenses

import (
	"context"
	"encoding/json"
	mod "expenses/expenses/models"
	lua "github.com/yuin/gopher-lua"
)

func (a *ExpApp) prepareExpense(exp *mod.Expense) error {
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

func (a *ExpApp) luaGetAllExpenses(l *lua.LState) int {
	ctx := context.Background()
	// TODO - what do we do about the LUA API
	expenses, err := a.ExpRepo.GetAll(ctx, 1)
	if err != nil {
		return returnWithError(l, err.Error())
	}

	tbl := l.NewTable()
	for _, e := range expenses {
		err = a.prepareExpense(&e)
		if err != nil {
			return returnWithError(l, err.Error())
		}

		jsonData, err := json.Marshal(&e)
		if err != nil {
			return returnWithError(l, err.Error())
		}

		et := l.NewTable()
		et.RawSetString("expense", lua.LString(jsonData))
		tbl.Append(et)
	}

	l.Push(lua.LBool(true))
	l.Push(tbl)
	return 2
}

func (a *ExpApp) luaGetExpense(l *lua.LState) int {
	expId := l.CheckInt(1)

	// TODO
	// This also doesnt make sense
	ctx := context.Background()

	expense, err := a.ExpRepo.Get(ctx, int32(expId))
	if err != nil {
		return returnWithError(l, err.Error())
	}

	err = a.prepareExpense(&expense)
	if err != nil {
		return returnWithError(l, err.Error())
	}

	jsonData, err := json.Marshal(&expense)
	if err != nil {
		return returnWithError(l, err.Error())
	}

	l.Push(lua.LBool(true))
	l.Push(lua.LString(jsonData))
	return 2
}

func (a *ExpApp) luaUpdateExpense(L *lua.LState) int {
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

func (a *ExpApp) luaNormalizeShares(L *lua.LState) int {
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
