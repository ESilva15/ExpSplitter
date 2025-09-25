package models

import (
	"encoding/json"
	"time"

	dec "github.com/shopspring/decimal"
)

// Expense is the model for the expense concept in our application.
type Expense struct {
	ExpID        int32       `json:"ExpID"`
	Description  string      `json:"Description"`
	Value        dec.Decimal `json:"Value"`
	Store        Store       `json:"Store"`
	Type         Type        `json:"Type"`
	Category     Category    `json:"Category"`
	Owner        User        `json:"Owner"`
	Date         time.Time   `json:"Date"`
	Payments     []Payment   `json:"Payments"`
	Shares       []Share     `json:"Shares"`
	Debts        Debts       `json:"Debts"`
	PaidOff      bool        `json:"PaidOff"`
	SharesEven   bool        `json:"SharesEven"`
	QRString     string      `json:"qr"`
	CreationDate time.Time   `json:"CreationDate"`
}

// Expenses defines a list of Expense.
type Expenses []Expense

// NewExpense returns an empty Expense.
func NewExpense() Expense {
	return Expense{
		ExpID:        -1,
		Description:  "",
		Value:        dec.NewFromFloat(0.0),
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

// ExpenseFromJSON takes []byte and returns an *Expense.
func ExpenseFromJSON(data []byte) (*Expense, error) {
	var expense Expense
	err := json.Unmarshal(data, &expense)
	return &expense, err
}

// PaidByUser returns the total paid by the user of a given ID.
func (e *Expense) PaidByUser(uID int32) dec.Decimal {
	total := dec.NewFromFloat(0.0)

	// TODO - Could we optimize this by storing the payments mapped to users in the list?
	for _, payment := range e.Payments {
		if payment.User.UserID == uID {
			total = total.Add(payment.PayedAmount)
		}
	}

	return total
}
