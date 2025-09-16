package repo

import (
	"context"

	mod "github.com/ESilva15/expenses/expenses/models"
)

// CategoryRepository defines the repository for the category data model.
type CategoryRepository interface {
	Get(ctx context.Context, id int32) (mod.Category, error)
	GetAll(ctx context.Context) (mod.Categories, error)
	Update(ctx context.Context, cat mod.Category) error
	Insert(ctx context.Context, cat mod.Category) error
	Delete(ctx context.Context, id int32) error

	Close()
}
