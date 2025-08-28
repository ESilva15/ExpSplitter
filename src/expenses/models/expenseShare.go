package models

import (
	"context"
	"expenses/expenses/db/repository"
	repo "expenses/expenses/db/repository"
	"database/sql"
	"fmt"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	dec "github.com/shopspring/decimal"
)

type Share struct {
	ExpShareID int64       `json:"ExpShareID"`
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

func (sh *Share) Insert(tx *sql.Tx, expID int64) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.InsertShare(ctx, repo.InsertShareParams{
		ExpID:      expID,
		Share:      sh.Share.String(),
		UserID:     sh.User.UserID,
		Calculated: sh.Calculated.String(),
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

func (sh *Share) Update(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.UpdateShare(ctx, repo.UpdateShareParams{
		ExpShareID: sh.ExpShareID,
		Share:      sh.Share.String(),
		UserID:     sh.User.UserID,
		Calculated: sh.Calculated.String(),
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

func (sh *Share) Delete(tx *sql.Tx) error {
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
