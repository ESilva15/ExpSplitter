package models

import (
	"time"

	dec "github.com/shopspring/decimal"
)

type MonthlySummary struct {
	MonthYear            string      `json:"month_year"`
	Despesa              dec.Decimal `json:"despesa"`
	Salario              dec.Decimal `json:"salario"`
	CartaoRefeicao       dec.Decimal `json:"cartao_refeicao"`
	Reembolso            dec.Decimal `json:"reembolso"`
	Prenda               dec.Decimal `json:"prenda"`
	Investment           dec.Decimal `json:"investment"`
	Adjustment           dec.Decimal `json:"adjustment"`
	AdjustmentLiquid     dec.Decimal `json:"adjustment_liquid"`
	Poupanca             dec.Decimal `json:"poupanca"`
	PoupancaLiquid       dec.Decimal `json:"poupanca_liquid"`
	RunningBalance       dec.Decimal `json:"running_balance"`
	RunningBalanceLiquid dec.Decimal `json:"running_balance_liquid"`
}

type MonthlySavingsAverage struct {
	AvgPoupanca       dec.Decimal `json:"avg_poupanca"`
	AvgPoupancaLiquid dec.Decimal `json:"avg_poupanca_liquid"`
}

// ExpFilter has the fields required to query the database.
type Filter struct {
	Start    *time.Time // Start is the earliest an expense can be
	End      *time.Time // End is the latest an expense can be
	UserIDs  int32      // UserIDs is a list of IDs to match against
	CatIDs   []int32    // CatIDs is a list of categories to match against
	StoreIDs []int32    // StoreIDs is a list of stores to match against
	TypeIDs  []int32    // TypeIDs is a list of types to match against
}

// NewExpFilter returns a newly initialized ExpFilter.
func NewFilter() Filter {
	return Filter{
		Start:    nil,
		End:      nil,
		UserIDs:  -1,
		CatIDs:   nil,
		StoreIDs: nil,
		TypeIDs:  nil,
	}
}
