package expenses

import (
	mod "expenses/expenses/models"
	dec "github.com/shopspring/decimal"
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

		payed, err := dec.NewFromString(values[k])
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

func (a *ExpensesApp) GetExpensePaymentByUserID(eId int64, uId int64,
) (mod.ExpensePayment, error) {
	tx, err := a.DB.Begin()
	if err != nil {
		return mod.ExpensePayment{}, err
	}
	defer tx.Rollback()

	payment, err := mod.GetExpensePaymentByUserID(tx, eId, uId)

	return payment, tx.Commit()
}

// insertPayment allows us to insert a payment manually
// for now its private, need to figure out if it needs to be public
func (a *ExpensesApp) insertPayment(payment mod.ExpensePayment, eIdx int64) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = payment.Insert(tx, eIdx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *ExpensesApp) DeletePayment(id int64) error {
	tx, err := a.DB.Begin()
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

func (a *ExpensesApp) UpdatePayment(payment mod.ExpensePayment) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = payment.Update(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *ExpensesApp) AddPayment(expID int64, userID int64, sum dec.Decimal) error {
	payment := mod.ExpensePayment{
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
