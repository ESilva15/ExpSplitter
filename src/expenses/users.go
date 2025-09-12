package expenses

import (
	"context"
	mod "expenses/expenses/models"
	"golang.org/x/crypto/bcrypt"
)

func (a *ExpensesApp) GetUser(id int32) (*mod.User, error) {
	ctx := context.Background()
	return a.UserRepo.Get(ctx, id)
}

func (a *ExpensesApp) GetUserByName(name string) (*mod.User, error) {
	ctx := context.Background()
	return a.UserRepo.GetByName(ctx, name)
}

func (a *ExpensesApp) GetAllUsers(ctx context.Context) ([]mod.User, error) {
	return a.UserRepo.GetAll(ctx)
}

func (a *ExpensesApp) ValidateCredentials(name string, pass string) (*mod.User, error) {
	ctx := context.Background()

	user, err := a.UserRepo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		return nil, err
	}

	return user, nil
}
