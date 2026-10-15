package repo

import (
	mod "github.com/ESilva15/expenses/expenses/models"
	pgsqlc "github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
)

func mapRepoMonthlySummary(r pgsqlc.GetMonthlySummaryRow) mod.MonthlySummary {
	return mod.MonthlySummary{
		MonthYear:            r.MonthYear,
		Despesa:              pgNumericToDecimal(r.Despesa),
		Salario:              pgNumericToDecimal(r.Salario),
		CartaoRefeicao:       pgNumericToDecimal(r.CartaoRefeicao),
		Reembolso:            pgNumericToDecimal(r.Reembolso),
		Prenda:               pgNumericToDecimal(r.Prenda),
		Investment:           pgNumericToDecimal(r.Investment),
		Adjustment:           pgNumericToDecimal(r.Adjustment),
		AdjustmentLiquid:     pgNumericToDecimal(r.AdjustmentLiquid),
		Poupanca:             pgNumericToDecimal(r.Poupanca),
		PoupancaLiquid:       pgNumericToDecimal(r.PoupancaLiquid),
		RunningBalance:       pgNumericToDecimal(r.RunningBalance),
		RunningBalanceLiquid: pgNumericToDecimal(r.RunningBalanceLiquid),
	}
}

func mapRepoMonthlySummaries(rows []pgsqlc.GetMonthlySummaryRow) []mod.MonthlySummary {
	summaries := make([]mod.MonthlySummary, len(rows))
	for i, row := range rows {
		summaries[i] = mapRepoMonthlySummary(row)
	}
	return summaries
}

func mapRepoMonthlySavingsAverage(r pgsqlc.GetMonthlySavingsAverageRow) mod.MonthlySavingsAverage {
	return mod.MonthlySavingsAverage{
		AvgPoupanca:       pgNumericToDecimal(r.AvgPoupanca),
		AvgPoupancaLiquid: pgNumericToDecimal(r.AvgPoupancaLiquid),
	}
}
