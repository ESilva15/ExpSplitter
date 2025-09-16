package repo

import (
	"context"

	mod "github.com/ESilva15/expenses/expenses/models"
)

// UserRepository defines the repository for the user data model.
type UserRepository interface {
	Get(ctx context.Context, id int32) (*mod.User, error)
	GetByName(ctx context.Context, name string) (*mod.User, error)
	GetAll(ctx context.Context) (mod.Users, error)

	Close()
}
