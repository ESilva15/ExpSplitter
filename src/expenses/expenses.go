package expenses

import (
	mod "expenses/expenses/models"
	"log"
	"time"

	"github.com/shopspring/decimal"
)

func (a *ExpensesApp) GetAllExpenses() ([]mod.Expense, error) {
	tx, err := a.DB.Begin()
	if err != nil {
		return []mod.Expense{}, err
	}
	defer tx.Rollback()

	expenses, err := mod.GetAllExpenses(tx)

	return expenses, tx.Commit()
}

func (a *ExpensesApp) GetExpensesRanged(startDate string, endDate string) ([]mod.Expense, error) {
	startDateTime, err := time.ParseInLocation("02-Jan-2006 15:04:05", startDate, time.UTC)
	if err != nil {
		log.Printf("error startDate: %v", err)
		return []mod.Expense{}, nil
	}
	start := startDateTime.Unix()

	endDateTime, err := time.ParseInLocation("02-Jan-2006 15:04:05", endDate, time.UTC)
	if err != nil {
		log.Printf("error endDate: %v", err)
		return []mod.Expense{}, nil
	}
	end := endDateTime.Unix()

	tx, err := a.DB.Begin()
	if err != nil {
		return []mod.Expense{}, err
	}
	defer tx.Rollback()

	expenses, err := mod.GetExpensesRange(tx, start, end)

	return expenses, tx.Commit()
}

func (a *ExpensesApp) GetExpense(id int64) (mod.Expense, error) {
	tx, err := a.DB.Begin()
	if err != nil {
		return mod.Expense{}, err
	}
	defer tx.Rollback()

	expense, err := mod.GetExpense(tx, id)
	if err != nil {
		return mod.Expense{}, err
	}

	return expense, tx.Commit()
}

func (a *ExpensesApp) LoadExpenseShares(e *mod.Expense) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = e.GetShares(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *ExpensesApp) LoadExpensePayments(e *mod.Expense) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = e.GetPayments(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *ExpensesApp) LoadExpenseDebts(e *mod.Expense) error {
	debts, _ := CalculateDebts(e)
	e.Debts = debts

	return nil
}

func (a *ExpensesApp) DeleteExpense(id int64) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	expense := mod.Expense{
		ExpID: id,
	}

	err = expense.Delete()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func mapShares(e *mod.Expense) map[mod.User]decimal.Decimal {
	shares := make(map[mod.User]decimal.Decimal)
	for _, share := range e.Shares {
		shares[share.User] = share.Share
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
	log.Printf("is exp %d even?", exp.ExpID)

	shares := mapShares(exp)
	payments := mapPayments(exp)

	for user, share := range shares {
		userShare := share.Mul(exp.Value)
		val, userHasPayment := payments[user]

		// If a user doesn't even have a payment but has a share, its not even
		if !userHasPayment {
			return false
		}

		if !val.Truncate(2).Equal(userShare.Truncate(2)) {
			log.Printf("No: %s != %s", val.Truncate(2), userShare.Truncate(2))
			return false
		}
	}

	log.Printf("Yes")

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
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	a.analyzeExpense(&exp)

	err = exp.Insert(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *ExpensesApp) UpdateExpense(exp mod.Expense) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	a.analyzeExpense(&exp)

	err = exp.Update(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
