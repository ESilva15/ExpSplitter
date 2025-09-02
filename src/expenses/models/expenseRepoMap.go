package models

import (
	repo "expenses/expenses/db/repository"
)

func mapRepoGetExpenseRow(e repo.GetExpenseRow) Expense {
	value := pgNumericToDecimal(e.Expense.Value)
	paidOff := pgBoolToBool(e.Expense.PaidOff)
	sharesEven := pgBoolToBool(e.Expense.SharesEven)
	expDate := pgTimestampToTime(e.Expense.ExpDate)
	creationDate := pgTimestampToTime(e.Expense.CreationDate)

	return Expense{
		ExpID:       e.Expense.ExpID,
		Description: e.Expense.Description,
		Value:       value,
		Store: Store{
			StoreID:   e.Store.StoreID,
			StoreName: e.Store.StoreName,
		},
		Type: Type{
			TypeID:   e.ExpenseType.TypeID,
			TypeName: e.ExpenseType.TypeName,
		},
		Category: Category{
			CategoryID:   e.Category.CategoryID,
			CategoryName: e.Category.CategoryName,
		},
		Owner: User{
			UserID:   e.User.UserID,
			UserName: e.User.UserName,
		},
		Date:         expDate,
		Shares:       []Share{},
		Payments:     []ExpensePayment{},
		PaidOff:      paidOff,
		SharesEven:   sharesEven,
		CreationDate: creationDate,
	}
}

func mapRepoGetExpenseRowMulti(e repo.GetExpensesRow) Expense {
	value := pgNumericToDecimal(e.Expense.Value)
	paidOff := pgBoolToBool(e.Expense.PaidOff)
	sharesEven := pgBoolToBool(e.Expense.SharesEven)
	expDate := pgTimestampToTime(e.Expense.ExpDate)
	creationDate := pgTimestampToTime(e.Expense.CreationDate)

	return Expense{
		ExpID:       e.Expense.ExpID,
		Description: e.Expense.Description,
		Value:       value,
		Store: Store{
			StoreID:   e.Store.StoreID,
			StoreName: e.Store.StoreName,
		},
		Type: Type{
			TypeID:   e.ExpenseType.TypeID,
			TypeName: e.ExpenseType.TypeName,
		},
		Category: Category{
			CategoryID:   e.Category.CategoryID,
			CategoryName: e.Category.CategoryName,
		},
		Owner: User{
			UserID:   e.User.UserID,
			UserName: e.User.UserName,
		},
		Date:         expDate,
		Shares:       []Share{},
		Payments:     []ExpensePayment{},
		PaidOff:      paidOff,
		SharesEven:   sharesEven,
		CreationDate: creationDate,
	}
}

func mapRepoGetExpensesRows(er []repo.GetExpensesRow) []Expense {
	expenses := make([]Expense, len(er))
	for k, exp := range er {
		expenses[k] = mapRepoGetExpenseRowMulti(exp)
	}
	return expenses
}
