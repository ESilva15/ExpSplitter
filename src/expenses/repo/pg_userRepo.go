package repo

import (
	"context"

	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PgUserRepo is a PG repository for users.
type PgUserRepo struct {
	DB *pgxpool.Pool
}

// NewPgUserRepo returns a new UserRepository of Postgres.
func NewPgUserRepo(db *pgxpool.Pool) PgUserRepo {
	return PgUserRepo{
		DB: db,
	}
}

// Close closes a PgUserRepo instance.
func (p PgUserRepo) Close() {
	p.DB.Close()
}

// Get returns an user by its ID.
func (p PgUserRepo) Get(ctx context.Context, id int32) (*mod.User, error) {
	queries := pgsqlc.New(p.DB)
	user, err := queries.GetUser(ctx, id)
	if err != nil {
		return &mod.User{}, err
	}

	u := mapRepoUser(user)

	return &u, nil
}

// GetByName returns a user by its name.
func (p PgUserRepo) GetByName(ctx context.Context, name string) (*mod.User, error) {
	queries := pgsqlc.New(p.DB)
	user, err := queries.GetUserByName(ctx, name)
	if err != nil {
		return &mod.User{}, err
	}

	u := mapRepoUser(user)

	return &u, nil
}

// GetAll returns all users.
func (p PgUserRepo) GetAll(ctx context.Context) (mod.Users, error) {
	queries := pgsqlc.New(p.DB)
	userList, err := queries.GetUsers(ctx)
	if err != nil {
		return mod.Users{}, err
	}

	return mapRepoUsers(userList), nil
}
