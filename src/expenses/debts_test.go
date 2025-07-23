package expenses

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The debt calculation is pretty much the core functionality of this thing
// Formula to calculate how much a given user owes:
// (share * total) - payedAmount

// TestCase1
// A simple expense of 150 with:
// Shares:
// User1 0.5
// User2 0.5

// Payments  | Debts
// User1 135 | (0.5 * 150) - 135 = -60 (is owed 60)
// User2  15 | (0.5 * 150) - 15  =  60 (owes 60)

// Summary:
// User2 owes User1 60

func TestCalculateDebtsTestCase1(t *testing.T) {
	expense := Expense{
		ExpID:       1,
		Description: "Test Exp",
		Value:       150.0,
		ExpType: Type{
			TypeID:   1,
			TypeName: "Despesa",
		},
		ExpCategory: Category{
			CategoryID:   0,
			CategoryName: "Groceries",
		},
		OwnerUser: User{
			UserID:   1,
			UserName: "ESilva",
		},
		ExpDate: 1753307982,
		Payments: []ExpensePayment{
			{
				ExpPaymID: 0,
				User: User{
					UserID:   1,
					UserName: "ESilva",
				},
				PayedAmount: 135,
			},
			{
				ExpPaymID: 1,
				User: User{
					UserID:   2,
					UserName: "Kika",
				},
				PayedAmount: 15,
			},
		},
		Shares: []ExpenseShare{
			{
				ExpShareID: 0,
				User: User{
					UserID:   1,
					UserName: "ESilva",
				},
				Share: 0.5,
			},
			{
				ExpShareID: 1,
				User: User{
					UserID:   2,
					UserName: "Kika",
				},
				Share: 0.5,
			},
		},
	}

	expectedDebts := []Debt{
		{
			Creditor: User{
				UserID:   1,
				UserName: "ESilva",
			},
			Debtor: User{
				UserID:   2,
				UserName: "Kika",
			},
			Sum: 60,
		},
	}
	slices.SortFunc(expectedDebts, sortBySum)

	resultDebts, err := expense.calculateDebts()
	if err != nil {
		t.Fatalf("Debt calculation failed: %+v", err)
	}
	slices.SortFunc(resultDebts, sortBySum)

	assert.Equal(t, expectedDebts, resultDebts)
}
