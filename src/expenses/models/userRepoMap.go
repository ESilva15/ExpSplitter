package models

import repo "expenses/expenses/db/repository"

func mapRepoUser(ru repo.User) User {
	return User{
		UserID:   ru.UserID,
		UserName: ru.UserName,
	}
}

func mapRepoUsers(ru []repo.User) []User {
	users := make([]User, len(ru))
	for k, u := range ru {
		users[k] = mapRepoUser(u)
	}
	return users
}
