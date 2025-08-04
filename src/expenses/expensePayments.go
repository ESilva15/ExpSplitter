package expenses

import (
	"context"
	"expenses/config"
	"expenses/expenses/db/repository"

	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type ExpensePayment struct {
	ExpPaymID   int
	User        User
	PayedAmount float32
}

func GetPayments(expID int64) ([]repository.GetPaymentsRow, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return []repository.GetPaymentsRow{}, err
	}
	defer db.Close()

	queries := repository.New(db)
	payments, err := queries.GetPayments(ctx, expID)
	if err != nil {
		return []repository.GetPaymentsRow{}, err
	}

	return payments, nil
}

func (p *ExpensePayment) Insert(expID int) error {
	cfg := config.GetInstance()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO expensesPayments(" +
		"ExpID,UserID,Payed" +
		") " +
		"VALUES(?, ?, ?)"

	res, err := db.Exec(query, expID, p.User.UserID, p.PayedAmount)
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

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE expensesPayments " +
		"SET " +
		"UserID = ?," +
		"Payed = ? " +
		"WHERE ExpPaymID = ?"

	res, err := db.Exec(query, p.User.UserID, p.PayedAmount, p.ExpPaymID)
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

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM expensesPayments " +
		"WHERE ExpPaymID = ?"

	res, err := db.Exec(query, p.ExpPaymID)
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
