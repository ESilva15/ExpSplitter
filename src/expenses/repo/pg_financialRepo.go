package repo

import (
	"context"

	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgFinancialRepo struct {
	DB *pgxpool.Pool
}

func NewPgFinancialRepo(db *pgxpool.Pool) PgFinancialRepo {
	return PgFinancialRepo{DB: db}
}

func (f PgFinancialRepo) Close() {
	f.DB.Close()
}

func (p PgFinancialRepo) GetMonthlySummary(
	ctx context.Context,
	filter *mod.Filter,
) ([]mod.MonthlySummary, error) {
	queries := pgsqlc.New(p.DB)

	var pgStart, pgEnd pgtype.Timestamp
	if filter.Start != nil {
		pgStart = pgtype.Timestamp{Time: *filter.Start, Valid: true}
	}
	if filter.End != nil {
		pgEnd = pgtype.Timestamp{Time: *filter.End, Valid: true}
	}

	rows, err := queries.GetMonthlySummary(ctx, pgsqlc.GetMonthlySummaryParams{
		Startdate: pgStart,
		Enddate:   pgEnd,
		UserID:    filter.UserIDs,
	})
	if err != nil {
		return nil, err
	}

	return mapRepoMonthlySummaries(rows), nil
}

func (p PgFinancialRepo) GetMonthlySavingsAverage(
	ctx context.Context,
	filter *mod.Filter,
) (mod.MonthlySavingsAverage, error) {
	queries := pgsqlc.New(p.DB)
	res, err := queries.GetMonthlySavingsAverage(ctx)
	if err != nil {
		return mod.MonthlySavingsAverage{}, err
	}

	return mapRepoMonthlySavingsAverage(res), nil
}
