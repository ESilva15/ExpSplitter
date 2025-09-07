package repo

import (
	"context"
	mod "expenses/expenses/models"
)

type CategoryRepository interface {
	Get(ctx context.Context, id int32) (mod.Category, error)
	GetAll(ctx context.Context) (mod.Categories, error)
	Update(ctx context.Context, cat mod.Category) error
	Insert(ctx context.Context, cat mod.Category) error
	Delete(ctx context.Context, id int32) error

	Close() 
}
