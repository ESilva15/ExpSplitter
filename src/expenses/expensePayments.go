package expenses

import (
	mod "expenses/expenses/models"

	"github.com/shopspring/decimal"
)

// TODO
// Do not pass an array of strings here, have a separate function to do the
// conversion that the user calls before calling this
// Can even create a struct to then hold the data
func ParseFormPayments(userIDs []string, paymentsIDs []string,
	values []string) ([]mod.ExpensePayment, error) {
	payments := []mod.ExpensePayment{}
	for k := range userIDs {
		userID, err := ParseID(userIDs[k])
		if err != nil {
			return nil, err
		}

		payed, err := decimal.NewFromString(values[k])
		if err != nil {
			return nil, err
		}

		newPayment := mod.ExpensePayment{
			ExpPaymID: -1,
			User: mod.User{
				UserID: userID,
			},
			PayedAmount: payed,
		}

		if paymentsIDs[k] != "" {
			id, err := ParseID(paymentsIDs[k])
			if err != nil {
				return nil, err
			}
			newPayment.ExpPaymID = id
		}

		payments = append(payments, newPayment)
	}

	return payments, nil
}

func (s *Service) DeletePayment(id int64) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	payment := mod.ExpensePayment{
		ExpPaymID: id,
	}

	err = payment.Delete(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
