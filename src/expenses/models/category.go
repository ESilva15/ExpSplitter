package models

import (
	"context"
	"database/sql"
	config "expenses/config"
	repo "expenses/expenses/db/repository"
	experr "expenses/expenses/errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Category struct {
	CategoryID   int64  `json:"CategoryID"`
	CategoryName string `json:"CategoryName"`
}

func NewCategory() Category {
	return Category{
		CategoryID:   -1,
		CategoryName: "",
	}
}

func GetAllCategories() ([]Category, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	queries := repo.New(db)
	categories, err := queries.GetCategories(ctx)
	if err != nil {
		return []Category{}, nil
	}

	return MapRepoCategories(categories), nil
}

func GetCategory(catID int64) (Category, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return Category{}, err
	}
	defer db.Close()

	queries := repo.New(db)
	category, err := queries.GetCategory(ctx, catID)

	return MapRepoCategory(category), err
}

func (cat *Category) Insert(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.InsertCategory(ctx, cat.CategoryName)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (cat *Category) Delete() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	res, err := queries.DeleteCategory(ctx, cat.CategoryID)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}

func (cat *Category) Update() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
	res, err := queries.UpdateCategory(ctx, repo.UpdateCategoryParams{
		CategoryName: cat.CategoryName,
		CategoryID:   cat.CategoryID,
	})
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}
