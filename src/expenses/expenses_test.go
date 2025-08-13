package expenses

import (
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
