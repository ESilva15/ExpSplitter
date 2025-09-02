package models

import (
	"context"
	"encoding/json"
	repo "expenses/expenses/db/repository"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type Expense struct {
	ExpID        int32            `json:"ExpID"`
	Description  string           `json:"Description"`
	Value        decimal.Decimal  `json:"Value"`
	Store        Store            `json:"Store"`
	Type         Type             `json:"Type"`
	Category     Category         `json:"Category"`
	Owner        User             `json:"Owner"`
	Date         time.Time        `json:"Date"`
	Payments     []ExpensePayment `json:"Payments"`
	Shares       []Share          `json:"Shares"`
	Debts        Debts            `json:"Debts"`
	PaidOff      bool             `json:"PaidOff"`
	SharesEven   bool             `json:"SharesEven"`
	CreationDate time.Time        `json:"CreationDate"`
}

func NewExpense() Expense {
	return Expense{
		ExpID:        -1,
		Description:  "",
		Value:        decimal.NewFromFloat(0.0),
		Store:        NewStore(),
		Category:     NewCategory(),
		Owner:        NewUser(),
		Date:         time.Now(),
		Payments:     []ExpensePayment{},
		Shares:       []Share{},
		PaidOff:      false,
		SharesEven:   false,
		CreationDate: time.Now(),
	}
}

// ExpenseFromJSON takes []byte and returns an *Expense
func ExpenseFromJSON(data []byte) (*Expense, error) {
	var expense Expense

	err := json.Unmarshal(data, &expense)

	return &expense, err
}

func GetAllExpenses(db repo.DBTX, tx pgx.Tx) ([]Expense, error) {
	ctx := context.Background()

	var start pgtype.Timestamp
	start.Valid = false
	var end pgtype.Timestamp
	end.Valid = false

	queries := repo.New(db).WithTx(tx)
	expenses, err := queries.GetExpenses(ctx, repo.GetExpensesParams{
		Startdate: start,
		Enddate:   end,
	})
	if err != nil {
		return []Expense{}, err
	}

	return mapRepoGetExpensesRows(expenses), nil
}

func GetExpense(db repo.DBTX, tx pgx.Tx, expID int32) (Expense, error) {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	expense, err := queries.GetExpense(ctx, expID)
	if err != nil {
		return Expense{}, err
	}

	return mapRepoGetExpenseRow(expense), nil
}

func (exp *Expense) Insert(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	value, err := decimalToNumeric(exp.Value)
	if err != nil {
		return err
	}

	paidOff, err := boolToPgBool(exp.PaidOff)
	if err != nil {
		return err
	}

	sharesEven, err := boolToPgBool(exp.SharesEven)
	if err != nil {
		return err
	}

	expDate, err := timeToTimestamp(exp.Date)
	if err != nil {
		return err
	}

	creationDate, err := timeToTimestamp(exp.CreationDate)
	if err != nil {
		return err
	}

	queries := repo.New(db).WithTx(tx)
	lastInsertedID, err := queries.InsertExpense(ctx, repo.InsertExpenseParams{
		Description:  exp.Description,
		Value:        value,
		StoreID:      exp.Store.StoreID,
		CategoryID:   exp.Category.CategoryID,
		TypeID:       exp.Type.TypeID,
		OwnerUserID:  exp.Owner.UserID,
		ExpDate:      expDate,
		PaidOff:      paidOff,
		SharesEven:   sharesEven,
		CreationDate: creationDate,
	})
	if err != nil {
		return err
	}

	exp.ExpID = lastInsertedID

	err = exp.InsertShares(db, tx)
	if err != nil {
		return err
	}

	err = exp.InsertPayments(db, tx)
	if err != nil {
		return err
	}

	return nil
}

func (exp *Expense) InsertShares(db repo.DBTX, tx pgx.Tx) error {
	for _, share := range exp.Shares {
		err := share.Insert(db, tx, exp.ExpID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (exp *Expense) InsertPayments(db repo.DBTX, tx pgx.Tx) error {
	for _, paym := range exp.Payments {
		err := paym.Insert(db, tx, exp.ExpID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *Expense) Update(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	value, err := decimalToNumeric(e.Value)
	if err != nil {
		return err
	}

	paidOff, err := boolToPgBool(e.PaidOff)
	if err != nil {
		return err
	}

	sharesEven, err := boolToPgBool(e.SharesEven)
	if err != nil {
		return err
	}

	expDate, err := timeToTimestamp(e.Date)
	if err != nil {
		return err
	}

	queries := repo.New(db).WithTx(tx)
	res, err := queries.UpdateExpense(ctx, repo.UpdateExpenseParams{
		ExpID:       e.ExpID,
		Description: e.Description,
		Value:       value,
		StoreID:     e.Store.StoreID,
		CategoryID:  e.Category.CategoryID,
		TypeID:      e.Type.TypeID,
		OwnerUserID: e.Owner.UserID,
		PaidOff:     paidOff,
		SharesEven:  sharesEven,
		ExpDate:     expDate,
	})
	if err != nil {
		return err
	}

	for _, share := range e.Shares {
		if share.ExpShareID == -1 {
			err := share.Insert(db, tx, e.ExpID)
			if err != nil {
				return err
			}
		} else {
			share.Update(db, tx)
		}
	}

	for _, paym := range e.Payments {
		if paym.ExpPaymID == -1 {
			err := paym.Insert(db, tx, e.ExpID)
			if err != nil {
				return err
			}
		} else {
			paym.Update(db, tx)
		}
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return nil
}

func (e *Expense) Delete(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	res, err := queries.DeleteExpense(ctx, e.ExpID)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

// SHITTY QUERIES THAT I NEED TO PUT IN SOMEWHERE MORE OGRANHJASLD
func GetExpensesRange(db repo.DBTX, tx pgx.Tx, start time.Time, end time.Time) ([]Expense, error) {
	ctx := context.Background()

	startPg, err := timeToTimestamp(start)
	if err != nil {
		return []Expense{}, err
	}

	endPg, err := timeToTimestamp(end)
	if err != nil {
		return []Expense{}, err
	}

	queries := repo.New(db).WithTx(tx)
	expenses, err := queries.GetExpenses(ctx, repo.GetExpensesParams{
		Startdate: startPg,
		Enddate:   endPg,
	})
	if err != nil {
		return []Expense{}, err
	}

	return mapRepoGetExpensesRows(expenses), nil
}
