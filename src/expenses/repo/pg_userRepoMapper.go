package repo

import (
	mod "expenses/expenses/models"
	"expenses/expenses/repo/pgdb/pgsqlc"
)

func mapRepoUser(ru pgsqlc.User) mod.User {
	return mod.User{
		UserID:   ru.UserID,
		UserName: ru.UserName,
		Password: ru.UserPass,
	}
}

func mapRepoUsers(ru []pgsqlc.User) mod.Users {
	users := make(mod.Users, len(ru))
	for k, u := range ru {
		users[k] = mapRepoUser(u)
	}
	return users
}
