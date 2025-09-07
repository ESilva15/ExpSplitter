package repo

import (
	"context"
	mod "expenses/expenses/models"
	"expenses/expenses/repo/pgdb/pgsqlc"
	"fmt"
)

func (p *PgExpRepo) insertShare(
	ctx context.Context,
	q *pgsqlc.Queries,
	eId int32,
	sh mod.Share) error {

	share, err := decimalToNumeric(sh.Share)
	if err != nil {
		return err
	}
	calculated, err := decimalToNumeric(sh.Calculated)
	if err != nil {
		return err
	}

	res, err := q.InsertShare(ctx, pgsqlc.InsertShareParams{
		ExpID:      eId,
		Share:      share,
		UserID:     sh.User.UserID,
		Calculated: calculated,
	})
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}

	return nil
}

func (p PgExpRepo) InsertShare(
	ctx context.Context,
	eId int32,
	sh mod.Share) error {

	return p.insertShare(ctx, pgsqlc.New(p.DB), eId, sh)
}

func (p PgExpRepo) InsertShares(
	ctx context.Context,
	eId int32,
	sh mod.Shares) error {

	return p.insertShares(ctx, pgsqlc.New(p.DB), eId, sh)
}

func (p *PgExpRepo) insertShares(
	ctx context.Context,
	q *pgsqlc.Queries,
	eId int32,
	sh []mod.Share) error {

	for k := range sh {
		err := p.insertShare(ctx, q, eId, sh[k])
		if err != nil {
			return err
		}
	}

	return nil
}

func (p PgExpRepo) updateShare(
	ctx context.Context, q *pgsqlc.Queries, sh mod.Share) error {

	share, err := decimalToNumeric(sh.Share)
	if err != nil {
		return err
	}
	calculated, err := decimalToNumeric(sh.Calculated)
	if err != nil {
		return err
	}

	res, err := q.UpdateShare(ctx, pgsqlc.UpdateShareParams{
		ExpShareID: sh.ExpShareID,
		Share:      share,
		UserID:     sh.User.UserID,
		Calculated: calculated,
	})
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return nil
}

func (p PgExpRepo) UpdateShare(ctx context.Context, sh mod.Share) error {
	return p.updateShare(ctx, pgsqlc.New(p.DB), sh)
}

func (p *PgExpRepo) deleteShare(
	ctx context.Context, q *pgsqlc.Queries, sh mod.Share) error {

	res, err := q.DeleteShare(ctx, sh.ExpShareID)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted")
	}

	return nil
}

func (p PgExpRepo) DeleteShare(ctx context.Context, sh mod.Share) error {
	return p.deleteShare(ctx, pgsqlc.New(p.DB), sh)
}

func (p *PgExpRepo) getShares(
	ctx context.Context, q *pgsqlc.Queries, eId int32) (mod.Shares, error) {

	shares, err := q.GetShares(ctx, eId)
	if err != nil {
		return mod.Shares{}, err
	}

	return mapRepoGetSharesRows(shares), nil
}

func (p PgExpRepo) GetShares(ctx context.Context, eId int32) (mod.Shares, error) {
	return p.getShares(ctx, pgsqlc.New(p.DB), eId)
}
