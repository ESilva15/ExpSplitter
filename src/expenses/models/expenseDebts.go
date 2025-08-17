package models

import (
	"github.com/shopspring/decimal"
)

type Debt struct {
	Creditor User
	Debtor   User
	Sum      decimal.Decimal
}

func (d *Debt) Insert() {
}

// func (d *Debt) Update() {
// }
