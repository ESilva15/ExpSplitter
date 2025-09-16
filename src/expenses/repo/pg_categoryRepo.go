package repo

import (
	"context"
	"fmt"

	experr "github.com/ESilva15/expenses/expenses/errors"
	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PgCatRepo is the PG repository for categories.
type PgCatRepo struct {
	DB *pgxpool.Pool
}

// NewPgCatRepo returns a new PG repository for categories.
func NewPgCatRepo(db *pgxpool.Pool) PgCatRepo {
	return PgCatRepo{
		DB: db,
	}
}

// Close closes a PgCatRepo.
func (p PgCatRepo) Close() {
	p.DB.Close()
}

// Get fetches a category by its id
//
// Returns an empty mod.Category and error if it fails.
func (p PgCatRepo) Get(ctx context.Context, id int32) (mod.Category, error) {
	queries := pgsqlc.New(p.DB)
	category, err := queries.GetCategory(ctx, id)

	return mapRepoCategory(category), err
}

// GetAll fetches all categories.
//
// Returns an empty mod.Categories if it fails and an error.
func (p PgCatRepo) GetAll(ctx context.Context) (mod.Categories, error) {
	queries := pgsqlc.New(p.DB)
	categories, err := queries.GetCategories(ctx)
	if err != nil {
		return mod.Categories{}, err
	}

	return mapRepoCategories(categories), nil
}

// Update updates the parameter cat
//
// Returns an error if it fails.
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

// Insert inserts a new category
//
// Returns an error if it fails.
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

// Delete deletes a category by ID
//
// Returns an error if it fails.
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
