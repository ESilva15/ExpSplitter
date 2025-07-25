package expenses

type Debt struct {
	Debtor User
	Sum    float32
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

func (e *Expense) CalculateDebts() ([]Debt, error) {
	dc := NewDebtCalculator(e)
	dc.mapShares()
	dc.mapPayments()

	debts := dc.getDebts()

	return debts, nil
}
