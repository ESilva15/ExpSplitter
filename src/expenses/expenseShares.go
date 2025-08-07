package expenses

import (
	"context"
	"expenses/config"
	"expenses/expenses/db/repository"
	repo "expenses/expenses/db/repository"

	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func (e *Expense) GetShares() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repository.New(db)
	shares, err := queries.GetShares(ctx, e.ExpID)
	if err != nil {
		return err
	}

	e.Shares = mapRepoGetSharesRows(shares)
	return nil
}

func (sh *ExpenseShare) Insert(expID int64) error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	res, err := queries.InsertShare(ctx, repo.InsertShareParams{
		ExpID:  expID,
		Share:  sh.Share,
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

func (sh *ExpenseShare) Update() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	res, err := queries.UpdateShare(ctx, repo.UpdateShareParams{
		ExpShareID: sh.ExpShareID,
		Share:      sh.Share,
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
