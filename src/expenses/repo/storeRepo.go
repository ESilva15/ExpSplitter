package repo

import (
	"context"

	mod "github.com/ESilva15/expenses/expenses/models"
)

// StoreRepository defines the repository for the store data model.
type StoreRepository interface {
	Get(ctx context.Context, id int32) (mod.Store, error)
	GetByNIF(ctx context.Context, nif string) (mod.Store, error)
	GetAll(ctx context.Context) (mod.Stores, error)
	Update(ctx context.Context, cat mod.Store) error
	Insert(ctx context.Context, cat mod.Store) error
	Delete(ctx context.Context, id int32) error

	Close()
}
