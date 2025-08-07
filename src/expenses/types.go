package expenses

import (
	"context"
	"expenses/config"
	repo "expenses/expenses/db/repository"
	experr "expenses/expenses/errors"

	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewType() Type {
	return Type{
		TypeID:   -1,
		TypeName: "",
	}
}

func GetAllTypes() ([]Type, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return []Type{}, err
	}
	defer db.Close()

	queries := repo.New(db)
	typeList, err := queries.GetTypes(ctx)
	if err != nil {
		return []Type{}, err
	}

	return mapRepoTypes(typeList), nil
}

func GetType(typID int64) (Type, error) {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
	if err != nil {
		return Type{}, err
	}
	defer db.Close()

	queries := repo.New(db)
	typ, err := queries.GetType(ctx, typID)
	if err != nil {
		return Type{}, err
	}

	return mapRepoType(typ), nil
}

func (typ *Type) Insert() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
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

func (typ *Type) Delete() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
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

func (typ *Type) Update() error {
	cfg := config.GetInstance()
	ctx := context.Background()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		return err
	}
	defer db.Close()

	queries := repo.New(db)
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
