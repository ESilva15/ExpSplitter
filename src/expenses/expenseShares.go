package expenses

import (
	"context"
	"expenses/config"
	"expenses/expenses/db/repository"

	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type ExpenseShare struct {
	ExpShareID int
	User       User
	Share      float32
}

func GetShares(expID int64) ([]repository.GetSharesRow, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return []repository.GetSharesRow{}, err
	}
	defer db.Close()

	queries := repository.New(db)
	shares, err := queries.GetShares(ctx, expID)
	if err != nil {
		return []repository.GetSharesRow{}, err
	}

	return shares, nil
}

func (sh *ExpenseShare) Insert(expID int) error {
	cfg := config.GetInstance()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO expensesShares(" +
		"ExpID,UserID,Share" +
		") " +
		"VALUES(?, ?, ?)"

	res, err := db.Exec(query, expID, sh.User.UserID, sh.Share)
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

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE expensesShares " +
		"SET " +
		"UserID = ?," +
		"Share = ? " +
		"WHERE ExpShareID = ?"

	res, err := db.Exec(query, sh.User.UserID, sh.Share, sh.ExpShareID)
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
