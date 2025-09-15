package expenses

import (
	mod "expenses/expenses/models"
	"fmt"
	"testing"

	dec "github.com/shopspring/decimal"
)

func TestFilterExpenseParticipants(t *testing.T) {
	testCases := []struct {
		name              string
		exp               *mod.Expense
		expectedDebtors   userTabs
		expectedCreditors userTabs
	}{
		{
			name: "Expense something something",
			exp:  &expense2,
			expectedCreditors: userTabs{
				{User: user1, Sum: dec.NewFromFloat(32.0)},
				{User: user3, Sum: dec.NewFromFloat(22.0)},
			},
			expectedDebtors: userTabs{
				{User: user2, Sum: dec.NewFromFloat(43.0)},
				{User: user4, Sum: dec.NewFromFloat(11.0)},
			},
		},
		{
			name: "Expense cents",
			exp:  &expense6,
			expectedCreditors: userTabs{
				{User: user1, Sum: dec.NewFromFloat(0.32)},
			},
			expectedDebtors: userTabs{
				{User: user2, Sum: dec.NewFromFloat(0.16)},
				{User: user3, Sum: dec.NewFromFloat(0.16)},
			},
		},
	}

	for _, tt := range testCases {
		fmt.Printf("Running test: %s\n", tt.name)
		debtors, creditors := filterExpenseParticipants(tt.exp)

		if !tt.expectedDebtors.Equal(debtors) {
			t.Errorf("Expected debtors:\n%+v\nGotten debtors:\n%+v\n",
				tt.expectedDebtors, debtors)
		}

		if !tt.expectedCreditors.Equal(creditors) {
			t.Errorf("Expected creditors:\n%+v\nGotten creditors:\n%+v\n",
				tt.expectedCreditors, creditors)
		}
	}
}

func TestResolveDebt_ExactDebtsForCredit(t *testing.T) {
	// Test1:
	// Debt is solvable
	debtors := userTabs{
		{User: user1, Sum: dec.NewFromFloat(10.0)},
		{User: user2, Sum: dec.NewFromFloat(10.0)},
		{User: user3, Sum: dec.NewFromFloat(10.0)},
	}
	creditor := userTab{User: user4, Sum: dec.NewFromFloat(30.0)}
	expectedDebts := mod.Debts{
		{Creditor: user4, Debtor: user1, Sum: dec.NewFromFloat(10.0)},
		{Creditor: user4, Debtor: user2, Sum: dec.NewFromFloat(10.0)},
		{Creditor: user4, Debtor: user3, Sum: dec.NewFromFloat(10.0)},
	}
	expectedDebtors := userTabs{
		{User: user1, Sum: dec.NewFromFloat(0.0)},
		{User: user2, Sum: dec.NewFromFloat(0.0)},
		{User: user3, Sum: dec.NewFromFloat(0.0)},
	}

	result := resolveDebt(creditor, debtors)

	if !result.Equal(expectedDebts) {
		t.Errorf("Expected debts:\n%+v\nResult debts:\n%+v\n",
			expectedDebts, result)
	}

	if !debtors.Equal(expectedDebtors) {
		t.Errorf("Expected debtors:\n%+v\nResult debtors:\n%+v\n",
			expectedDebtors, debtors)
	}
}

func TestResolveDebt_LessDebtsThanCredit(t *testing.T) {
	// Test2:
	// Debt is not solvable:
	debtors := userTabs{
		{User: user1, Sum: dec.NewFromFloat(10.0)},
		{User: user2, Sum: dec.NewFromFloat(10.0)},
	}
	creditor := userTab{User: user4, Sum: dec.NewFromFloat(30.0)}
	expectedDebts := mod.Debts{
		{Creditor: user4, Debtor: user1, Sum: dec.NewFromFloat(10.0)},
		{Creditor: user4, Debtor: user2, Sum: dec.NewFromFloat(10.0)},
	}

	result := resolveDebt(creditor, debtors)

	if !result.Equal(expectedDebts) {
		t.Errorf("[Test2] Expected debts:\n%+v\nResult debts:\n%+v\n",
			expectedDebts, result)
	}
}

func TestResolveDebt_EnoughDebtsForCredit(t *testing.T) {
	debtors := userTabs{
		{User: user1, Sum: dec.NewFromFloat(10.0)},
		{User: user2, Sum: dec.NewFromFloat(10.0)},
		{User: user3, Sum: dec.NewFromFloat(10.0)},
	}
	creditor := userTab{User: user4, Sum: dec.NewFromFloat(10.0)}
	expectedDebts := mod.Debts{
		{Creditor: user4, Debtor: user1, Sum: dec.NewFromFloat(10.0)},
	}

	result := resolveDebt(creditor, debtors)

	if !result.Equal(expectedDebts) {
		t.Errorf("[Test3] Expected debts:\n%+v\nResult debts:\n%+v\n",
			expectedDebts, result)
	}
}

func TestResolveDebt_RenameMe(t *testing.T) {
	debtors := userTabs{
		{User: user1, Sum: dec.NewFromFloat(10.0)},
	}
	expectedDebtors := userTabs{
		{User: user1, Sum: dec.NewFromFloat(3.0)},
	}
	creditor := userTab{User: user4, Sum: dec.NewFromFloat(7.0)}
	expectedDebts := mod.Debts{
		{Creditor: user4, Debtor: user1, Sum: dec.NewFromFloat(7.0)},
	}

	result := resolveDebt(creditor, debtors)

	if !result.Equal(expectedDebts) {
		t.Errorf("[Test3] Expected debts:\n%+v\nResult debts:\n%+v\n",
			expectedDebts, result)
	}

	if !debtors.Equal(expectedDebtors) {
		t.Errorf("Expected debtors:\n%+v\nResult debtors:\n%+v\n",
			expectedDebtors, debtors)
	}
}

// TODO
// Create more test cases for this thing, make sure it really works
func TestResolveDebts(t *testing.T) {
	expectedDebts := mod.Debts{
		{Creditor: user3, Debtor: user4, Sum: dec.NewFromFloat(11.0)},
		{Creditor: user3, Debtor: user2, Sum: dec.NewFromFloat(11.0)},
		{Creditor: user1, Debtor: user2, Sum: dec.NewFromFloat(32.0)},
	}
	expectedDebts.SortBySum()

	debtors, creditors := filterExpenseParticipants(&expense2)
	debts := resolveDebts(debtors, creditors)

	debts.SortBySum()

	if !expectedDebts.Equal(debts) {
		t.Errorf("Expected debts:\n%+v\nResult debts:\n%+v\n", expectedDebts, debts)
	}
}
