package expenses

import (
	"reflect"
	"testing"
)

var (
	user1 = User{
		UserID:   1,
		UserName: "Gustavo Gomes",
	}
	user2 = User{
		UserID:   2,
		UserName: "Paulo Fultre",
	}
	user3 = User{
		UserID:   3,
		UserName: "Afonso Esteves",
	}

	expense1 = Expense{
		ExpID: 0,
		Value: 150,
		Shares: []ExpenseShare{
			{
				ExpShareID: 0,
				User:       user1,
				Share:      0.4,
			},
			{
				ExpShareID: 1,
				User:       user2,
				Share:      0.4,
			},
			{
				ExpShareID: 2,
				User:       user3,
				Share:      0.2,
			},
		},
		Payments: []ExpensePayment{
			{
				ExpPaymID:   0,
				User:        user1,
				PayedAmount: 50,
			},
			{
				ExpPaymID:   1,
				User:        user1,
				PayedAmount: 50,
			},
			{
				ExpPaymID:   2,
				User:        user2,
				PayedAmount: 0,
			},
			{
				ExpPaymID:   3,
				User:        user3,
				PayedAmount: 40,
			},
			{
				ExpPaymID:   4,
				User:        user3,
				PayedAmount: 20,
			},
		},
	}
)

func TestUserPayments(t *testing.T) {
	expected := map[User]float32{
		user1: 100,
		user2: 0,
		user3: 60,
	}

	result := userPayments(&expense1)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected payments total and result is different:\n%+v\n%+v\n",
			expected, result)
	}
}

func TestUserShares(t *testing.T) {
	expected := map[User]float32{
		user1: 0.4,
		user2: 0.4,
		user3: 0.2,
	}

	result := userShares(&expense1)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected user shares and result is different:\n%+v\n%+v\n",
			expected, result)
	}
}
