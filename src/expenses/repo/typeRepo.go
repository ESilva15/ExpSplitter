package repo

import (
	"context"
	mod "expenses/expenses/models"
)

type TypeRepository interface {
	Get(ctx context.Context, id int32) (mod.Type, error)
	GetAll(ctx context.Context) (mod.Types, error)
	Update(ctx context.Context, typ mod.Type) error
	Insert(ctx context.Context, typ mod.Type) error
	Delete(ctx context.Context, id int32) error

	Close()
}
