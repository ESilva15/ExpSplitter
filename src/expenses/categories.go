package expenses

import (
	"context"
	"database/sql"
	"fmt"

	"expenses/config"
	repo "expenses/expenses/db/repository"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

type Category struct {
	CategoryID   int
	CategoryName string
}

func NewCategory() Category {
	return Category{
		CategoryID:   -1,
		CategoryName: "",
	}
}

func GetAllCategories() ([]repo.Category, error) {
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
		return []repo.Category{}, nil
	}

	return categories, nil
}

func GetCategory(catID int64) (repo.Category, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return repo.Category{}, err
	}
	defer db.Close()

	queries := repo.New(db)
	category, err := queries.GetCategory(ctx, catID)

	return category, nil
}

func (cat *Category) Insert() error {
	cfg := config.GetInstance()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO categories(CategoryName) VALUES(?)"
	res, err := db.Exec(query, cat.CategoryName)
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

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM categories WHERE CategoryID = ?"
	res, err := db.Exec(query, cat.CategoryID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (cat *Category) Update() error {
	cfg := config.GetInstance()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE categories SET CategoryName = ? WHERE CategoryID = ?"
	res, err := db.Exec(query, cat.CategoryName, cat.CategoryID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
