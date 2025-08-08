package models

import repo "expenses/expenses/db/repository"

func mapRepoGetPaymentsRow(s repo.GetPaymentsRow) ExpensePayment {
	return ExpensePayment{
		ExpPaymID:   s.ExpensesPayment.ExpPaymID,
		PayedAmount: s.ExpensesPayment.Payed,
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
