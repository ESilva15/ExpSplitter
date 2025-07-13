package expenses

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Expense struct {
	ExpID        int
	Description  string
	Value        float32
	StoreID      int
	CategoryID   int
	OwnerUserID  int
	ExpDate      int
	CreationDate int
}

func GetAllExpenses() ([]Expense, error) {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT ExpID,Description,Value,StoreID,CategoryID," +
		"OwnerUserID,ExpDate,CreationDate " +
		"FROM expenses"

	var expList []Expense
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		exp := &Expense{}
		err := rows.Scan(&exp.ExpID, &exp.Description, &exp.Value, &exp.StoreID,
			&exp.CategoryID, &exp.OwnerUserID, &exp.ExpDate, &exp.CreationDate,
		)
		if err != nil {
			log.Fatalf("Failed to parse data from db: %v", err)
		}
		expList = append(expList, *exp)
	}

	return expList, nil
}

func GetExpense(expID int) (Expense, error) {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return Expense{}, err
	}
	defer db.Close()

	query := "SELECT ExpID,Description,Value,StoreID,CategoryID," +
		"OwnerUserID,ExpDate,CreationDate " +
		"FROM expenses " +
		"WHERE ExpID = " + strconv.Itoa(expID)

	var exp Expense
	err = db.QueryRow(query).Scan(
		&exp.ExpID, &exp.Description, &exp.Value, &exp.StoreID,
		&exp.CategoryID, &exp.OwnerUserID, &exp.ExpDate, &exp.CreationDate,
	)
	if err != nil {
		return Expense{}, nil
	}

	return exp, nil
}
