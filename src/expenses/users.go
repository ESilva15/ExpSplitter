package expenses

import (
	"context"
	mod "expenses/expenses/models"

	"github.com/jackc/pgx/v5"
)

func (a *ExpensesApp) GetAllUsers() ([]mod.User, error) {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return []mod.User{}, err
	}
	defer tx.Rollback(ctx)

	users, err := mod.GetAllUsers(a.DB, tx)
	if err != nil {
		return []mod.User{}, err
	}

	return users, tx.Commit(ctx)
}
