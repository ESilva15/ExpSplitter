package expenses

import (
	"context"
	mod "expenses/expenses/models"
	"github.com/jackc/pgx/v5"
)

func (a *ExpensesApp) GetAllCategories() ([]mod.Category, error) {
	ctx := context.Background()

	categories, err := a.CategoryRepo.GetAll(ctx)
	if err != nil {
		return []mod.Category{}, err
	}

	return categories, nil
}

func (a *ExpensesApp) GetCategory(id int32) (mod.Category, error) {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return mod.Category{}, err
	}
	defer tx.Rollback(ctx)

	category, err := mod.GetCategory(a.DB, tx, id)
	if err != nil {
		return mod.Category{}, err
	}

	return category, tx.Commit(ctx)
}

func (a *ExpensesApp) CreateCategory(name string) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	newCat := mod.Category{
		CategoryName: name,
	}

	err = newCat.Insert(a.DB, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (a *ExpensesApp) UpdateCategory(id int32, name string) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	cat := mod.Category{
		CategoryID:   id,
		CategoryName: name,
	}
	err = cat.Update(a.DB, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (a *ExpensesApp) DeleteCategory(id int32) error {
	ctx := context.Background()
	tx, err := a.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	cat := mod.Category{
		CategoryID: id,
	}

	err = cat.Delete(a.DB, tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
