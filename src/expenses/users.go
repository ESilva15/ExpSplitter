package expenses

import (
	"context"
	mod "expenses/expenses/models"
)

func (a *ExpensesApp) GetUser(id int32) (*mod.User, error) {
	ctx := context.Background()
	return a.UserRepo.Get(ctx, id)
}

func (a *ExpensesApp) GetUserByName(name string) (*mod.User, error) {
	ctx := context.Background()
	return a.UserRepo.GetByName(ctx, name)
}

func (a *ExpensesApp) GetAllUsers() ([]mod.User, error) {
	ctx := context.Background()
	return a.UserRepo.GetAll(ctx)
}
