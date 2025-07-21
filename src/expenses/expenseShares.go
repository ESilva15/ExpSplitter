package expenses

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type ExpenseShare struct {
	ExpShareID int
	User       User
	Share      float32
}

func (e *Expense) GetShares() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "SELECT ExpShareID,Share,users.UserID,users.UserName " +
		"FROM expensesShares " +
		"JOIN users ON users.UserID = expensesShares.UserID " +
		"WHERE ExpID = ?"

	rows, err := db.Query(query, e.ExpID)
	if err != nil {
		return err
	}

	for rows.Next() {
		share := &ExpenseShare{}
		err := rows.Scan(
			&share.ExpShareID, &share.Share,
			&share.User.UserID, &share.User.UserName,
		)
		if err != nil {
			log.Fatalf("Failed to parse data from db: %v", err)
		}
		e.Shares = append(e.Shares, *share)
	}

	return nil
}

func (sh *ExpenseShare) Insert(expID int) error {
	db, err := sql.Open("sqlite3", "./data/data.db")
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
	db, err := sql.Open("sqlite3", "./data/data.db")
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
