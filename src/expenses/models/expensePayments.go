package models

import (
	"context"
	"expenses/config"
	repo "expenses/expenses/db/repository"

	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type ExpensePayment struct {
	ExpPaymID   int64
	User        User
	PayedAmount float64
}

func (e *Expense) GetPayments() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	payments, err := queries.GetPayments(ctx, e.ExpID)
	if err != nil {
		return err
	}

	e.Payments = mapRepoGetPaymentsRows(payments)
	return nil
}

func (p *ExpensePayment) Insert(expID int64) error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	res, err := queries.InsertPayment(ctx, repo.InsertPaymentParams{
		ExpID:  expID,
		UserID: p.User.UserID,
		Payed:  p.PayedAmount,
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

func (p *ExpensePayment) Update() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	res, err := queries.UpdatePayment(ctx, repo.UpdatePaymentParams{
		ExpPaymID: p.ExpPaymID,
		Payed:     p.PayedAmount,
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

func (p *ExpensePayment) Delete() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
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
