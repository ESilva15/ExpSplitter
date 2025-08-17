package expenses

import (
	mod "expenses/expenses/models"
	"testing"

	"github.com/shopspring/decimal"
)

var (
	// Use this data across the other tests if necessary
	user1 = mod.User{
		UserID:   1,
		UserName: "Fernando Silva",
	}
	user2 = mod.User{
		UserID:   2,
		UserName: "Paulo Fultre",
	}
	user3 = mod.User{
		UserID:   3,
		UserName: "Afonso Esteves",
	}
	user4 = mod.User{
		UserID:   4,
		UserName: "Sílvio Vieira",
	}

	// this expense isn't evenly shared
	// this expense is paid off
	expense1 = mod.Expense{
		ExpID: 0,
		Value: decimal.NewFromInt(150),
		Shares: []mod.ExpenseShare{
			{ExpShareID: 0, User: user1, Share: decimal.NewFromFloat(0.4)},
			{ExpShareID: 1, User: user2, Share: decimal.NewFromFloat(0.4)},
			{ExpShareID: 2, User: user3, Share: decimal.NewFromFloat(0.2)},
		},
		Payments: []mod.ExpensePayment{
			{ExpPaymID: 0, User: user1, PayedAmount: decimal.NewFromInt(50)},
			{ExpPaymID: 1, User: user1, PayedAmount: decimal.NewFromInt(50)},
			{ExpPaymID: 2, User: user2, PayedAmount: decimal.NewFromInt(0)},
			{ExpPaymID: 3, User: user3, PayedAmount: decimal.NewFromInt(30)},
			{ExpPaymID: 4, User: user3, PayedAmount: decimal.NewFromInt(20)},
		},
	}

	// Expense isn't evenly shared
	// Expense is paid off
	expense2 = mod.Expense{
		ExpID: 0,
		Value: decimal.NewFromInt(160),
		Shares: []mod.ExpenseShare{
			{ExpShareID: 0, User: user1, Share: decimal.NewFromFloat(0.3)},
			{ExpShareID: 1, User: user2, Share: decimal.NewFromFloat(0.3)},
			{ExpShareID: 2, User: user3, Share: decimal.NewFromFloat(0.3)},
			{ExpShareID: 3, User: user4, Share: decimal.NewFromFloat(0.1)},
		},
		Payments: []mod.ExpensePayment{
			{ExpPaymID: 0, User: user1, PayedAmount: decimal.NewFromInt(40)},
			{ExpPaymID: 1, User: user1, PayedAmount: decimal.NewFromInt(40)},
			{ExpPaymID: 2, User: user2, PayedAmount: decimal.NewFromInt(5)},
			{ExpPaymID: 3, User: user3, PayedAmount: decimal.NewFromInt(30)},
			{ExpPaymID: 4, User: user3, PayedAmount: decimal.NewFromInt(30)},
			{ExpPaymID: 5, User: user3, PayedAmount: decimal.NewFromInt(10)},
			{ExpPaymID: 6, User: user4, PayedAmount: decimal.NewFromInt(4)},
			{ExpPaymID: 7, User: user4, PayedAmount: decimal.NewFromInt(1)},
		},
	}

	// Expense is: evenly shared and paid off
	expense3 = mod.Expense{
		ExpID: 0,
		Value: decimal.NewFromInt(160),
		Shares: []mod.ExpenseShare{
			{ExpShareID: 0, User: user1, Share: decimal.NewFromFloat(0.3)},
			{ExpShareID: 1, User: user2, Share: decimal.NewFromFloat(0.3)},
			{ExpShareID: 2, User: user3, Share: decimal.NewFromFloat(0.3)},
			{ExpShareID: 3, User: user4, Share: decimal.NewFromFloat(0.1)},
		},
		Payments: []mod.ExpensePayment{
			{ExpPaymID: 0, User: user1, PayedAmount: decimal.NewFromInt(40)},
			{ExpPaymID: 1, User: user1, PayedAmount: decimal.NewFromInt(8)},
			{ExpPaymID: 2, User: user2, PayedAmount: decimal.NewFromInt(48)},
			{ExpPaymID: 3, User: user3, PayedAmount: decimal.NewFromInt(15)},
			{ExpPaymID: 4, User: user3, PayedAmount: decimal.NewFromInt(15)},
			{ExpPaymID: 5, User: user3, PayedAmount: decimal.NewFromInt(18)},
			{ExpPaymID: 6, User: user4, PayedAmount: decimal.NewFromInt(8)},
			{ExpPaymID: 7, User: user4, PayedAmount: decimal.NewFromInt(8)},
		},
	}
)

func TestSharesAndPaymentsMapping(t *testing.T) {
	// expectedPayments := map[mod.User]float32{
	// 	user1: 100,
	// 	user2: 0,
	// 	user3: 60,
	// }
	//
	// expectedShares := map[mod.User]float32{
	// 	user1: 0.4,
	// 	user2: 0.4,
	// 	user3: 0.2,
	// }
	//
	// dc := NewDebtCalculator(&expense1)
	// dc.mapShares()
	// dc.mapPayments()
	//
	// if !reflect.DeepEqual(expectedPayments, dc.Payments) {
	// 	t.Errorf("expected payments total and result is different:\n%+v\n%+v\n",
	// 		expectedPayments, dc.Payments)
	// }
	//
	// if !reflect.DeepEqual(expectedShares, dc.Shares) {
	// 	t.Errorf("expected user shares and result is different:\n%+v\n%+v\n",
	// 		expectedShares, dc.Shares)
	// }
}

// func TestFilterExpenseParticipants(t *testing.T) {
	// expectedDebtors := []Debt{
	// 	{user2, decimal.NewFromInt(43)},
	// 	{user4, decimal.NewFromInt(11)},
	// }
	// slices.SortFunc(expectedDebtors, sortBySum)
	//
	// dc := NewDebtCalculator(&expense2)
	// dc.mapShares()
	// dc.mapPayments()
	//
	// debts := dc.getDebts()
	// slices.SortFunc(debts, sortBySum)
	//
	// if !reflect.DeepEqual(expectedDebtors, debts) {
	// 	t.Errorf("expected creditors and resulting creditors don't match:\n%+v\n%+v\n",
	// 		expectedDebtors, debts)
	// }
// }
