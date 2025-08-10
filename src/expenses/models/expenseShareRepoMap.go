package models

import (
	repo "expenses/expenses/db/repository"

	"github.com/shopspring/decimal"
)

func mapRepoGetSharesRow(s repo.GetSharesRow) ExpenseShare {
	share, _ := decimal.NewFromString(s.ExpensesShare.Share)

	return ExpenseShare{
		ExpShareID: s.ExpensesShare.ExpShareID,
		Share:      share,
		User: User{
			UserID:   s.User.UserID,
			UserName: s.User.UserName,
		},
	}
}

func mapRepoGetSharesRows(es []repo.GetSharesRow) []ExpenseShare {
	shares := make([]ExpenseShare, len(es))
	for k, exp := range es {
		shares[k] = mapRepoGetSharesRow(exp)
	}
	return shares
}
