package models

import (
	"context"
	"database/sql"
	repo "expenses/expenses/db/repository"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UserID   int64
	UserName string
}

func NewUser() User {
	return User{
		UserID:   -1,
		UserName: "",
	}
}

func GetAllUsers(tx *sql.Tx) ([]User, error) {
	ctx := context.Background()

	queries := repo.New(tx)
	userList, err := queries.GetUsers(ctx)
	if err != nil {
		return []User{}, err
	}

	return mapRepoUsers(userList), nil
}
