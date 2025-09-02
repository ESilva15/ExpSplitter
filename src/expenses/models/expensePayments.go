package models

import (
	"context"
	"encoding/json"
	repo "expenses/expenses/db/repository"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
)

type ExpensePayment struct {
	ExpPaymID   int32
	User        User
	PayedAmount decimal.Decimal
}

// PaymentFromJSON takes []byte and returns an *ExpensePayment
func PaymentFromJSON(data []byte) (*ExpensePayment, error) {
	var payment ExpensePayment
	err := json.Unmarshal(data, &payment)
	return &payment, err
}

func (e *Expense) GetPayments(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	payments, err := queries.GetPayments(ctx, e.ExpID)
	if err != nil {
		return err
	}

	e.Payments = mapRepoGetPaymentsRows(payments)
	return nil
}

func GetExpensePaymentByUserID(db repo.DBTX, tx pgx.Tx, eId int32, uId int32,
) (ExpensePayment, error) {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	payment, err := queries.GetExpensePaymentByUser(ctx, repo.GetExpensePaymentByUserParams{
		ExpID:  eId,
		UserID: uId,
	})
	if err != nil {
		return ExpensePayment{}, err
	}

	return mapRepoGetPaymentRow(payment), nil
}

func (p *ExpensePayment) Insert(db repo.DBTX, tx pgx.Tx, expID int32) error {
	ctx := context.Background()

	payed, err := decimalToNumeric(p.PayedAmount)
	if err != nil {
		return err
	}

	queries := repo.New(db).WithTx(tx)
	res, err := queries.InsertPayment(ctx, repo.InsertPaymentParams{
		ExpID:  expID,
		UserID: p.User.UserID,
		Payed:  payed,
	})
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (p *ExpensePayment) Update(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	payed, err := decimalToNumeric(p.PayedAmount)
	if err != nil {
		return err
	}

	queries := repo.New(db).WithTx(tx)
	res, err := queries.UpdatePayment(ctx, repo.UpdatePaymentParams{
		ExpPaymID: p.ExpPaymID,
		Payed:     payed,
		UserID:    p.User.UserID,
	})
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (p *ExpensePayment) Delete(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	res, err := queries.DeletePayment(ctx, p.ExpPaymID)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}
