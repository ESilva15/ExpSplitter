package expenses

import (
	"context"
	mod "expenses/expenses/models"
	dec "github.com/shopspring/decimal"
)

// TODO
// Do not pass an array of strings here, have a separate function to do the
// conversion that the user calls before calling this
// Can even create a struct to then hold the data
func ParseFormPayments(userIDs []string, paymentsIDs []string,
	values []string) ([]mod.Payment, error) {
	payments := []mod.Payment{}
	for k := range userIDs {
		userID, err := ParseID(userIDs[k])
		if err != nil {
			return nil, err
		}

		payed, err := dec.NewFromString(values[k])
		if err != nil {
			return nil, err
		}

		newPayment := mod.Payment{
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

func (a *ExpApp) GetExpensePaymentByUserID(eId int32, uId int32,
) (mod.Payment, error) {
	ctx := context.Background()
	return a.ExpRepo.GetExpensePaymentByUserID(ctx, eId, uId)
}

// insertPayment allows us to insert a payment manually
// for now its private, need to figure out if it needs to be public
func (a *ExpApp) insertPayment(payment mod.Payment, eIdx int32) error {
	ctx := context.Background()
	return a.ExpRepo.InsertPayment(ctx, eIdx, payment)
}

func (a *ExpApp) DeletePayment(id int32) error {
	ctx := context.Background()
	return a.ExpRepo.DeletePayment(ctx, id)
}

func (a *ExpApp) UpdatePayment(payment mod.Payment) error {
	ctx := context.Background()
	return a.ExpRepo.UpdatePayment(ctx, payment)
}

func (a *ExpApp) AddPayment(expID int32, userID int32, sum dec.Decimal) error {
	payment := mod.Payment{
		User: mod.User{
			UserID: userID,
		},
		PayedAmount: sum,
	}

	err := a.insertPayment(payment, expID)
	if err != nil {
		return err
	}

	return nil
}
