package repo

import (
	"context"
	experr "expenses/expenses/errors"
	mod "expenses/expenses/models"
	"expenses/expenses/repo/pgdb/pgsqlc"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgStoreRepo struct {
	DB *pgxpool.Pool
}

func NewPgStoreRepo(db *pgxpool.Pool) StoreRepository {
	return PgStoreRepo{
		DB: db,
	}
}

func (p PgStoreRepo) Close() {
	p.DB.Close()
}

func (p PgStoreRepo) Get(ctx context.Context, id int32) (mod.Store, error) {
	queries := pgsqlc.New(p.DB)
	store, err := queries.GetStore(ctx, id)
	if err != nil {
		return mod.Store{}, err
	}

	return mapRepoStore(store), nil
}

func (p PgStoreRepo) GetAll(ctx context.Context) (mod.Stores, error) {
	queries := pgsqlc.New(p.DB)
	storeList, err := queries.GetStores(ctx)
	if err != nil {
		return mod.Stores{}, err
	}

	return mapRepoStores(storeList), nil
}

func (p PgStoreRepo) Update(ctx context.Context, s mod.Store) error {
	queries := pgsqlc.New(p.DB)
	_, err := queries.UpdateStore(ctx, pgsqlc.UpdateStoreParams{
		StoreID:   s.StoreID,
		StoreName: s.StoreName,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p PgStoreRepo) Insert(ctx context.Context, s mod.Store) error {
	queries := pgsqlc.New(p.DB)
	res, err := queries.InsertStore(ctx, s.StoreName)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (p PgStoreRepo) Delete(ctx context.Context, id int32) error {
	queries := pgsqlc.New(p.DB)
	res, err := queries.DeleteStore(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}
