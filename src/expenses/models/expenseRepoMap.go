package models

import (
	repo "expenses/expenses/db/repository"

	"github.com/shopspring/decimal"
)

func mapRepoGetExpenseRow(e repo.GetExpenseRow) Expense {
	value, _ := decimal.NewFromString(e.Expense.Value)

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
		Date:         e.Expense.ExpDate,
		Shares:       []ExpenseShare{},
		Payments:     []ExpensePayment{},
		CreationDate: e.Expense.CreationDate,
	}
}

func mapRepoGetExpenseRowMulti(e repo.GetExpensesRow) Expense {
	value, _ := decimal.NewFromString(e.Expense.Value)

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
		Date:         e.Expense.ExpDate,
		Shares:       []ExpenseShare{},
		Payments:     []ExpensePayment{},
		CreationDate: e.Expense.CreationDate,
	}
}

func mapRepoGetExpensesRows(er []repo.GetExpensesRow) []Expense {
	expenses := make([]Expense, len(er))
	for k, exp := range er {
		expenses[k] = mapRepoGetExpenseRowMulti(exp)
	}
	return expenses
}
