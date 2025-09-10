package repo

import (
	"context"
	mod "expenses/expenses/models"
)

type UserRepository interface {
	Get(ctx context.Context, id int32) (*mod.User, error)
	GetByName(ctx context.Context, name string) (*mod.User, error)
	GetAll(ctx context.Context) (mod.Users, error)

	Close()
}
