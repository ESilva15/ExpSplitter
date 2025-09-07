package repo

import (
	"context"
	mod "expenses/expenses/models"
)

type UserRepository interface {
	GetAll(ctx context.Context) (mod.Users, error)
}
