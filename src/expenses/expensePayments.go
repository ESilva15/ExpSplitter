package expenses

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type ExpensePayment struct {
	ExpPaymID   int
	User        User
	PayedAmount float32
}

func (e *Expense) GetPayments() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "SELECT ExpPaymID,Payed,users.UserID,users.UserName " +
		"FROM expensesPayments " +
		"JOIN users ON users.UserID = expensesPayments.UserID " +
		"WHERE ExpID = ?"

	rows, err := db.Query(query, e.ExpID)
	if err != nil {
		return err
	}

	for rows.Next() {
		paym := &ExpensePayment{}
		err := rows.Scan(
			&paym.ExpPaymID, &paym.PayedAmount,
			&paym.User.UserID, &paym.User.UserName,
		)
		if err != nil {
			log.Fatalf("Failed to parse data from db: %v", err)
		}
		e.Payments = append(e.Payments, *paym)
	}

	return nil
}

func (p *ExpensePayment) Insert(expID int) error {
	db, err := sql.Open("sqlite3", "./data/data.db")
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
	db, err := sql.Open("sqlite3", "./data/data.db")
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
	db, err := sql.Open("sqlite3", "./data/data.db")
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
