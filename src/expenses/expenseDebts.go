package expenses

import (
	mod "expenses/expenses/models"

	dec "github.com/shopspring/decimal"
)

func sortBySum(a, b mod.Debt) int {
	if a.Sum.LessThan(b.Sum) {
		return -1
	}

	if a.Sum.GreaterThan(b.Sum) {
		return 1
	}

	return 0
}

func filterExpenseParticipants(e *mod.Expense,
) (map[mod.User]dec.Decimal, map[mod.User]dec.Decimal) {
	shares := mapShares(e)
	payments := mapPayments(e)

	debtors := make(map[mod.User]dec.Decimal)
	creditors := make(map[mod.User]dec.Decimal)

	for user, share := range shares {
		debt := (share.Mul(e.Value)).Sub(payments[user])
		if debt.LessThan(dec.NewFromFloat(0.0)) {
			creditors[user] = debt.Abs()
		}

		if debt.GreaterThan(dec.NewFromFloat(0.0)) {
			debtors[user] = debt.Abs()
		}
	}

	return debtors, creditors
}

func resolveDebts(debtors map[mod.User]dec.Decimal,
	creditors map[mod.User]dec.Decimal) []mod.Debt {
	debts := []mod.Debt{}

	// maybe create a map point to the debt of each user and count from there?
	// I should sketch this one first

	return debts
}

func CalculateDebts(e *mod.Expense) ([]mod.Debt, error) {
	// dc := NewDebtCalculator(e)
	// dc.mapShares()
	// dc.mapPayments()
	//
	// debts := dc.getDebts()

	debtors, creditors := filterExpenseParticipants(e)
	debts := resolveDebts(debtors, creditors)

	return []mod.Debt{}, nil
}
