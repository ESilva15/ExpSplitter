// Package repo defines the repositories for the expenses app
package repo

import (
	"context"
	"time"

	mod "github.com/ESilva15/expenses/expenses/models"
)

// ExpenseRepository defines the repository for the expense data model.
type ExpenseRepository interface {
	// Direct expense methods
	Get(ctx context.Context, id int32) (mod.Expense, error)
	GetAll(ctx context.Context, uID int32) (mod.Expenses, error)
	GetExpensesRange(ctx context.Context, start time.Time,
		end time.Time, uID int32) (mod.Expenses, error)
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
