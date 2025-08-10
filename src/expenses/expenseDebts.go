package expenses

import (
	mod "expenses/expenses/models"

	"github.com/shopspring/decimal"
)

type Debt struct {
	Debtor mod.User
	Sum    decimal.Decimal
}

func sortBySum(a, b Debt) int {
	if a.Sum.LessThan(b.Sum) {
		return -1
	}

	if a.Sum.GreaterThan(b.Sum) {
		return 1
	}

	return 0
}

func CalculateDebts(e *mod.Expense) ([]Debt, error) {
	dc := NewDebtCalculator(e)
	dc.mapShares()
	dc.mapPayments()

	debts := dc.getDebts()

	return debts, nil
}
