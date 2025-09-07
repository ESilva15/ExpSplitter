package models

import (
	"encoding/json"
	"github.com/shopspring/decimal"
)

type Payment struct {
	ExpPaymID   int32
	User        User
	PayedAmount decimal.Decimal
}
type Payments []Payment

// PaymentFromJSON takes []byte and returns an *ExpensePayment
func PaymentFromJSON(data []byte) (*Payment, error) {
	var payment Payment
	err := json.Unmarshal(data, &payment)
	return &payment, err
}
