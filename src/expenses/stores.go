package expenses

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	StoreID   int
	StoreName string
}

func NewStore() Store {
	return Store{
		StoreID:   -1,
		StoreName: "",
	}
}

func GetAllStores() ([]Store, error) {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT StoreID,StoreName " +
		"FROM stores"

	var storeList []Store
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		store := &Store{}
		err := rows.Scan(&store.StoreID, &store.StoreName)
		if err != nil {
			log.Fatalf("Failed to parse data from db: %v", err)
		}
		storeList = append(storeList, *store)
	}

	return storeList, nil
}

func GetStore(storeID int) (Store, error) {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return Store{}, err
	}
	defer db.Close()

	query := "SELECT StoreID,StoreName " +
		"FROM stores " +
		"WHERE StoreID = ?"

	store := Store{StoreID: -1}
	err = db.QueryRow(query, storeID).Scan(&store.StoreID, &store.StoreName)
	if err != nil {
		return Store{StoreID: -1}, nil
	}

	return store, nil
}

func (s *Store) Insert() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO stores(StoreName) VALUES(?)"
	res, err := db.Exec(query, s.StoreName)
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

func (s *Store) Update() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}

	query := "UPDATE stores SET StoreName = ? WHERE StoreID = ?"
	res, err := db.Exec(query, s.StoreName, s.StoreID)
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

func (s *Store) Delete() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}

	query := "DELETE FROM stores WHERE StoreID = ?"
	res, err := db.Exec(query, s.StoreID)
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
