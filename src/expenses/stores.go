package expenses

import (
	"context"
	"github.com/jackc/pgx/v5"
	mod "expenses/expenses/models"
)

func (a *ExpensesApp) GetAllStores() ([]mod.Store, error) {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return []mod.Store{}, err
	}
	defer tx.Rollback(ctx)

	stores, err := mod.GetAllStores(a.DB, tx)
	if err != nil {
		return []mod.Store{}, err
	}

	return stores, tx.Commit(ctx)
}

func (a *ExpensesApp) GetStore(id int32) (mod.Store, error) {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return mod.Store{}, err
	}
	defer tx.Rollback(ctx)

	store, err := mod.GetStore(a.DB, tx, id)
	if err != nil {
		return mod.Store{}, err
	}

	return store, tx.Commit(ctx)
}

func (a *ExpensesApp) NewStore(name string) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	newStore := mod.Store{
		StoreName: name,
	}

	err = newStore.Insert(a.DB, tx)

	return tx.Commit(ctx)
}

func (a *ExpensesApp) UpdateStore(id int32, name string) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	store := mod.Store{
		StoreID:   id,
		StoreName: name,
	}

	err = store.Update(a.DB, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (a *ExpensesApp) DeleteStore(id int32) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	store := mod.Store{
		StoreID: id,
	}

	err = store.Delete(a.DB, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
