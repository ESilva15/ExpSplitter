package repo

import (
	"context"
	experr "expenses/expenses/errors"
	mod "expenses/expenses/models"
	"expenses/expenses/repo/pgdb/pgsqlc"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgCatRepo struct {
	DB *pgxpool.Pool
}

func NewPgCatRepo(db *pgxpool.Pool) CategoryRepository {
	return PgCatRepo{
		DB: db,
	}
}

func (p PgCatRepo) Get(ctx context.Context, id int32) (mod.Category, error) {
	queries := pgsqlc.New(p.DB)
	category, err := queries.GetCategory(ctx, id)
	return MapRepoCategory(category), err
}

func (p PgCatRepo) GetAll(ctx context.Context) (mod.Categories, error) {
	queries := pgsqlc.New(p.DB)
	categories, err := queries.GetCategories(ctx)
	if err != nil {
		return mod.Categories{}, nil
	}
	return MapRepoCategories(categories), nil
}

func (p PgCatRepo) Update(ctx context.Context, cat mod.Category) error {
	queries := pgsqlc.New(p.DB)
	res, err := queries.UpdateCategory(ctx, pgsqlc.UpdateCategoryParams{
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

func (p PgCatRepo) Insert(ctx context.Context, cat mod.Category) error {
	queries := pgsqlc.New(p.DB)
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

func (p PgCatRepo) Delete(ctx context.Context, id int32) error {
	queries := pgsqlc.New(p.DB)
	res, err := queries.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}
