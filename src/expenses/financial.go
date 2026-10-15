package expenses

import (
	"context"

	mod "github.com/ESilva15/expenses/expenses/models"
)

func (a *ExpApp) GetMonthlySummary(
	ctx context.Context,
	filter *mod.Filter,
) ([]mod.MonthlySummary, error) {
	return a.FinancialRepo.GetMonthlySummary(ctx, filter)
}

func (a *ExpApp) GetMonthlySavingsAverage(
	ctx context.Context,
	filter *mod.Filter,
) (mod.MonthlySavingsAverage, error) {
	return a.FinancialRepo.GetMonthlySavingsAverage(ctx, filter)
}
