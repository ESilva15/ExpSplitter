package models

import (
	"context"
	repo "expenses/expenses/db/repository"
	experr "expenses/expenses/errors"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Category struct {
	CategoryID   int32  `json:"CategoryID"`
	CategoryName string `json:"CategoryName"`
}

func NewCategory() Category {
	return Category{
		CategoryID:   -1,
		CategoryName: "",
	}
}

func GetAllCategories(db repo.DBTX, tx pgx.Tx) ([]Category, error) {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	categories, err := queries.GetCategories(ctx)
	if err != nil {
		return []Category{}, nil
	}

	return MapRepoCategories(categories), nil
}

func GetCategory(db repo.DBTX, tx pgx.Tx, catID int32) (Category, error) {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	category, err := queries.GetCategory(ctx, catID)

	return MapRepoCategory(category), err
}

func (cat *Category) Insert(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	res, err := queries.InsertCategory(ctx, cat.CategoryName)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (cat *Category) Delete(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	res, err := queries.DeleteCategory(ctx, cat.CategoryID)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}

func (cat *Category) Update(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	res, err := queries.UpdateCategory(ctx, repo.UpdateCategoryParams{
		CategoryName: cat.CategoryName,
		CategoryID:   cat.CategoryID,
	})
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}
