package expenses

import (
	mod "expenses/expenses/models"
	"fmt"
	"reflect"
	"testing"

	dec "github.com/shopspring/decimal"
)

func TestNormalizeShares(t *testing.T) {
	testCases := []struct {
		name           string
		exp            *mod.Expense
		expectedShares mod.Shares
	}{
		{
			name: "Split with remainder",
			exp:  &expense5,
			expectedShares: mod.Shares{
				{ExpShareID: 0, User: user1,
					Share: dec.NewFromFloat(0.34), Calculated: dec.NewFromFloat(37.79)},
				{ExpShareID: 1, User: user2,
					Share: dec.NewFromFloat(0.33), Calculated: dec.NewFromFloat(36.66)},
				{ExpShareID: 2, User: user3,
					Share: dec.NewFromFloat(0.33), Calculated: dec.NewFromFloat(36.66)},
			},
		},
		{
			name: "Even split",
			exp:  &expense3,
			expectedShares: mod.Shares{
				{ExpShareID: 0, User: user1,
					Share: dec.NewFromFloat(0.3), Calculated: dec.NewFromFloat(48.00)},
				{ExpShareID: 1, User: user2,
					Share: dec.NewFromFloat(0.3), Calculated: dec.NewFromFloat(48.00)},
				{ExpShareID: 2, User: user3,
					Share: dec.NewFromFloat(0.3), Calculated: dec.NewFromFloat(48.00)},
				{ExpShareID: 3, User: user4,
					Share: dec.NewFromFloat(0.1), Calculated: dec.NewFromFloat(16.00)},
			},
		},
	}

	// dummy service to call it
	app := ExpApp{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := app.NormalizeShares(tc.exp)
			if err != nil {
				t.Errorf("Failed to normalize shares: %+v", err)
			}

			if !tc.expectedShares.Equal(tc.exp.Shares) {
				t.Errorf("Resulting shares: %+v\nAre not correct: %+v", tc.exp.Shares,
					tc.expectedShares)
			}
		})
	}

}

func TestMapShares(t *testing.T) {
	testCases := []struct {
		name           string
		exp            *mod.Expense
		expectedShares map[mod.User]dec.Decimal
	}{
		{
			name: "Map a share",
			exp:  &expense6,
			expectedShares: map[mod.User]dec.Decimal{
				user1: dec.NewFromFloat(0.18),
				user2: dec.NewFromFloat(0.16),
				user3: dec.NewFromFloat(0.16),
			},
		},
	}

	for _, tc := range testCases {
		fmt.Printf("Running test: %s", tc.name)

		shares := mapShares(tc.exp)
		if !reflect.DeepEqual(shares, tc.expectedShares) {
			t.Errorf("Resulting shares: %+v\nAre not correct: %+v", shares,
				tc.expectedShares)
		}
	}
}

func TestMapPayments(t *testing.T) {
	testCases := []struct {
		name             string
		exp              *mod.Expense
		expectedPayments map[mod.User]dec.Decimal
	}{
		{
			name: "Map a share",
			exp:  &expense6,
			expectedPayments: map[mod.User]dec.Decimal{
				user1: dec.NewFromFloat(0.50),
				user2: dec.NewFromFloat(0.0),
				user3: dec.NewFromFloat(0.0),
			},
		},
	}

	for _, tc := range testCases {
		fmt.Printf("Running test: %s", tc.name)

		payments := mapPayments(tc.exp)
		if !reflect.DeepEqual(payments, tc.expectedPayments) {
			t.Errorf("Resulting payments: %+v\nAre not correct: %+v", payments,
				tc.expectedPayments)
		}
	}
}
