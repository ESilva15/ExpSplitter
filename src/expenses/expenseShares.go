package expenses

import (
	"strconv"

	mod "expenses/expenses/models"
)

// TODO
// Do not pass an array of strings here, have a separate function to do the
// conversion that the user calls before calling this
// Can even create a struct to then hold the data
func ParseFormShares(userIDs []string, shares []string, sharesIDs []string,
) ([]mod.ExpenseShare, error) {
	shareList := []mod.ExpenseShare{}
	for i := range userIDs {
		userID, err := strconv.ParseInt(userIDs[i], 10, 16)
		if err != nil {
			return nil, err
		}

		share, err := strconv.ParseFloat(shares[i], 32)
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
			id, err := strconv.ParseInt(sharesIDs[i], 10, 16)
			if err != nil {
				return nil, err
			}
			newShare.ExpShareID = id
		}

		shareList = append(shareList, newShare)
	}

	return shareList, nil
}
