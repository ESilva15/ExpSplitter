package models

import (
	"context"
	repo "expenses/expenses/db/repository"
	experr "expenses/expenses/errors"

	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Type struct {
	TypeID   int64  `json:"TypeID"`
	TypeName string `json:"TypeName"`
}

func NewType() Type {
	return Type{
		TypeID:   -1,
		TypeName: "",
	}
}

func GetAllTypes(tx *sql.Tx) ([]Type, error) {
	ctx := context.Background()

	queries := repo.New(tx)
	typeList, err := queries.GetTypes(ctx)
	if err != nil {
		return []Type{}, err
	}

	return mapRepoTypes(typeList), nil
}

func GetType(tx *sql.Tx, typID int64) (Type, error) {
	ctx := context.Background()

	queries := repo.New(tx)
	typ, err := queries.GetType(ctx, typID)
	if err != nil {
		return Type{}, err
	}

	return mapRepoType(typ), nil
}

func (typ *Type) Insert(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.InsertType(ctx, typ.TypeName)
	if err != nil {
		return err
	}

	// TODO
	// Move all theses things outside of this, I'll handle it wherever I'm doing
	// logics
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (typ *Type) Delete(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.DeleteType(ctx, typ.TypeID)
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

func (typ *Type) Update(tx *sql.Tx) error {
	ctx := context.Background()

	queries := repo.New(tx)
	res, err := queries.UpdateType(ctx, repo.UpdateTypeParams{
		TypeID:   typ.TypeID,
		TypeName: typ.TypeName,
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
