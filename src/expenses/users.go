package expenses

import (
	"context"
	mod "expenses/expenses/models"

	"golang.org/x/crypto/bcrypt"
)

// GetUser returns an user by its `id`
func (a *ExpApp) GetUser(id int32) (*mod.User, error) {
	ctx := context.Background()
	return a.UserRepo.Get(ctx, id)
}

// GetUserByName returns a user by its `name`
func (a *ExpApp) GetUserByName(name string) (*mod.User, error) {
	ctx := context.Background()
	return a.UserRepo.GetByName(ctx, name)
}

// GetAllUsers returns all the users
func (a *ExpApp) GetAllUsers(ctx context.Context) ([]mod.User, error) {
	return a.UserRepo.GetAll(ctx)
}

// ValidateCredentials validates the input credentials against the database
func (a *ExpApp) ValidateCredentials(name string, pass string) (*mod.User, error) {
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
