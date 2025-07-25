package expenses

type DebtCalculator struct {
	Shares    map[User]float32
	Payments  map[User]float32
	Expense   *Expense
}

func NewDebtCalculator(e *Expense) *DebtCalculator {
	return &DebtCalculator{
		Shares:    make(map[User]float32),
		Payments:  make(map[User]float32),
		Expense:   e,
	}
}

func (dc *DebtCalculator) mapShares() {
	for _, share := range dc.Expense.Shares {
		dc.Shares[share.User] = share.Share
		dc.Payments[share.User] = 0
	}
}

func (dc *DebtCalculator) mapPayments() {
	for _, payment := range dc.Expense.Payments {
		dc.Payments[payment.User] += payment.PayedAmount
	}
}

func (dc *DebtCalculator) getDebts() []Debt {
	debts := []Debt{}	

	for user := range dc.Payments {
		debt := (dc.Shares[user] * dc.Expense.Value) - dc.Payments[user]

		if debt > 0 {
			debts = append(debts, Debt{Debtor: user, Sum: debt})
		}
	}

	return debts
}
