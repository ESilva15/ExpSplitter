package models

import (
	"context"
	"encoding/json"
	repo "expenses/expenses/db/repository"

	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
)

type ExpensePayment struct {
	ExpPaymID   int64
	User        User
	PayedAmount decimal.Decimal
}

// PaymentFromJSON takes []byte and returns an *ExpensePayment
func PaymentFromJSON(data []byte) (*ExpensePayment, error) {
	var payment ExpensePayment
	err := json.Unmarshal(data, &payment)
	return &payment, err
}

func (e *Expense) GetPayments(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	payments, err := queries.GetPayments(ctx, e.ExpID)
	if err != nil {
		return err
	}

	e.Payments = mapRepoGetPaymentsRows(payments)
	return nil
}

func GetExpensePaymentByUserID(tx *sql.Tx, eId int64, uId int64,
) (ExpensePayment, error) {
	ctx := context.Background()

	queries := repo.New(tx)
	payment, err := queries.GetExpensePaymentByUser(ctx, repo.GetExpensePaymentByUserParams{
		ExpID:  eId,
		UserID: uId,
	})
	if err != nil {
		return ExpensePayment{}, err
	}

	return mapRepoGetPaymentRow(payment), nil
}

func (p *ExpensePayment) Insert(tx *sql.Tx, expID int64) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.InsertPayment(ctx, repo.InsertPaymentParams{
		ExpID:  expID,
		UserID: p.User.UserID,
		Payed:  p.PayedAmount.String(),
	})
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (p *ExpensePayment) Update(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.UpdatePayment(ctx, repo.UpdatePaymentParams{
		ExpPaymID: p.ExpPaymID,
		Payed:     p.PayedAmount.String(),
		UserID:    p.User.UserID,
	})
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (p *ExpensePayment) Delete(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.DeletePayment(ctx, p.ExpPaymID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}
