package expenses

import (
	"database/sql"
	"log"

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
