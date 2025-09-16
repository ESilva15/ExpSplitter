package repo

import (
	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
)

func mapRepoGetPaymentsRow(s pgsqlc.GetPaymentsRow) mod.Payment {
	payed := pgNumericToDecimal(s.ExpensesPayment.Payed)

	return mod.Payment{
		ExpPaymID:   s.ExpensesPayment.ExpPaymID,
		PayedAmount: payed,
		User: mod.User{
			UserID:   s.User.UserID,
			UserName: s.User.UserName,
		},
	}
}

func mapRepoGetPaymentsRows(ep []pgsqlc.GetPaymentsRow) []mod.Payment {
	payments := make([]mod.Payment, len(ep))
	for k, exp := range ep {
		payments[k] = mapRepoGetPaymentsRow(exp)
	}
	return payments
}

func mapRepoGetPaymentRow(s pgsqlc.GetExpensePaymentByUserRow) mod.Payment {
	payed := pgNumericToDecimal(s.ExpensesPayment.Payed)

	return mod.Payment{
		ExpPaymID:   s.ExpensesPayment.ExpPaymID,
		PayedAmount: payed,
		User: mod.User{
			UserID:   s.User.UserID,
			UserName: s.User.UserName,
		},
	}
}
