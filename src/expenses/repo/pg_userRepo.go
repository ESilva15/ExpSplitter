package repo

import (
	"context"
	mod "expenses/expenses/models"
	"expenses/expenses/repo/pgdb/pgsqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgUserRepo struct {
	DB *pgxpool.Pool
}

func NewPgUserRepo(db *pgxpool.Pool) UserRepository {
	return PgUserRepo{
		DB: db,
	}
}

func (p PgUserRepo) Close() {
	p.DB.Close()
}

func (p PgUserRepo) GetAll(ctx context.Context) (mod.Users, error) {
	queries := pgsqlc.New(p.DB)
	userList, err := queries.GetUsers(ctx)
	if err != nil {
		return mod.Users{}, err
	}

	return mapRepoUsers(userList), nil
}
