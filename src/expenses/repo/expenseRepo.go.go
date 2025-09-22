// Package repo defines the repositories for the expenses app
package repo

import (
	"context"
	"time"

	mod "github.com/ESilva15/expenses/expenses/models"
)

// ExpFilter has the fields required to query the database.
type ExpFilter struct {
	Start    *time.Time // Start is the earliest an expense can be
	End      *time.Time // End is the latest an expense can be
	CatIDs   []int32    // CatIDs is a list of categories to match against
	StoreIDs []int32    // StoreIDs is a list of stores to match against
	TypeIDs  []int32    // TypeIDs is a list of types to match against
}

// NewExpFilter returns a newly initialized ExpFilter.
func NewExpFilter() ExpFilter {
	return ExpFilter{
		Start:    nil,
		End:      nil,
		CatIDs:   nil,
		StoreIDs: nil,
		TypeIDs:  nil,
	}
}

// ExpenseRepository defines the repository for the expense data model.
type ExpenseRepository interface {
	// Direct expense methods
	Get(ctx context.Context, id int32) (mod.Expense, error)
	GetAll(ctx context.Context, filter ExpFilter, uID int32) (mod.Expenses, error)
	Update(ctx context.Context, exp mod.Expense) error
	Insert(ctx context.Context, exp mod.Expense) error
	Delete(ctx context.Context, id int32) error

	// Share methods
	GetShares(ctx context.Context, eID int32) (mod.Shares, error)
	UpdateShare(ctx context.Context, sh mod.Share) error
	InsertShare(ctx context.Context, eID int32, sh mod.Share) error
	InsertShares(ctx context.Context, eID int32, sh mod.Shares) error
	DeleteShare(ctx context.Context, id int32) error

	// Payment methods
	GetPayments(ctx context.Context, eID int32) (mod.Payments, error)
	GetExpensePaymentByUserID(ctx context.Context, eID int32,
		uID int32) (mod.Payment, error)
	UpdatePayment(ctx context.Context, pm mod.Payment) error
	InsertPayment(ctx context.Context, eID int32, pm mod.Payment) error
	InsertPayments(ctx context.Context, eID int32, pm mod.Payments) error
	DeletePayment(ctx context.Context, id int32) error

	// Debts
	SettleDebt(ctx context.Context, eID int32,
		payment mod.Payment, credit mod.Payment) error

	Close()
}
