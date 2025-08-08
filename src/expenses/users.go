package expenses

import (
	mod "expenses/expenses/models"
)

func GetAllUsers() ([]mod.User, error) {
	return mod.GetAllUsers()
}
