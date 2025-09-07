package expenses

import (
	"context"
	mod "expenses/expenses/models"

	"github.com/jackc/pgx/v5"
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

func (a *ExpensesApp) GetExpensePaymentByUserID(eId int32, uId int32,
) (mod.Payment, error) {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return mod.Payment{}, err
	}
	defer tx.Rollback(ctx)

	payment, err := mod.GetExpensePaymentByUserID(a.DB, tx, eId, uId)

	return payment, tx.Commit(ctx)
}

// insertPayment allows us to insert a payment manually
// for now its private, need to figure out if it needs to be public
func (a *ExpensesApp) insertPayment(payment mod.Payment, eIdx int32) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = payment.Insert(a.DB, tx, eIdx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (a *ExpensesApp) DeletePayment(id int32) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	payment := mod.Payment{
		ExpPaymID: id,
	}

	err = payment.Delete(a.DB, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (a *ExpensesApp) UpdatePayment(payment mod.Payment) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = payment.Update(a.DB, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (a *ExpensesApp) AddPayment(expID int32, userID int32, sum dec.Decimal) error {
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
