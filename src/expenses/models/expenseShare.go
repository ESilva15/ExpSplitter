package models

import (
	"context"
	"expenses/expenses/db/repository"
	repo "expenses/expenses/db/repository"

	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shopspring/decimal"
)

type ExpenseShare struct {
	ExpShareID int64
	User       User
	Share      decimal.Decimal
}

func (e *Expense) GetShares(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repository.New(tx)
	shares, err := queries.GetShares(ctx, e.ExpID)
	if err != nil {
		return err
	}

	e.Shares = mapRepoGetSharesRows(shares)
	return nil
}

func (sh *ExpenseShare) Insert(tx *sql.Tx, expID int64) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.InsertShare(ctx, repo.InsertShareParams{
		ExpID:  expID,
		Share:  sh.Share.String(),
		UserID: sh.User.UserID,
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

func (sh *ExpenseShare) Update(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.UpdateShare(ctx, repo.UpdateShareParams{
		ExpShareID: sh.ExpShareID,
		Share:      sh.Share.String(),
		UserID:     sh.User.UserID,
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

func (sh *ExpenseShare) Delete(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.DeleteShare(ctx, sh.ExpShareID)
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
