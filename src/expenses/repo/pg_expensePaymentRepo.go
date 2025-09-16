package repo

import (
	"context"
	"fmt"

	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
)

func (p PgExpRepo) insertPayment(
	ctx context.Context, q *pgsqlc.Queries, eID int32, pm mod.Payment,
) error {
	paid, err := decimalToNumeric(pm.PayedAmount)
	if err != nil {
		return err
	}

	res, err := q.InsertPayment(ctx, pgsqlc.InsertPaymentParams{
		ExpID:  eID,
		UserID: pm.User.UserID,
		Payed:  paid,
	})
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (p PgExpRepo) InsertPayment(
	ctx context.Context, eId int32, pm mod.Payment) error {
	return p.insertPayment(ctx, pgsqlc.New(p.DB), eId, pm)
}

func (p *PgExpRepo) insertPayments(
	ctx context.Context, q *pgsqlc.Queries, eId int32, pm []mod.Payment) error {
	for k := range pm {
		err := p.insertPayment(ctx, q, eId, pm[k])
		if err != nil {
			return err
		}
	}

	return nil
}

func (p PgExpRepo) InsertPayments(
	ctx context.Context, eId int32, pm mod.Payments) error {
	return p.insertPayments(ctx, pgsqlc.New(p.DB), eId, pm)
}

func (p *PgExpRepo) updatePayment(
	ctx context.Context, q *pgsqlc.Queries, pm mod.Payment) error {

	payed, err := decimalToNumeric(pm.PayedAmount)
	if err != nil {
		return err
	}

	res, err := q.UpdatePayment(ctx, pgsqlc.UpdatePaymentParams{
		ExpPaymID: pm.ExpPaymID,
		Payed:     payed,
		UserID:    pm.User.UserID,
	})
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (p PgExpRepo) UpdatePayment(
	ctx context.Context, pm mod.Payment) error {
	return p.updatePayment(ctx, pgsqlc.New(p.DB), pm)
}

func (p *PgExpRepo) deletePayment(
	ctx context.Context, q *pgsqlc.Queries, id int32) error {

	res, err := q.DeletePayment(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (p PgExpRepo) DeletePayment(ctx context.Context, id int32) error {
	return p.deletePayment(ctx, pgsqlc.New(p.DB), id)
}

func (p *PgExpRepo) getPayments(
	ctx context.Context, q *pgsqlc.Queries, eId int32) (mod.Payments, error) {

	payments, err := q.GetPayments(ctx, eId)
	if err != nil {
		return mod.Payments{}, err
	}

	return mapRepoGetPaymentsRows(payments), nil
}

func (p PgExpRepo) GetPayments(
	ctx context.Context, eId int32) (mod.Payments, error) {
	return p.getPayments(ctx, pgsqlc.New(p.DB), eId)
}

func (p PgExpRepo) GetExpensePaymentByUserID(ctx context.Context, eId int32, uId int32,
) (mod.Payment, error) {
	queries := pgsqlc.New(p.DB)
	payment, err := queries.GetExpensePaymentByUser(ctx, pgsqlc.GetExpensePaymentByUserParams{
		ExpID:  eId,
		UserID: uId,
	})
	if err != nil {
		return mod.Payment{}, err
	}

	return mapRepoGetPaymentRow(payment), nil
}
