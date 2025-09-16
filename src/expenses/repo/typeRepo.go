package repo

import (
	"context"

	mod "github.com/ESilva15/expenses/expenses/models"
)

// TypeRepository defines the repository for the type data model.
type TypeRepository interface {
	Get(ctx context.Context, id int32) (mod.Type, error)
	GetAll(ctx context.Context) (mod.Types, error)
	Update(ctx context.Context, typ mod.Type) error
	Insert(ctx context.Context, typ mod.Type) error
	Delete(ctx context.Context, id int32) error

	Close()
}
