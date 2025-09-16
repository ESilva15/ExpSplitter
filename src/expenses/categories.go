package expenses

import (
	"context"

	mod "github.com/ESilva15/expenses/expenses/models"
)

// GetAllCategories returns all categories.
func (a *ExpApp) GetAllCategories(ctx context.Context) ([]mod.Category, error) {
	return a.CategoryRepo.GetAll(ctx)
}

// GetCategory returns a category by its ID.
func (a *ExpApp) GetCategory(id int32) (mod.Category, error) {
	ctx := context.Background()

	return a.CategoryRepo.Get(ctx, id)
}

// CreateCategory creates a new category.
func (a *ExpApp) CreateCategory(name string) error {
	ctx := context.Background()

	newCat := mod.Category{
		CategoryName: name,
	}

	return a.CategoryRepo.Insert(ctx, newCat)
}

// UpdateCategory updates a category by its ID.
func (a *ExpApp) UpdateCategory(id int32, name string) error {
	ctx := context.Background()

	cat := mod.Category{
		CategoryID:   id,
		CategoryName: name,
	}

	return a.CategoryRepo.Update(ctx, cat)
}

// DeleteCategory deletes a category by its ID.
func (a *ExpApp) DeleteCategory(id int32) error {
	ctx := context.Background()

	return a.CategoryRepo.Delete(ctx, id)
}
