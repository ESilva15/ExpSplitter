package expenses

import (
	"reflect"
	"slices"
	"testing"
)

var (
	user1 = User{
		UserID:   1,
		UserName: "Fernando Silva",
	}
	user2 = User{
		UserID:   2,
		UserName: "Paulo Fultre",
	}
	user3 = User{
		UserID:   3,
		UserName: "Afonso Esteves",
	}
	user4 = User{
		UserID:   4,
		UserName: "Sílvio Vieira",
	}

	expense1 = Expense{
		ExpID: 0,
		Value: 150,
		Shares: []ExpenseShare{
			{ExpShareID: 0, User: user1, Share: 0.4},
			{ExpShareID: 1, User: user2, Share: 0.4},
			{ExpShareID: 2, User: user3, Share: 0.2},
		},
		Payments: []ExpensePayment{
			{ExpPaymID: 0, User: user1, PayedAmount: 50},
			{ExpPaymID: 1, User: user1, PayedAmount: 50},
			{ExpPaymID: 2, User: user2, PayedAmount: 0},
			{ExpPaymID: 3, User: user3, PayedAmount: 40},
			{ExpPaymID: 4, User: user3, PayedAmount: 20},
		},
	}

	expense2 = Expense{
		ExpID: 0,
		Value: 160,
		Shares: []ExpenseShare{
			{ExpShareID: 0, User: user1, Share: 0.3},
			{ExpShareID: 1, User: user2, Share: 0.3},
			{ExpShareID: 2, User: user3, Share: 0.3},
			{ExpShareID: 3, User: user4, Share: 0.1},
		},
		Payments: []ExpensePayment{
			{ExpPaymID: 0, User: user1, PayedAmount: 40},
			{ExpPaymID: 1, User: user1, PayedAmount: 40},
			{ExpPaymID: 2, User: user2, PayedAmount: 5},
			{ExpPaymID: 3, User: user3, PayedAmount: 30},
			{ExpPaymID: 3, User: user3, PayedAmount: 30},
			{ExpPaymID: 3, User: user3, PayedAmount: 10},
			{ExpPaymID: 4, User: user4, PayedAmount: 4},
			{ExpPaymID: 5, User: user4, PayedAmount: 1},
		},
	}
)

func TestSharesAndPaymentsMapping(t *testing.T) {
	expectedPayments := map[User]float32{
		user1: 100,
		user2: 0,
		user3: 60,
	}

	expectedShares := map[User]float32{
		user1: 0.4,
		user2: 0.4,
		user3: 0.2,
	}

	dc := NewDebtCalculator(&expense1)
	dc.mapShares()
	dc.mapPayments()

	if !reflect.DeepEqual(expectedPayments, dc.Payments) {
		t.Errorf("expected payments total and result is different:\n%+v\n%+v\n",
			expectedPayments, dc.Payments)
	}

	if !reflect.DeepEqual(expectedShares, dc.Shares) {
		t.Errorf("expected user shares and result is different:\n%+v\n%+v\n",
			expectedShares, dc.Shares)
	}
}

func TestFilterExpenseParticipants(t *testing.T) {
	expectedDebtors := []Debt{
		{user2, 43},
		{user4, 11},
	}
	slices.SortFunc(expectedDebtors, sortBySum)

	dc := NewDebtCalculator(&expense2)
	dc.mapShares()
	dc.mapPayments()

	debts := dc.getDebts()
	slices.SortFunc(debts, sortBySum)

	if !reflect.DeepEqual(expectedDebtors, debts) {
		t.Errorf("expected creditors and resulting creditors don't match:\n%+v\n%+v\n",
			expectedDebtors, debts)
	}
}
