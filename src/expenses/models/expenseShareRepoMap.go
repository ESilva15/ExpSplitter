package models

import repo "expenses/expenses/db/repository"

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
