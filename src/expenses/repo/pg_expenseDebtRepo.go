package repo

import (
	"context"
	mod "expenses/expenses/models"
	"expenses/expenses/repo/pgdb/pgsqlc"
)

func (p PgExpRepo) SettleDebt(
	ctx context.Context, eId int32, payment mod.Payment, credit mod.Payment) error {

	return withTx(p.DB, ctx, func(q *pgsqlc.Queries) error {
		err := p.insertPayment(ctx, q, eId, payment)
		if err != nil {
			return err
		}

		err = p.updatePayment(ctx, q, credit)
		if err != nil {
			return err
		}

		return nil // yeah we could just return the last op, but I like this way
	})
}
