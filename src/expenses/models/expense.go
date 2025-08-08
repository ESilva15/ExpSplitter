package models

import (
	"context"
	"expenses/config"
	repo "expenses/expenses/db/repository"

	"database/sql"
	"fmt"
)

type Expense struct {
	ExpID        int64
	Description  string
	Value        float64
	Store        Store
	Type         Type
	Category     Category
	Owner        User
	Date         int64
	Payments     []ExpensePayment
	Shares       []ExpenseShare
	CreationDate int64
}

func NewExpense() Expense {
	return Expense{
		ExpID:        -1,
		Description:  "",
		Value:        0.0,
		Store:        NewStore(),
		Category:     NewCategory(),
		Owner:        NewUser(),
		Date:         0,
		Payments:     []ExpensePayment{},
		Shares:       []ExpenseShare{},
		CreationDate: 0,
	}
}

func mapRepoGetExpensesRows(er []repo.GetExpensesRow) []Expense {
	expenses := make([]Expense, len(er))
	for k, exp := range er {
		expenses[k] = mapRepoGetExpenseRowMulti(exp)
	}
	return expenses
}

func GetAllExpenses() ([]Expense, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return []Expense{}, err
	}
	defer db.Close()

	queries := repo.New(db)
	expenses, err := queries.GetExpenses(ctx, repo.GetExpensesParams{
		Startdate: nil,
		Enddate:   nil,
	})
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

	// tx, err := db.Begin()
	// if err != nil {
	// 	return err
	// }
	// defer tx.Rollback()

	queries := repo.New(db)
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

	// if err = tx.Commit(); err != nil {
	// TODO
	// Add some kind of log here otherwise we could jusr return the commit res
	// 	return err
	// }

	return nil
}

func (e *Expense) Update() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	res, err := queries.UpdateExpense(ctx, repo.UpdateExpenseParams{
		ExpID:       e.ExpID,
		Description: e.Description,
		Value:       e.Value,
		StoreID:     e.Store.StoreID,
		CategoryID:  e.Category.CategoryID,
		TypeID:      e.Type.TypeID,
		OwnerUserID: e.Owner.UserID,
		ExpDate:     e.Date,
	})
	if err != nil {
		return err
	}

	for _, share := range e.Shares {
		if share.ExpShareID == -1 {
			err := share.Insert(e.ExpID)
			if err != nil {
				return err
			}
		} else {
			share.Update()
		}
	}

	for _, paym := range e.Payments {
		if paym.ExpPaymID == -1 {
			err := paym.Insert(e.ExpID)
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
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	res, err := queries.DeleteExpense(ctx, e.ExpID)
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
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return []Expense{}, err
	}
	defer db.Close()

	queries := repo.New(db)
	expenses, err := queries.GetExpenses(ctx, repo.GetExpensesParams{
		Startdate: start,
		Enddate:   end,
	})
	if err != nil {
		return []Expense{}, err
	}

	return mapRepoGetExpensesRows(expenses), nil
}
