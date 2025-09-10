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

func (p PgUserRepo) Get(ctx context.Context, id int32) (*mod.User, error) {
	queries := pgsqlc.New(p.DB)
	user, err := queries.GetUser(ctx, id)
	if err != nil {
		return &mod.User{}, err
	}

	u := mapRepoUser(user)

	return &u, nil
}

func (p PgUserRepo)	GetByName(ctx context.Context, name string) (*mod.User, error) {
	queries := pgsqlc.New(p.DB)
	user, err := queries.GetUserByName(ctx, name)
	if err != nil {
		return &mod.User{}, err
	}

	u := mapRepoUser(user)

	return &u, nil
}

func (p PgUserRepo) GetAll(ctx context.Context) (mod.Users, error) {
	queries := pgsqlc.New(p.DB)
	userList, err := queries.GetUsers(ctx)
	if err != nil {
		return mod.Users{}, err
	}

	return mapRepoUsers(userList), nil
}
