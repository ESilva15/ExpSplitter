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
	ExpStore     Store
	ExpCategory  Category
	OwnerUser    User
	ExpDate      int
	CreationDate int
}

func NewExpense() Expense {
	return Expense{
		ExpID:        -1,
		Description:  "",
		Value:        0.0,
		ExpStore:     NewStore(),
		ExpCategory:  NewCategory(),
		OwnerUser:    NewUser(),
		ExpDate:      0,
		CreationDate: 0,
	}
}

func GetAllExpenses() ([]Expense, error) {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT ExpID,Description,Value," +
		"Stores.StoreID,Stores.StoreName," +
		"Categories.CategoryID,Categories.CategoryName," +
		"Users.UserID,Users.UserName," +
		"ExpDate,CreationDate " +
		"FROM expenses " +
		"JOIN Stores, Categories " +
		"JOIN Users ON UserID = OwnerUserId"

	var expList []Expense
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		exp := &Expense{}
		err := rows.Scan(
			&exp.ExpID, &exp.Description, &exp.Value,
			&exp.ExpStore.StoreID, &exp.ExpStore.StoreName,
			&exp.ExpCategory.CategoryID, &exp.ExpCategory.CategoryName,
			&exp.OwnerUser.UserID, &exp.OwnerUser.UserName,
			&exp.ExpDate, &exp.CreationDate,
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

	query := "SELECT ExpID,Description,Value," +
		"Stores.StoreID,Stores.StoreName," +
		"Categories.CategoryID,Categories.CategoryName," +
		"Users.UserID,Users.UserName," +
		"ExpDate,CreationDate " +
		"FROM expenses " +
		"JOIN Stores, Categories " +
		"JOIN Users ON UserID = OwnerUserId " +
		"WHERE ExpID = " + strconv.Itoa(expID)

	log.Println(query)

	exp := Expense{ExpID: -1}
	err = db.QueryRow(query).Scan(
		&exp.ExpID, &exp.Description, &exp.Value,
		&exp.ExpStore.StoreID, &exp.ExpStore.StoreName,
		&exp.ExpCategory.CategoryID, &exp.ExpCategory.CategoryName,
		&exp.OwnerUser.UserID, &exp.OwnerUser.UserName,
		&exp.ExpDate, &exp.CreationDate,
	)
	if err == sql.ErrNoRows {
		return Expense{ExpID: -1}, nil
	}

	if err != nil {
		return Expense{}, nil
	}

	return exp, nil
}
