package expenses

import (
	mod "expenses/expenses/models"
)

func (a *ExpensesApp) GetAllUsers() ([]mod.User, error) {
	tx, err := a.DB.Begin()
	if err != nil {
		return []mod.User{}, err
	}
	defer tx.Rollback()

	users, err := mod.GetAllUsers(tx)
	if err != nil {
		return []mod.User{}, err
	}
	
	return users, tx.Commit()
}
