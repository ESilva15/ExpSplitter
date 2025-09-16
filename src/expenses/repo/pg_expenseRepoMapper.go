package repo

import (
	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
)

type ExpenseRow interface {
	GetExpense() pgsqlc.Expense
	GetStore() pgsqlc.Store
	GetType() pgsqlc.ExpenseType
	GetCategory() pgsqlc.Category
	GetUser() pgsqlc.User
}

type (
	ExpenseRowSingle pgsqlc.GetExpenseRow
	ExpenseRowMulti  pgsqlc.GetExpensesRow
)

func (e ExpenseRowSingle) GetExpense() pgsqlc.Expense   { return e.Expense }
func (e ExpenseRowSingle) GetStore() pgsqlc.Store       { return e.Store }
func (e ExpenseRowSingle) GetType() pgsqlc.ExpenseType  { return e.ExpenseType }
func (e ExpenseRowSingle) GetCategory() pgsqlc.Category { return e.Category }
func (e ExpenseRowSingle) GetUser() pgsqlc.User         { return e.User }

func (e ExpenseRowMulti) GetExpense() pgsqlc.Expense   { return e.Expense }
func (e ExpenseRowMulti) GetStore() pgsqlc.Store       { return e.Store }
func (e ExpenseRowMulti) GetType() pgsqlc.ExpenseType  { return e.ExpenseType }
func (e ExpenseRowMulti) GetCategory() pgsqlc.Category { return e.Category }
func (e ExpenseRowMulti) GetUser() pgsqlc.User         { return e.User }

func mapRepoExpenseRow(e ExpenseRow) mod.Expense {
	exp := e.GetExpense()
	value := pgNumericToDecimal(exp.Value)
	paidOff := pgBoolToBool(exp.PaidOff)
	sharesEven := pgBoolToBool(exp.SharesEven)
	expDate := pgTimestampToTime(exp.ExpDate)
	creationDate := pgTimestampToTime(exp.CreationDate)

	return mod.Expense{
		ExpID:       exp.ExpID,
		Description: exp.Description,
		Value:       value,
		Store: mod.Store{
			StoreID:   e.GetStore().StoreID,
			StoreName: e.GetStore().StoreName,
		},
		Type: mod.Type{
			TypeID:   e.GetType().TypeID,
			TypeName: e.GetType().TypeName,
		},
		Category: mod.Category{
			CategoryID:   e.GetCategory().CategoryID,
			CategoryName: e.GetCategory().CategoryName,
		},
		Owner: mod.User{
			UserID:   e.GetUser().UserID,
			UserName: e.GetUser().UserName,
		},
		Date:         expDate,
		Shares:       []mod.Share{},
		Payments:     []mod.Payment{},
		PaidOff:      paidOff,
		SharesEven:   sharesEven,
		QRString:     exp.Qr,
		CreationDate: creationDate,
	}
}

func mapRepoGetExpensesRows(er []pgsqlc.GetExpensesRow) []mod.Expense {
	expenses := make([]mod.Expense, len(er))
	for k, exp := range er {
		expenses[k] = mapRepoExpenseRow(ExpenseRowSingle(exp))
	}
	return expenses
}
