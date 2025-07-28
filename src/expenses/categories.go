package expenses

import (
	"database/sql"
	"fmt"
	"log"

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

func GetAllCategories() ([]Category, error) {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT CategoryID,CategoryName " +
		"FROM categories"

	var categoryList []Category
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		exp := &Category{}
		err := rows.Scan(&exp.CategoryID, &exp.CategoryName)
		if err != nil {
			log.Fatalf("Failed to parse data from db: %v", err)
		}
		categoryList = append(categoryList, *exp)
	}

	return categoryList, nil
}

func GetCategory(catID int) (Category, error) {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return Category{}, err
	}
	defer db.Close()

	query := "SELECT CategoryID,CategoryName " +
		"FROM categories " +
		"WHERE CategoryID = ?"

	cat := Category{CategoryID: -1}
	err = db.QueryRow(query, catID).Scan(&cat.CategoryID, &cat.CategoryName)
	if err != nil {
		return Category{CategoryID: -1}, nil
	}

	return cat, nil
}

func (cat *Category) Insert() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
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
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}

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
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}

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
