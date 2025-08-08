package expenses

import (
	mod "expenses/expenses/models"
)

type Debt struct {
	Debtor mod.User
	Sum    float64
}

func sortBySum(a, b Debt) int {
	if a.Sum < b.Sum {
		return -1
	}
	if a.Sum > b.Sum {
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
