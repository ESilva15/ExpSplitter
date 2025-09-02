package models

import (
	"context"
	repo "expenses/expenses/db/repository"
	"github.com/jackc/pgx/v5"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UserID   int32  `json:"UserID"`
	UserName string `json:"UserName"`
}

func NewUser() User {
	return User{
		UserID:   -1,
		UserName: "",
	}
}

func GetAllUsers(db repo.DBTX, tx pgx.Tx) ([]User, error) {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	userList, err := queries.GetUsers(ctx)
	if err != nil {
		return []User{}, err
	}

	return mapRepoUsers(userList), nil
}
