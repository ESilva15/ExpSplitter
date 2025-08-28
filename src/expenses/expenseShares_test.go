package expenses

import (
	mod "expenses/expenses/models"
	dec "github.com/shopspring/decimal"
	"testing"
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
	app := ExpensesApp{}

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
