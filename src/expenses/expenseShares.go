package expenses

import (
	"database/sql"
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
