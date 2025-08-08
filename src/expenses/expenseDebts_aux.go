package expenses

import (
	mod "expenses/expenses/models"
)

type DebtCalculator struct {
	Shares   map[mod.User]float64
	Payments map[mod.User]float64
	Expense  *mod.Expense
}

func NewDebtCalculator(e *mod.Expense) *DebtCalculator {
	return &DebtCalculator{
		Shares:   make(map[mod.User]float64),
		Payments: make(map[mod.User]float64),
		Expense:  e,
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
