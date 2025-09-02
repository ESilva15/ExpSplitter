package models

import (
	repo "expenses/expenses/db/repository"
)

func mapRepoGetSharesRow(s repo.GetSharesRow) Share {
	share := pgNumericToDecimal(s.ExpensesShare.Share)
	calculated := pgNumericToDecimal(s.ExpensesShare.Calculated)

	return Share{
		ExpShareID: s.ExpensesShare.ExpShareID,
		Share:      share,
		User: User{
			UserID:   s.User.UserID,
			UserName: s.User.UserName,
		},
		Calculated: calculated,
	}
}

func mapRepoGetSharesRows(es []repo.GetSharesRow) []Share {
	shares := make([]Share, len(es))
	for k, exp := range es {
		shares[k] = mapRepoGetSharesRow(exp)
	}
	return shares
}
