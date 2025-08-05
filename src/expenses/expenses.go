package expenses

import (
	"context"
	"expenses/config"
	repo "expenses/expenses/db/repository"

	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetAllExpenses() ([]Expense, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return []Expense{}, err
	}
	defer db.Close()

	queries := repo.New(db)
	expenses, err := queries.GetExpenses(ctx)
	if err != nil {
		return []Expense{}, err
	}

	return mapRepoGetExpensesRows(expenses), nil
}

func GetExpense(expID int64) (Expense, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return Expense{}, err
	}
	defer db.Close()

	queries := repo.New(db)
	expense, err := queries.GetExpense(ctx, expID)
	if err != nil {
		return Expense{}, err
	}

	return mapRepoGetExpenseRow(expense), nil
}

func (exp *Expense) Insert() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queries := repo.New(tx)
	res, err := queries.InsertExpense(ctx, repo.InsertExpenseParams{
		Description:  exp.Description,
		Value:        exp.Value,
		StoreID:      exp.Store.StoreID,
		CategoryID:   exp.Category.CategoryID,
		TypeID:       exp.Type.TypeID,
		OwnerUserID:  exp.Owner.UserID,
		ExpDate:      exp.Date,
		CreationDate: exp.CreationDate,
	})
	if err != nil {
		return err
	}

	expenseID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve last inserted expense ID: %v", err)
	}

	exp.ExpID = expenseID

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

	if err = tx.Commit(); err != nil {
		// TODO
		// Add some kind of log here otherwise we could jusr return the commit res
		return err
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
		exp.Description, exp.Value, exp.Store.StoreID,
		exp.Category.CategoryID, exp.Type.TypeID, exp.Owner.UserID,
		exp.Date, exp.ExpID,
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
		"JOIN Users ON UserID = OwnerUserId " +
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
			&exp.Store.StoreID, &exp.Store.StoreName,
			&exp.Category.CategoryID, &exp.Category.CategoryName,
			&exp.Owner.UserID, &exp.Owner.UserName,
			&exp.Date, &exp.CreationDate,
		)
		if err != nil {
			log.Fatalf("Failed to parse data from db: %v", err)
		}
		expList = append(expList, *exp)
	}

	return expList, nil
}
