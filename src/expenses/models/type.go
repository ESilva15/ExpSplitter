package models

import (
	"context"
	repo "expenses/expenses/db/repository"
	experr "expenses/expenses/errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/mattn/go-sqlite3"
)

type Type struct {
	TypeID   int32  `json:"TypeID"`
	TypeName string `json:"TypeName"`
}

func NewType() Type {
	return Type{
		TypeID:   -1,
		TypeName: "",
	}
}

func GetAllTypes(db repo.DBTX, tx pgx.Tx) ([]Type, error) {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	typeList, err := queries.GetTypes(ctx)
	if err != nil {
		return []Type{}, err
	}

	return mapRepoTypes(typeList), nil
}

func GetType(db repo.DBTX, tx pgx.Tx, typID int32) (Type, error) {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	typ, err := queries.GetType(ctx, typID)
	if err != nil {
		return Type{}, err
	}

	return mapRepoType(typ), nil
}

func (typ *Type) Insert(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	res, err := queries.InsertType(ctx, typ.TypeName)
	if err != nil {
		return err
	}

	// TODO
	// Move all theses things outside of this, I'll handle it wherever I'm doing
	// logics
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (typ *Type) Delete(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	res, err := queries.DeleteType(ctx, typ.TypeID)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}

func (typ *Type) Update(db repo.DBTX, tx pgx.Tx) error {
	ctx := context.Background()

	queries := repo.New(db).WithTx(tx)
	res, err := queries.UpdateType(ctx, repo.UpdateTypeParams{
		TypeID:   typ.TypeID,
		TypeName: typ.TypeName,
	})
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}
