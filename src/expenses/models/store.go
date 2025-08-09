package models

import (
	"context"
	repo "expenses/expenses/db/repository"
	experr "expenses/expenses/errors"

	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	StoreID   int64
	StoreName string
}

func NewStore() Store {
	return Store{
		StoreID:   -1,
		StoreName: "",
	}
}

func GetAllStores(tx *sql.Tx) ([]Store, error) {
	ctx := context.Background()

	queries := repo.New(tx)
	storeList, err := queries.GetStores(ctx)
	if err != nil {
		return []Store{}, err
	}

	return mapRepoStores(storeList), nil
}

func GetStore(tx *sql.Tx, storeID int64) (Store, error) {
	ctx := context.Background()

	queries := repo.New(tx)
	store, err := queries.GetStore(ctx, storeID)
	if err != nil {
		return Store{}, err
	}

	return mapRepoStore(store), nil
}

func (s *Store) Insert(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.InsertStore(ctx, s.StoreName)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (s *Store) Update(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.UpdateStore(ctx, repo.UpdateStoreParams{
		StoreID:   s.StoreID,
		StoreName: s.StoreName,
	})

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}

func (s *Store) Delete(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.DeleteStore(ctx, s.StoreID)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}
