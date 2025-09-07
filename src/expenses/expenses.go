package expenses

import (
	"context"
	mod "expenses/expenses/models"
	"github.com/shopspring/decimal"
	"log"
	"time"
)

func (a *ExpensesApp) GetAllExpenses() ([]mod.Expense, error) {
	ctx := context.Background()

	expenses, err := a.ExpRepo.GetAll(ctx)
	if err != nil {
		return []mod.Expense{}, err
	}

	return expenses, nil
}

func (a *ExpensesApp) GetExpensesRanged(startDate string, endDate string) ([]mod.Expense, error) {
	startDateTime, err := time.ParseInLocation("02-Jan-2006 15:04:05", startDate, time.UTC)
	if err != nil {
		log.Printf("error startDate: %v", err)
		return []mod.Expense{}, nil
	}

	endDateTime, err := time.ParseInLocation("02-Jan-2006 15:04:05", endDate, time.UTC)
	if err != nil {
		log.Printf("error endDate: %v", err)
		return []mod.Expense{}, nil
	}

	ctx := context.Background()

	return a.ExpRepo.GetExpensesRange(ctx, startDateTime, endDateTime)
}

func (a *ExpensesApp) GetExpense(id int32) (mod.Expense, error) {
	ctx := context.Background()

	expense, err := a.ExpRepo.Get(ctx, id)
	if err != nil {
		return mod.Expense{}, err
	}

	return expense, nil
}

func (a *ExpensesApp) LoadExpenseShares(e *mod.Expense) error {
	ctx := context.Background()

	var err error
	e.Shares, err = a.ExpRepo.GetShares(ctx, e.ExpID)

	return err
}

func (a *ExpensesApp) LoadExpensePayments(e *mod.Expense) error {
	ctx := context.Background()

	var err error
	e.Payments, err = a.ExpRepo.GetPayments(ctx, e.ExpID)

	return err
}

func (a *ExpensesApp) LoadExpenseDebts(e *mod.Expense) error {
	debts, _ := CalculateDebts(e)
	e.Debts = debts

	return nil
}

func (a *ExpensesApp) DeleteExpense(id int32) error {
	ctx := context.Background()
	err := a.ExpRepo.Delete(ctx, id)
	return err
}

func mapShares(e *mod.Expense) map[mod.User]decimal.Decimal {
	shares := make(map[mod.User]decimal.Decimal)
	for _, share := range e.Shares {
		shares[share.User] = share.Calculated
	}

	return shares
}

func mapPayments(e *mod.Expense) map[mod.User]decimal.Decimal {
	payments := make(map[mod.User]decimal.Decimal)
	for _, p := range e.Payments {
		payments[p.User] = payments[p.User].Add(p.PayedAmount)
	}

	return payments
}

func ExpenseTotalPayed(exp *mod.Expense) decimal.Decimal {
	total := decimal.NewFromFloat(0.0)
	for _, p := range exp.Payments {
		total = total.Add(p.PayedAmount)
	}

	return total
}

func ExpenseIsEvenlyShared(exp *mod.Expense) bool {
	shares := mapShares(exp)
	payments := mapPayments(exp)

	for user, share := range shares {
		val, userHasPayment := payments[user]

		// If a user doesn't even have a payment but has a share, its not even
		if !userHasPayment {
			return false
		}

		if !val.Truncate(2).Equal(share.Truncate(2)) {
			return false
		}
	}

	return true
}

func (a *ExpensesApp) analyzeExpense(e *mod.Expense) {
	// Figure out if its paid off or not by adding the existing payments
	e.PaidOff = e.Value.Equal(ExpenseTotalPayed(e))

	// Figure out if its evenly shared by the people associated to it
	e.SharesEven = ExpenseIsEvenlyShared(e)

	// Update the calculated fields on the shares
	a.NormalizeShares(e)
}

func (a *ExpensesApp) NewExpense(exp mod.Expense) error {
	ctx := context.Background()

	a.analyzeExpense(&exp)

	return a.ExpRepo.Insert(ctx, exp)
}

func (a *ExpensesApp) UpdateExpense(exp mod.Expense) error {
	ctx := context.Background()

	// TODO
	// Move these functions to the models
	a.analyzeExpense(&exp)

	return a.ExpRepo.Update(ctx, exp)
}
