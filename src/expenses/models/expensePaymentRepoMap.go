package models

import (
	repo "expenses/expenses/db/repository"

	"github.com/shopspring/decimal"
)

func mapRepoGetPaymentsRow(s repo.GetPaymentsRow) ExpensePayment {
	payed, _ := decimal.NewFromString(s.ExpensesPayment.Payed)

	return ExpensePayment{
		ExpPaymID:   s.ExpensesPayment.ExpPaymID,
		PayedAmount: payed,
		User: User{
			UserID:   s.User.UserID,
			UserName: s.User.UserName,
		},
	}
}

func mapRepoGetPaymentsRows(ep []repo.GetPaymentsRow) []ExpensePayment {
	payments := make([]ExpensePayment, len(ep))
	for k, exp := range ep {
		payments[k] = mapRepoGetPaymentsRow(exp)
	}
	return payments
}

func mapRepoGetPaymentRow(s repo.GetExpensePaymentByUserRow) ExpensePayment {
	payed, _ := decimal.NewFromString(s.ExpensesPayment.Payed)

	return ExpensePayment{
		ExpPaymID:   s.ExpensesPayment.ExpPaymID,
		PayedAmount: payed,
		User: User{
			UserID:   s.User.UserID,
			UserName: s.User.UserName,
		},
	}
}
