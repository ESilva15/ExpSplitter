package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	dec "github.com/shopspring/decimal"
)

// CustomDate is here so we can scan it.
type CustomDate struct {
	time.Time
}

// CustomDateLayout defines the format we are using for expenses dates.
const CustomDateLayout = "02-Jan-2006"

// UnmarshalJSON unmarshalles our custom date to time.Time.
func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return fmt.Errorf("time cannot be empty")
	}

	t, err := time.Parse(CustomDateLayout, s)
	if err != nil {
		return err
	}

	cd.Time = t
	return nil
}

// Expense is the model for the expense concept in our application.
type Expense struct {
	ExpID        int32       `json:"ExpID"`
	Description  string      `json:"Description"`
	Value        dec.Decimal `json:"Value"`
	Store        Store       `json:"Store"`
	Type         Type        `json:"Type"`
	Category     Category    `json:"Category"`
	Owner        User        `json:"Owner"`
	Date         CustomDate  `json:"Date"`
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
		Date:         CustomDate{time.Now()},
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
