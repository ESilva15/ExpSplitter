package expenses

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

type Category struct {
	CategoryID   int
	CategoryName string
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
		"WHERE CategoryID = " + strconv.Itoa(catID)

	var cat Category
	err = db.QueryRow(query).Scan(&cat.CategoryID, &cat.CategoryName)
	if err != nil {
		return Category{}, nil
	}

	return cat, nil
}
