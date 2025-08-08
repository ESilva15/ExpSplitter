package models

import (
	"context"
	"expenses/config"
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

func GetAllStores() ([]Store, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return []Store{}, err
	}
	defer db.Close()

	queries := repo.New(db)
	storeList, err := queries.GetStores(ctx)
	if err != nil {
		return []Store{}, err
	}

	return mapRepoStores(storeList), nil
}

func GetStore(storeID int64) (Store, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return Store{}, err
	}
	defer db.Close()

	queries := repo.New(db)
	store, err := queries.GetStore(ctx, storeID)
	if err != nil {
		return Store{}, err
	}

	return mapRepoStore(store), nil
}

func (s *Store) Insert() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	res, err := queries.InsertStore(ctx, s.StoreName)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (s *Store) Update() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
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

func (s *Store) Delete() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	res, err := queries.DeleteStore(ctx, s.StoreID)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}
