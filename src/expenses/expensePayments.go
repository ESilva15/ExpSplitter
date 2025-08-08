package expenses

import (
	mod "expenses/expenses/models"
	"strconv"
)

// TODO
// Do not pass an array of strings here, have a separate function to do the
// conversion that the user calls before calling this
// Can even create a struct to then hold the data
func ParseFormPayments(userIDs []string, paymentsIDs []string,
	values []string) ([]mod.ExpensePayment, error) {
	payments := []mod.ExpensePayment{}
	for k := range userIDs {
		userID, err := strconv.ParseInt(userIDs[k], 10, 16)
		if err != nil {
			return nil, err
		}

		payed, err := strconv.ParseFloat(values[k], 32)
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
			id, err := strconv.ParseInt(paymentsIDs[k], 10, 16)
			if err != nil {
				return nil, err
			}
			newPayment.ExpPaymID = id
		}

		payments = append(payments, newPayment)
	}

	return payments, nil
}

func DeletePayment(id int64) error {
	payment := mod.ExpensePayment{
		ExpPaymID: id,
	}

	return payment.Delete()
}
