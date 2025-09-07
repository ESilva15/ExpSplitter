package expenses

import (
	"context"
	mod "expenses/expenses/models"
)

func (a *ExpensesApp) GetAllUsers() ([]mod.User, error) {
	ctx := context.Background()
	return a.UserRepo.GetAll(ctx)
}
