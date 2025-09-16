package repo

import (
	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
)

func mapRepoGetSharesRow(s pgsqlc.GetSharesRow) mod.Share {
	share := pgNumericToDecimal(s.ExpensesShare.Share)
	calculated := pgNumericToDecimal(s.ExpensesShare.Calculated)

	return mod.Share{
		ExpShareID: s.ExpensesShare.ExpShareID,
		Share:      share,
		User: mod.User{
			UserID:   s.User.UserID,
			UserName: s.User.UserName,
		},
		Calculated: calculated,
	}
}

func mapRepoGetSharesRows(es []pgsqlc.GetSharesRow) mod.Shares {
	shares := make(mod.Shares, len(es))
	for k, exp := range es {
		shares[k] = mapRepoGetSharesRow(exp)
	}
	return shares
}
