package expenses

import (
	mod "expenses/expenses/models"

	"github.com/shopspring/decimal"
)

// TODO
// Do not pass an array of strings here, have a separate function to do the
// conversion that the user calls before calling this
// Can even create a struct to then hold the data
func ParseFormShares(userIDs []string, shares []string, sharesIDs []string,
) ([]mod.ExpenseShare, error) {
	shareList := []mod.ExpenseShare{}
	for i := range userIDs {
		userID, err := ParseID(userIDs[i])
		if err != nil {
			return nil, err
		}

		share, err := decimal.NewFromString(shares[i])
		if err != nil {
			return nil, err
		}

		newShare := mod.ExpenseShare{
			ExpShareID: -1,
			User: mod.User{
				UserID: userID,
			},
			Share: share,
		}

		if sharesIDs[i] != "" {
			id, err := ParseID(sharesIDs[i])
			if err != nil {
				return nil, err
			}
			newShare.ExpShareID = id
		}

		shareList = append(shareList, newShare)
	}

	return shareList, nil
}
