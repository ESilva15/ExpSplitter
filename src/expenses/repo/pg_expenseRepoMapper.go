package repo

import (
	mod "expenses/expenses/models"
	"expenses/expenses/repo/pgdb/pgsqlc"
)

func mapRepoGetExpenseRow(e pgsqlc.GetExpenseRow) mod.Expense {
	value := pgNumericToDecimal(e.Expense.Value)
	paidOff := pgBoolToBool(e.Expense.PaidOff)
	sharesEven := pgBoolToBool(e.Expense.SharesEven)
	expDate := pgTimestampToTime(e.Expense.ExpDate)
	creationDate := pgTimestampToTime(e.Expense.CreationDate)

	return mod.Expense{
		ExpID:       e.Expense.ExpID,
		Description: e.Expense.Description,
		Value:       value,
		Store: mod.Store{
			StoreID:   e.Store.StoreID,
			StoreName: e.Store.StoreName,
		},
		Type: mod.Type{
			TypeID:   e.ExpenseType.TypeID,
			TypeName: e.ExpenseType.TypeName,
		},
		Category: mod.Category{
			CategoryID:   e.Category.CategoryID,
			CategoryName: e.Category.CategoryName,
		},
		Owner: mod.User{
			UserID:   e.User.UserID,
			UserName: e.User.UserName,
		},
		Date:         expDate,
		Shares:       []mod.Share{},
		Payments:     []mod.Payment{},
		PaidOff:      paidOff,
		SharesEven:   sharesEven,
		QRString:     e.Expense.Qr,
		CreationDate: creationDate,
	}
}

func mapRepoGetExpenseRowMulti(e pgsqlc.GetExpensesRow) mod.Expense {
	value := pgNumericToDecimal(e.Expense.Value)
	paidOff := pgBoolToBool(e.Expense.PaidOff)
	sharesEven := pgBoolToBool(e.Expense.SharesEven)
	expDate := pgTimestampToTime(e.Expense.ExpDate)
	creationDate := pgTimestampToTime(e.Expense.CreationDate)

	return mod.Expense{
		ExpID:       e.Expense.ExpID,
		Description: e.Expense.Description,
		Value:       value,
		Store: mod.Store{
			StoreID:   e.Store.StoreID,
			StoreName: e.Store.StoreName,
		},
		Type: mod.Type{
			TypeID:   e.ExpenseType.TypeID,
			TypeName: e.ExpenseType.TypeName,
		},
		Category: mod.Category{
			CategoryID:   e.Category.CategoryID,
			CategoryName: e.Category.CategoryName,
		},
		Owner: mod.User{
			UserID:   e.User.UserID,
			UserName: e.User.UserName,
		},
		Date:         expDate,
		Shares:       []mod.Share{},
		Payments:     []mod.Payment{},
		PaidOff:      paidOff,
		SharesEven:   sharesEven,
		QRString:     e.Expense.Qr,
		CreationDate: creationDate,
	}
}

func mapRepoGetExpensesRows(er []pgsqlc.GetExpensesRow) []mod.Expense {
	expenses := make([]mod.Expense, len(er))
	for k, exp := range er {
		expenses[k] = mapRepoGetExpenseRowMulti(exp)
	}
	return expenses
}
