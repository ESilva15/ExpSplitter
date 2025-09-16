package repo

import (
	"context"

	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
)

// SettleDebt will update the PG storage with a debt settlement.
//
// Returns an error if it fails.
func (p PgExpRepo) SettleDebt(
	ctx context.Context, eID int32, payment mod.Payment, credit mod.Payment,
) error {
	return withTx(p.DB, ctx, func(q *pgsqlc.Queries) error {
		err := p.insertPayment(ctx, q, eID, payment)
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
