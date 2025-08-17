package expenses

import (
	mod "expenses/expenses/models"
	"sort"

	dec "github.com/shopspring/decimal"
)

type UserTab struct {
	User mod.User
	Sum  dec.Decimal
}
type UserTabs []UserTab

func (ut UserTabs) SortBySum() {
	sort.Slice(ut, func(i, j int) bool {
		return ut[i].Sum.Cmp(ut[j].Sum) > 0
	})
}

func (ut UserTabs) Equal(t UserTabs) bool {
	if len(ut) != len(t) {
		return false
	}

	for k := range ut {
		if ut[k].User != t[k].User || !ut[k].Sum.Equal(t[k].Sum) {
			return false
		}
	}

	return true
}

func filterExpenseParticipants(e *mod.Expense) (UserTabs, UserTabs) {
	shares := mapShares(e)
	payments := mapPayments(e)

	debtors := UserTabs{}
	creditors := UserTabs{}

	for user, share := range shares {
		debt := (share.Mul(e.Value)).Sub(payments[user])
		if debt.LessThan(dec.NewFromFloat(0.0)) {
			creditors = append(creditors, UserTab{
				User: user,
				Sum:  debt.Abs(),
			})
		}

		if debt.GreaterThan(dec.NewFromFloat(0.0)) {
			debtors = append(debtors, UserTab{
				User: user,
				Sum:  debt.Abs(),
			})
		}
	}

	debtors.SortBySum()
	creditors.SortBySum()

	return debtors, creditors
}

func resolveDebts(debtors UserTabs, creditors UserTabs) []mod.Debt {
	debts := []mod.Debt{}

	// maybe create a map point to the debt of each user and count from there?
	// I should sketch this one first

	return debts
}

func CalculateDebts(e *mod.Expense) ([]mod.Debt, error) {
	// dc := NewDebtCalculator(e)
	// dc.mapShares()
	// dc.mapPayments()
	//
	// debts := dc.getDebts()

	debtors, creditors := filterExpenseParticipants(e)
	_ = resolveDebts(debtors, creditors)

	return []mod.Debt{}, nil
}
