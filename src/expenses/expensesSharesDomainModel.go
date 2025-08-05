package expenses

import (
	repo "expenses/expenses/db/repository"
)

type ExpenseShare struct {
	ExpShareID int64
	User       User
	Share      float64
}

func mapRepoGetSharesRow(s repo.GetSharesRow) ExpenseShare {
	return ExpenseShare{
		ExpShareID: s.ExpensesShare.ExpShareID,
		Share:      s.ExpensesShare.Share,
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
