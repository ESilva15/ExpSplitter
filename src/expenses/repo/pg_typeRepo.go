package repo

import (
	"context"
	experr "expenses/expenses/errors"
	mod "expenses/expenses/models"
	"expenses/expenses/repo/pgdb/pgsqlc"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgTypeRepo struct {
	DB *pgxpool.Pool
}

func NewPgTypeRepo(db *pgxpool.Pool) TypeRepository {
	return PgTypeRepo{
		DB: db,
	}
}

func (p PgTypeRepo) Close() {
	p.DB.Close()
}

func (p PgTypeRepo) Get(ctx context.Context, id int32) (mod.Type, error) {
	queries := pgsqlc.New(p.DB)
	typ, err := queries.GetType(ctx, id)
	if err != nil {
		return mod.Type{}, err
	}

	return mapRepoType(typ), nil
}

func (p PgTypeRepo) GetAll(ctx context.Context) (mod.Types, error) {
	queries := pgsqlc.New(p.DB)
	typeList, err := queries.GetTypes(ctx)
	if err != nil {
		return mod.Types{}, err
	}

	return mapRepoTypes(typeList), nil
}

func (p PgTypeRepo) Update(ctx context.Context, typ mod.Type) error {
	queries := pgsqlc.New(p.DB)
	res, err := queries.UpdateType(ctx, pgsqlc.UpdateTypeParams{
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

func (p PgTypeRepo) Insert(ctx context.Context, typ mod.Type) error {
	queries := pgsqlc.New(p.DB)
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

func (p PgTypeRepo) Delete(ctx context.Context, id int32) error {
	queries := pgsqlc.New(p.DB)
	res, err := queries.DeleteType(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return experr.ErrNotFound
	}

	return nil
}
