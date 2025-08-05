package expenses

import (
	repo "expenses/expenses/db/repository"
)

type Expense struct {
	ExpID        int64
	Description  string
	Value        float64
	Store        Store
	Type         Type
	Category     Category
	Owner        User
	Date         int64
	Payments     []ExpensePayment
	Shares       []ExpenseShare
	CreationDate int64
}

func NewExpense() Expense {
	return Expense{
		ExpID:        -1,
		Description:  "",
		Value:        0.0,
		Store:        NewStore(),
		Category:     NewCategory(),
		Owner:        NewUser(),
		Date:         0,
		Payments:     []ExpensePayment{},
		Shares:       []ExpenseShare{},
		CreationDate: 0,
	}
}

func mapRepoGetExpenseRow(e repo.GetExpenseRow) Expense {
	return Expense{
		ExpID:       e.Expense.ExpID,
		Description: e.Expense.Description,
		Value:       e.Expense.Value,
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
	return Expense{
		ExpID:       e.Expense.ExpID,
		Description: e.Expense.Description,
		Value:       e.Expense.Value,
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
