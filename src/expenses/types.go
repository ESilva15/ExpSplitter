package expenses

import (
	"context"
	mod "expenses/expenses/models"
	"github.com/jackc/pgx/v5"
)

func (a *ExpensesApp) GetAllTypes() ([]mod.Type, error) {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return []mod.Type{}, err
	}
	defer tx.Rollback(ctx)

	types, err := mod.GetAllTypes(a.DB, tx)
	if err != nil {
		return []mod.Type{}, err
	}

	return types, tx.Commit(ctx)
}

func (a *ExpensesApp) GetType(id int32) (mod.Type, error) {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return mod.Type{}, err
	}
	defer tx.Rollback(ctx)

	typ, err := mod.GetType(a.DB, tx, id)
	if err != nil {
		return mod.Type{}, err
	}

	return typ, tx.Commit(ctx)
}

func (a *ExpensesApp) NewType(name string) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	newTyp := mod.Type{
		TypeName: name,
	}

	err = newTyp.Insert(a.DB, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (a *ExpensesApp) DeleteType(id int32) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	typ := mod.Type{
		TypeID: id,
	}

	err = typ.Delete(a.DB, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (a *ExpensesApp) UpdateType(id int32, name string) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	newTyp := mod.Type{
		TypeID:   id,
		TypeName: name,
	}

	err = newTyp.Update(a.DB, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
