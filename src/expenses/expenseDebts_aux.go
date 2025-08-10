package expenses

import (
	mod "expenses/expenses/models"

	"github.com/shopspring/decimal"
)

type DebtCalculator struct {
	Shares   map[mod.User]decimal.Decimal
	Payments map[mod.User]decimal.Decimal
	Expense  *mod.Expense
}

func NewDebtCalculator(e *mod.Expense) *DebtCalculator {
	return &DebtCalculator{
		Shares:   make(map[mod.User]decimal.Decimal),
		Payments: make(map[mod.User]decimal.Decimal),
		Expense:  e,
	}
}

func (dc *DebtCalculator) mapShares() {
	for _, share := range dc.Expense.Shares {
		dc.Shares[share.User] = share.Share
		dc.Payments[share.User] = decimal.NewFromInt(0)
	}
}

func (dc *DebtCalculator) mapPayments() {
	for _, payment := range dc.Expense.Payments {
		dc.Payments[payment.User] = dc.Payments[payment.User].Add(payment.PayedAmount)
	}
}

func (dc *DebtCalculator) getDebts() []Debt {
	debts := []Debt{}

	for user := range dc.Payments {
		// debt := (dc.Shares[user] * dc.Expense.Value) - dc.Payments[user]
		debt := (dc.Shares[user].Mul(dc.Expense.Value)).Sub(dc.Payments[user])

		if debt.GreaterThan(decimal.NewFromInt(0)) {
			debts = append(debts, Debt{Debtor: user, Sum: debt})
		}
	}

	return debts
}
