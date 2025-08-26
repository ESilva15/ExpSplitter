package expenses

import (
	"encoding/json"
	mod "expenses/expenses/models"

	lua "github.com/yuin/gopher-lua"
)

func (s *Service) LuaGetAllExpenses(L *lua.LState) int {
	tx, err := s.DB.Begin()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	defer tx.Rollback()

	expenses, err := mod.GetAllExpenses(tx)

	tbl := L.NewTable()
	for _, e := range expenses {
		err := s.LoadExpensePayments(&e)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			continue
		}

		err = s.LoadExpenseShares(&e)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			continue
		}

		jsonData, err := json.Marshal(e)

		et := L.NewTable()
		et.RawSetString("expense", lua.LString(jsonData))
		tbl.Append(et)
	}

	L.Push(tbl)
	return 1
}
