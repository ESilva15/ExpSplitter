package repo

import (
	"context"
	mod "expenses/expenses/models"
	"expenses/expenses/repo/pgdb/pgsqlc"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

func (r *PgExpRepo) withTx(ctx context.Context, fn func(q *pgsqlc.Queries) error) error {
	tx, err := r.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	q := pgsqlc.New(tx) // bind sqlc to this tx

	if err := fn(q); err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (p *PgExpRepo) getAll(ctx context.Context, q *pgsqlc.Queries) (mod.Expenses, error) {
	// We want to select all here, no timestamps
	var start pgtype.Timestamp
	start.Valid = false
	var end pgtype.Timestamp
	end.Valid = false

	expenses, err := q.GetExpenses(ctx, pgsqlc.GetExpensesParams{
		Startdate: start,
		Enddate:   end,
	})
	if err != nil {
		return []mod.Expense{}, err
	}

	return mapRepoGetExpensesRows(expenses), err
}

func (p PgExpRepo) GetAll(ctx context.Context) (mod.Expenses, error) {
	return p.getAll(ctx, pgsqlc.New(p.DB))
}

func (p *PgExpRepo) get(ctx context.Context, q *pgsqlc.Queries, id int32) (mod.Expense, error) {
	expense, err := q.GetExpense(ctx, id)
	if err != nil {
		return mod.Expense{}, err
	}
	return mapRepoGetExpenseRow(expense), nil
}

func (p PgExpRepo) Get(ctx context.Context, id int32) (mod.Expense, error) {
	return p.get(ctx, pgsqlc.New(p.DB), id)
}

func (p PgExpRepo) Update(
	ctx context.Context,
	exp mod.Expense) error {

	return p.withTx(ctx, func(q *pgsqlc.Queries) error {
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
					return nil
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
	exp mod.Expense) error {

	return p.withTx(ctx, func(q *pgsqlc.Queries) error {
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

func (p PgExpRepo) GetExpensesRange(
	ctx context.Context, start time.Time, end time.Time) (mod.Expenses, error) {

	startPg, err := timeToTimestamp(start)
	if err != nil {
		return mod.Expenses{}, err
	}

	endPg, err := timeToTimestamp(end)
	if err != nil {
		return mod.Expenses{}, err
	}

	queries := pgsqlc.New(p.DB)
	expenses, err := queries.GetExpenses(ctx, pgsqlc.GetExpensesParams{
		Startdate: startPg,
		Enddate:   endPg,
	})
	if err != nil {
		return mod.Expenses{}, err
	}

	return mapRepoGetExpensesRows(expenses), nil
}
