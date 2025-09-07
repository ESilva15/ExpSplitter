package expenses

import (
	"context"
	mod "expenses/expenses/models"
)

func (a *ExpensesApp) GetAllCategories() ([]mod.Category, error) {
	ctx := context.Background()
	return a.CategoryRepo.GetAll(ctx)
}

func (a *ExpensesApp) GetCategory(id int32) (mod.Category, error) {
	ctx := context.Background()
	return a.CategoryRepo.Get(ctx, id)
}

func (a *ExpensesApp) CreateCategory(name string) error {
	ctx := context.Background()

	newCat := mod.Category{
		CategoryName: name,
	}

	return a.CategoryRepo.Insert(ctx, newCat)
}

func (a *ExpensesApp) UpdateCategory(id int32, name string) error {
	ctx := context.Background()

	cat := mod.Category{
		CategoryID:   id,
		CategoryName: name,
	}

	return a.CategoryRepo.Update(ctx, cat)
}

func (a *ExpensesApp) DeleteCategory(id int32) error {
	ctx := context.Background()
	return a.CategoryRepo.Delete(ctx, id)
}
