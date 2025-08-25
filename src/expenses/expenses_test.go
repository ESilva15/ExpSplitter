package expenses

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestExpenseTotalPayed(t *testing.T) {
	expectedTotal := decimal.NewFromInt(160)

	result := ExpenseTotalPayed(&expense1)

	if result.Equal(expectedTotal) {
		t.Errorf("Expected %v, got %v\n", result, expectedTotal)
	}
}

func TestExpenseIsEvenlyShared(t *testing.T) {
	// expense3 is evenly shared
	expected := true
	result := ExpenseIsEvenlyShared(&expense3)
	if expected != result {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}

	// expense1 isn't evenly shared
	expected = false
	result = ExpenseIsEvenlyShared(&expense1)
	if expected != result {
		t.Errorf("Expected %v, got %v\n", expected, result)
	}

	// expense4 is evenly shared but with the caveat of fractional cents
	expected = true
	result = ExpenseIsEvenlyShared(&expense4)
	if expected != result {
		fmt.Printf("Expense: %+v\n", expense4)
		t.Errorf("Expected %v, got %v\n", expected, result)
	}
}
