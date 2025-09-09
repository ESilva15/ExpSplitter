package models

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"time"
)

type Expense struct {
	ExpID        int32           `json:"ExpID"`
	Description  string          `json:"Description"`
	Value        decimal.Decimal `json:"Value"`
	Store        Store           `json:"Store"`
	Type         Type            `json:"Type"`
	Category     Category        `json:"Category"`
	Owner        User            `json:"Owner"`
	Date         time.Time       `json:"Date"`
	Payments     []Payment       `json:"Payments"`
	Shares       []Share         `json:"Shares"`
	Debts        Debts           `json:"Debts"`
	PaidOff      bool            `json:"PaidOff"`
	SharesEven   bool            `json:"SharesEven"`
	QRString     string          `json:"qr"`
	CreationDate time.Time       `json:"CreationDate"`
}
type Expenses []Expense

func NewExpense() Expense {
	return Expense{
		ExpID:        -1,
		Description:  "",
		Value:        decimal.NewFromFloat(0.0),
		Store:        NewStore(),
		Category:     NewCategory(),
		Owner:        NewUser(),
		Date:         time.Now(),
		Payments:     []Payment{},
		Shares:       []Share{},
		PaidOff:      false,
		SharesEven:   false,
		CreationDate: time.Now(),
	}
}

// ExpenseFromJSON takes []byte and returns an *Expense
func ExpenseFromJSON(data []byte) (*Expense, error) {
	var expense Expense
	err := json.Unmarshal(data, &expense)
	return &expense, err
}
