package models

import (
	"context"
	"encoding/json"
	"expenses/expenses/db/repository"
	repo "expenses/expenses/db/repository"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/mattn/go-sqlite3"
	dec "github.com/shopspring/decimal"
)

type Share struct {
	ExpShareID int32       `json:"ExpShareID"`
	User       User        `json:"User"`
	Share      dec.Decimal `json:"Share"`
	Calculated dec.Decimal `json:"Calculated"`
}
type Shares []Share

// ShareFromJSON takes []byte and returns an *Share
func ShareFromJSON(data []byte) (*Share, error) {
	var share Share

	err := json.Unmarshal(data, &share)

	return &share, err
}

func (sh Shares) Equal(other Shares) bool {
	if len(sh) != len(other) {
		return false
	}

	for k := range sh {
		if sh[k].User != other[k].User ||
			!sh[k].Calculated.Equal(other[k].Calculated) ||
			!sh[k].Share.Equal(other[k].Share) {
			return false
		}
	}

	return true
}

func (e *Expense) GetShares(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repository.New(db).WithTx(tx)
	shares, err := queries.GetShares(ctx, e.ExpID)
	if err != nil {
		return err
	}

	e.Shares = mapRepoGetSharesRows(shares)
	return nil
}

func (sh *Share) Insert(db repo.DBTX, tx pgx.Tx, expID int32) error {
	ctx := context.Background()

	share, err := decimalToNumeric(sh.Share)
	if err != nil {
		return err
	}
	calculated, err := decimalToNumeric(sh.Calculated)
	if err != nil {
		return err
	}

	queries := repo.New(db).WithTx(tx)
	res, err := queries.InsertShare(ctx, repo.InsertShareParams{
		ExpID:      expID,
		Share:      share,
		UserID:     sh.User.UserID,
		Calculated: calculated,
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

func (sh *Share) Update(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	share, err := decimalToNumeric(sh.Share)
	if err != nil {
		return err
	}
	calculated, err := decimalToNumeric(sh.Calculated)
	if err != nil {
		return err
	}

	queries := repo.New(db).WithTx(tx)
	res, err := queries.UpdateShare(ctx, repo.UpdateShareParams{
		ExpShareID: sh.ExpShareID,
		Share:      share,
		UserID:     sh.User.UserID,
		Calculated: calculated,
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

func (sh *Share) Delete(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	res, err := queries.DeleteShare(ctx, sh.ExpShareID)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}
