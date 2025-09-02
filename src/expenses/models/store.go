package models

import (
	"context"
	repo "expenses/expenses/db/repository"
	experr "expenses/expenses/errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	StoreID   int32  `json:"StoreID"`
	StoreName string `json:"StoreName"`
}

func NewStore() Store {
	return Store{
		StoreID:   -1,
		StoreName: "",
	}
}

func GetAllStores(db repo.DBTX, tx pgx.Tx) ([]Store, error) {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	storeList, err := queries.GetStores(ctx)
	if err != nil {
		return []Store{}, err
	}

	return mapRepoStores(storeList), nil
}

func GetStore(db repo.DBTX, tx pgx.Tx, storeID int32) (Store, error) {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	store, err := queries.GetStore(ctx, storeID)
	if err != nil {
		return Store{}, err
	}

	return mapRepoStore(store), nil
}

func (s *Store) Insert(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
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

func (s *Store) Update(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	_, err := queries.UpdateStore(ctx, repo.UpdateStoreParams{
		StoreID:   s.StoreID,
		StoreName: s.StoreName,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Delete(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	res, err := queries.DeleteStore(ctx, s.StoreID)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}
