package models

import (
	"github.com/shopspring/decimal"
	"sort"
)

type Debt struct {
	Creditor User
	Debtor   User
	Sum      decimal.Decimal
}
type Debts []Debt

func (ds Debts) SortBySum() {
	sort.SliceStable(ds, func(i, j int) bool {
		if cmp := ds[i].Sum.Cmp(ds[j].Sum); cmp != 0 {
			return cmp > 0
		}
		if ds[i].Creditor.UserID != ds[j].Creditor.UserID {
			return ds[i].Creditor.UserID < ds[j].Creditor.UserID
		}
		return ds[i].Debtor.UserID < ds[j].Debtor.UserID
	})
}

func (ds Debts) Equal(t Debts) bool {
	if len(ds) != len(t) {
		return false
	}

	for k := range ds {
		if ds[k].Creditor != t[k].Creditor || ds[k].Debtor != t[k].Debtor ||
			!ds[k].Sum.Equal(t[k].Sum) {
			return false
		}
	}

	return true
}
