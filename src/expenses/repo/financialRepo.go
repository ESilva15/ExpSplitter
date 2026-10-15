package repo

import (
	"context"

	mod "github.com/ESilva15/expenses/expenses/models"
)

// CategoryRepository defines the repository for the category data model.
type FinancialRepository interface {
	GetMonthlySummary(context.Context, *mod.Filter) ([]mod.MonthlySummary, error)
	GetMonthlySavingsAverage(context.Context, *mod.Filter) (mod.MonthlySavingsAverage, error)
	Close()
}
