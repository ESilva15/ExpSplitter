package expenses

import (
	"expenses/config"
	mod "expenses/expenses/models"

	dec "github.com/shopspring/decimal"
)

// TODO
// Do not pass an array of strings here, have a separate function to do the
// conversion that the user calls before calling this
// Can even create a struct to then hold the data
func ParseFormShares(userIDs []string, shares []string, sharesIDs []string,
) ([]mod.Share, error) {
	shareList := []mod.Share{}
	for i := range userIDs {
		userID, err := ParseID(userIDs[i])
		if err != nil {
			return nil, err
		}

		share, err := dec.NewFromString(shares[i])
		if err != nil {
			return nil, err
		}

		newShare := mod.Share{
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

// NormalizeShares will take the total of an expense and the proposed shares
// and calculate how much each user actually has to pay - avoids fracd cents
func (a *ExpensesApp) NormalizeShares(e *mod.Expense) error {
	cfg := config.GetInstance()
	affi := cfg.CurrencyAffinity

	excess := dec.NewFromFloat(0.0)
	ownerShIdx := -1
	for k := range e.Shares {
		if e.Shares[k].User.UserID == e.Owner.UserID {
			ownerShIdx = k
		}
		owed := e.Value.Mul(e.Shares[k].Share)
		excess = excess.Add(owed.Sub(owed.Truncate(affi)))

		e.Shares[k].Calculated = owed.Truncate(affi)
	}

	// It fails here for some reason
	e.Shares[ownerShIdx].Calculated = e.Shares[ownerShIdx].Calculated.Add(excess)

	return nil
}

// insertShare allows us to insert a share manually
// for now its private, need to figure out if it needs to be public
func (a *ExpensesApp) insertShare(share *mod.Share, eIdx int64) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = share.Insert(tx, eIdx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *ExpensesApp) DeleteShare(id int64) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	share := mod.Share{
		ExpShareID: id,
	}

	err = share.Delete(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
