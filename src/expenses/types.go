package expenses

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Type struct {
	TypeID   int
	TypeName string
}

func NewType() Type {
	return Type{
		TypeID:   -1,
		TypeName: "",
	}
}

func GetAllTypes() ([]Type, error) {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT TypeID,TypeName " +
		"FROM expenseTypes"

	var typeList []Type
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		nType := &Type{}
		err := rows.Scan(&nType.TypeID, &nType.TypeName)
		if err != nil {
			log.Fatalf("Failed to parse data from db: %v", err)
		}
		typeList = append(typeList, *nType)
	}

	return typeList, nil
}

func GetType(typID int) (Type, error) {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return Type{}, err
	}
	defer db.Close()

	query := "SELECT TypeID,TypeName " +
		"FROM expenseTypes " +
		"WHERE TypeID = " + strconv.Itoa(typID)

	var nType Type
	err = db.QueryRow(query).Scan(&nType.TypeID, &nType.TypeName)
	if err != nil {
		return Type{}, nil
	}

	return nType, nil
}

func (typ *Type) Insert() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO expenseTypes(TypeName) VALUES(?)"
	res, err := db.Exec(query, typ.TypeName)
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

func (typ *Type) Delete() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}

	query := "DELETE FROM expenseTypes WHERE TypeID = ?"
	res, err := db.Exec(query, typ.TypeID)
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
