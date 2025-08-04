package expenses

import (
	"context"
	"expenses/config"
	"expenses/expenses/db/repository"

	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Expense struct {
	ExpID        int
	Description  string
	Value        float32
	ExpStore     Store
	ExpType      Type
	ExpCategory  Category
	OwnerUser    User
	ExpDate      int64
	Payments     []ExpensePayment
	Shares       []ExpenseShare
	CreationDate int64
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
		Payments:     []ExpensePayment{},
		Shares:       []ExpenseShare{},
		CreationDate: 0,
	}
}

func GetAllExpenses() ([]repository.GetExpensesRow, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := repository.New(db)
	expenses, err := queries.GetExpenses(ctx)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func GetExpense(expID int64) (repository.GetExpenseRow, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return repository.GetExpenseRow{}, err
	}
	defer db.Close()

	queries := repository.New(db)
	expense, err := queries.GetExpense(ctx, expID)
	if err != nil {
		return repository.GetExpenseRow{}, err
	}

	return expense, nil
}

func (exp *Expense) Insert() error {
	cfg := config.GetInstance()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO expenses(" +
		"Description,Value,StoreID,CategoryID,TypeID,OwnerUserID,ExpDate,CreationDate" +
		") " +
		"VALUES(?, ?, ? , ?, ?, ?, ?, ?)"

	// TODO
	// Add a transaction here so that it fails when necessary
	res, err := db.Exec(query, exp.Description, exp.Value, exp.ExpStore.StoreID,
		exp.ExpCategory.CategoryID, exp.ExpType.TypeID, 1, exp.ExpDate, exp.CreationDate,
	)
	if err != nil {
		return err
	}

	expenseID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve last inserted expense ID: %v", err)
	}
	exp.ExpID = int(expenseID)
	for _, share := range exp.Shares {
		err := share.Insert(exp.ExpID)
		if err != nil {
			return err
		}
	}
	for _, paym := range exp.Payments {
		err := paym.Insert(exp.ExpID)
		if err != nil {
			return err
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (exp *Expense) Update() error {
	cfg := config.GetInstance()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE expenses " +
		"SET " +
		"Description = ?," +
		"Value = ?," +
		"StoreID = ?," +
		"CategoryID = ?," +
		"TypeID = ?," +
		"OwnerUserID = ?," +
		"ExpDate = ?" +
		"WHERE ExpID = ?"

	res, err := db.Exec(query,
		exp.Description, exp.Value, exp.ExpStore.StoreID, 
		exp.ExpCategory.CategoryID, exp.ExpType.TypeID, exp.OwnerUser.UserID, 
		exp.ExpDate, exp.ExpID,
	)
	if err != nil {
		return err
	}

	for _, share := range exp.Shares {
		if share.ExpShareID == -1 {
			err := share.Insert(exp.ExpID)
			if err != nil {
				return err
			}
		} else {
			share.Update()
		}
	}

	for _, paym := range exp.Payments {
		if paym.ExpPaymID == -1 {
			err := paym.Insert(exp.ExpID)
			if err != nil {
				return err
			}
		} else {
			paym.Update()
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return nil
}

func (e *Expense) Delete() error {
	cfg := config.GetInstance()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM expenses " +
		"WHERE ExpID = ?"

	res, err := db.Exec(query, e.ExpID)
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

// SHITTY QUERIES THAT I NEED TO PUT IN SOMEWHERE MORE OGRANHJASLD
func GetExpensesRange(start int64, end int64) ([]Expense, error) {
	cfg := config.GetInstance()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
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
		"JOIN Stores ON stores.StoreID = expenses.StoreID " +
		"JOIN Categories ON categories.CategoryID = expenses.CategoryID " +
		"JOIN Users ON UserID = OwnerUserId "+
		"WHERE ExpDate >= ? and ExpDate <= ?"

	var expList []Expense
	rows, err := db.Query(query, start, end)
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
