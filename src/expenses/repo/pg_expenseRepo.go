package repo

import (
	"context"
	"fmt"

	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgExpRepo struct {
	DB *pgxpool.Pool
}

func NewPgExpRepo(db *pgxpool.Pool) ExpenseRepository {
	return PgExpRepo{
		DB: db,
	}
}

func (p PgExpRepo) Close() {
	p.DB.Close()
}

func withTx(
	ctx context.Context,
	pgPool *pgxpool.Pool,
	fn func(q *pgsqlc.Queries) error,
) error {
	tx, err := pgPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	q := pgsqlc.New(tx)

	if err := fn(q); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

// GetAll fetches all expenses.
func (p PgExpRepo) GetAll(ctx context.Context, filter ExpFilter, uID int32,
) (mod.Expenses, error) {
	startPg, err := timeToTimestamp(filter.Start)
	if err != nil {
		return mod.Expenses{}, err
	}

	endPg, err := timeToTimestamp(filter.End)
	if err != nil {
		return mod.Expenses{}, err
	}

	queries := pgsqlc.New(p.DB)
	expenses, err := queries.GetExpenses(ctx, pgsqlc.GetExpensesParams{
		Startdate: startPg,
		Enddate:   endPg,
		UserID:    uID,
		Catids:    filter.CatIDs,
		Storeids:  filter.StoreIDs,
		Typeids:   filter.TypeIDs,
	})
	if err != nil {
		return mod.Expenses{}, err
	}

	return mapRepoGetExpensesRows(expenses), nil
}

func (p *PgExpRepo) get(ctx context.Context, q *pgsqlc.Queries, id int32) (mod.Expense, error) {
	expense, err := q.GetExpense(ctx, id)
	if err != nil {
		return mod.Expense{}, err
	}

	return mapRepoExpenseRow(ExpenseRowSingle(expense)), nil
}

// Get fetches a single expense by its ID.
func (p PgExpRepo) Get(ctx context.Context, id int32) (mod.Expense, error) {
	return p.get(ctx, pgsqlc.New(p.DB), id)
}

// Update updates a given expense.
func (p PgExpRepo) Update(
	ctx context.Context,
	exp mod.Expense,
) error {
	return withTx(ctx, p.DB, func(q *pgsqlc.Queries) error {
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

		expDate, err := timeToTimestamp(&exp.Date)
		if err != nil {
			return err
		}

		res, err := q.UpdateExpense(ctx, pgsqlc.UpdateExpenseParams{
			ExpID:       exp.ExpID,
			Description: exp.Description,
			Value:       value,
			StoreID:     exp.Store.StoreID,
			CategoryID:  exp.Category.CategoryID,
			TypeID:      exp.Type.TypeID,
			OwnerUserID: exp.Owner.UserID,
			PaidOff:     paidOff,
			SharesEven:  sharesEven,
			Qr:          exp.QRString,
			ExpDate:     expDate,
		})
		if err != nil {
			return err
		}

		for _, share := range exp.Shares {
			if share.ExpShareID == -1 {
				err := p.insertShare(ctx, q, exp.ExpID, share)
				if err != nil {
					return err
				}
			} else {
				err := p.updateShare(ctx, q, share)
				if err != nil {
					return err
				}
			}
		}

		for _, paym := range exp.Payments {
			if paym.ExpPaymID == -1 {
				err := p.insertPayment(ctx, q, exp.ExpID, paym)
				if err != nil {
					return err
				}
			} else {
				err := p.updatePayment(ctx, q, paym)
				if err != nil {
					return err
				}
			}
		}

		rowsAffected := res.RowsAffected()
		if rowsAffected == 0 {
			return fmt.Errorf("no rows were updated")
		}

		return nil
	})
}

func (p PgExpRepo) Insert(
	ctx context.Context,
	exp mod.Expense,
) error {
	return withTx(ctx, p.DB, func(q *pgsqlc.Queries) error {
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

		expDate, err := timeToTimestamp(&exp.Date)
		if err != nil {
			return err
		}

		creationDate, err := timeToTimestamp(&exp.CreationDate)
		if err != nil {
			return err
		}

		lastInsertedID, err := q.InsertExpense(ctx, pgsqlc.InsertExpenseParams{
			Description:  exp.Description,
			Value:        value,
			StoreID:      exp.Store.StoreID,
			CategoryID:   exp.Category.CategoryID,
			TypeID:       exp.Type.TypeID,
			OwnerUserID:  exp.Owner.UserID,
			ExpDate:      expDate,
			PaidOff:      paidOff,
			SharesEven:   sharesEven,
			Qr:           exp.QRString,
			CreationDate: creationDate,
		})
		if err != nil {
			return err
		}

		exp.ExpID = lastInsertedID

		err = p.insertShares(ctx, q, exp.ExpID, exp.Shares)
		if err != nil {
			return err
		}

		err = p.insertPayments(ctx, q, exp.ExpID, exp.Payments)
		if err != nil {
			return err
		}

		return nil
	})
}

// Delete deletes a given expense by its ID.
func (p PgExpRepo) Delete(ctx context.Context, id int32) error {
	queries := pgsqlc.New(p.DB)
	res, err := queries.DeleteExpense(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were delete")
	}

	return nil
}
