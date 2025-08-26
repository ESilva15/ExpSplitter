package expenses

import (
	mod "expenses/expenses/models"
	"testing"

	dec "github.com/shopspring/decimal"
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
		Value: dec.NewFromInt(150),
		Shares: []mod.Share{
			{ExpShareID: 0, User: user1, Share: dec.NewFromFloat(0.4)},
			{ExpShareID: 1, User: user2, Share: dec.NewFromFloat(0.4)},
			{ExpShareID: 2, User: user3, Share: dec.NewFromFloat(0.2)},
		},
		Payments: []mod.ExpensePayment{
			{ExpPaymID: 0, User: user1, PayedAmount: dec.NewFromInt(50)},
			{ExpPaymID: 1, User: user1, PayedAmount: dec.NewFromInt(50)},
			{ExpPaymID: 2, User: user2, PayedAmount: dec.NewFromInt(0)},
			{ExpPaymID: 3, User: user3, PayedAmount: dec.NewFromInt(30)},
			{ExpPaymID: 4, User: user3, PayedAmount: dec.NewFromInt(20)},
		},
	}

	// Expense isn't evenly shared
	// Expense is paid off
	expense2 = mod.Expense{
		ExpID: 0,
		Value: dec.NewFromInt(160),
		Shares: []mod.Share{
			{ExpShareID: 0, User: user1, Share: dec.NewFromFloat(0.3)},
			{ExpShareID: 1, User: user2, Share: dec.NewFromFloat(0.3)},
			{ExpShareID: 2, User: user3, Share: dec.NewFromFloat(0.3)},
			{ExpShareID: 3, User: user4, Share: dec.NewFromFloat(0.1)},
		},
		Payments: []mod.ExpensePayment{
			{ExpPaymID: 0, User: user1, PayedAmount: dec.NewFromInt(40)},
			{ExpPaymID: 1, User: user1, PayedAmount: dec.NewFromInt(40)},
			{ExpPaymID: 2, User: user2, PayedAmount: dec.NewFromInt(5)},
			{ExpPaymID: 3, User: user3, PayedAmount: dec.NewFromInt(30)},
			{ExpPaymID: 4, User: user3, PayedAmount: dec.NewFromInt(30)},
			{ExpPaymID: 5, User: user3, PayedAmount: dec.NewFromInt(10)},
			{ExpPaymID: 6, User: user4, PayedAmount: dec.NewFromInt(4)},
			{ExpPaymID: 7, User: user4, PayedAmount: dec.NewFromInt(1)},
		},
	}
	/*
		Shares:
		User1 0.3
		User2 0.3
		User3 0.3
		User4 0.1

		Payments  | Debts
		User1  80 | (0.3 * 160) - 80 = -32 (is owed 32)
		User2   5 | (0.3 * 160) -  5 =  43 (owes 43)
		User3  70 | (0.3 * 160) - 70 = -22 (is owed 22)
		User4   5 | (0.1 * 160) -  5 =  11 (owes 11)

		Debtors   {u4, 11}, {u2, 43},
		Creditors {u3, 22}, {u1, 32}

		{u3, 22} - {u4, 11} = {u3, 11}, {u4, 00}
		{u3, 11} - {u2, 43} = {u3, 00}, {u2, 32}
		Debts:
		{Debtor: u4, Creditor: u3, Sum: 11}
		{Debtor: u2, Creditor: u3, Sum: 11}

		Debtors   {u4, 00}, {u2, 32}
		Creditors {u3, 00}, {u1, 32}

		{u1, 32} - {u2, 32} = {u1, 00} {u2, 00}
		Debts:
		{Debtor: u2, Creditor: u1, Sum: 32}

		Final debts:
		{Debtor: u4, Creditor: u3, Sum: 11}
		{Debtor: u2, Creditor: u3, Sum: 11}
		{Debtor: u2, Creditor: u1, Sum: 32}



															 1                   2
		Debtors: [U2: 43, U4: 11] -> [U2: 21, U4: 11] -> [U4: 11]
		                               1            2
		Creditors: [U3: -22, U1: -32] -> [U1: -32] -> [U1: -11]
		Kill the debts with:
		1. User3.Credit + User2.Debt -> -22 + 43 =  21 - can pay the 22
		2. User1.Credit + User2.Debt -> -32 + 21 = -11 - can only pay 21
		3. User1.Credit + User4.Debt -> -11 + 11 =   0 - can pay the 11

		Summary:
		{creditor: U3, debtor: U2, debt: 22} - U2 owes U3 22
		{creditor: U1, debtor: U2, debt: 21} - U2 owes U1 21
		{creditor: U1, debtor: U4, debt: 11} - U4 owes U1 11
	*/

	// Expense is: evenly shared and paid off
	expense3 = mod.Expense{
		ExpID: 0,
		Owner: mod.User{
			UserID: user1.UserID,
		},
		Value: dec.NewFromInt(160),
		Shares: []mod.Share{
			{ExpShareID: 0, User: user1, Share: dec.NewFromFloat(0.3)},
			{ExpShareID: 1, User: user2, Share: dec.NewFromFloat(0.3)},
			{ExpShareID: 2, User: user3, Share: dec.NewFromFloat(0.3)},
			{ExpShareID: 3, User: user4, Share: dec.NewFromFloat(0.1)},
		},
		Payments: []mod.ExpensePayment{
			{ExpPaymID: 0, User: user1, PayedAmount: dec.NewFromInt(40)},
			{ExpPaymID: 1, User: user1, PayedAmount: dec.NewFromInt(8)},
			{ExpPaymID: 2, User: user2, PayedAmount: dec.NewFromInt(48)},
			{ExpPaymID: 3, User: user3, PayedAmount: dec.NewFromInt(15)},
			{ExpPaymID: 4, User: user3, PayedAmount: dec.NewFromInt(15)},
			{ExpPaymID: 5, User: user3, PayedAmount: dec.NewFromInt(18)},
			{ExpPaymID: 6, User: user4, PayedAmount: dec.NewFromInt(8)},
			{ExpPaymID: 7, User: user4, PayedAmount: dec.NewFromInt(8)},
		},
	}

	expense4 = mod.Expense{
		ExpID: 7,
		Value: dec.NewFromFloat(0.50),
		Shares: []mod.Share{
			{ExpShareID: 0, User: user1, Share: dec.NewFromFloat(0.34)},
			{ExpShareID: 1, User: user2, Share: dec.NewFromFloat(0.33)},
			{ExpShareID: 2, User: user3, Share: dec.NewFromFloat(0.33)},
		},
		Payments: []mod.ExpensePayment{
			{ExpPaymID: 0, User: user1, PayedAmount: dec.NewFromFloat(0.18)},
			{ExpPaymID: 1, User: user1, PayedAmount: dec.NewFromFloat(0.16)},
			{ExpPaymID: 2, User: user2, PayedAmount: dec.NewFromFloat(0.16)},
		},
	}

	expense5 = mod.Expense{
		ExpID: 9,
		Owner: mod.User{
			UserID: user1.UserID,
		},
		Value: dec.NewFromFloat(111.11),
		Shares: []mod.Share{
			{ExpShareID: 0, User: user1,
				Share: dec.NewFromFloat(0.34), Calculated: dec.NewFromFloat(37.79)},
			{ExpShareID: 1, User: user2,
				Share: dec.NewFromFloat(0.33), Calculated: dec.NewFromFloat(36.66)},
			{ExpShareID: 2, User: user3,
				Share: dec.NewFromFloat(0.33), Calculated: dec.NewFromFloat(36.66)},
		},
		Payments: []mod.ExpensePayment{
			{ExpPaymID: 0, User: user1, PayedAmount: dec.NewFromFloat(37.79)},
			{ExpPaymID: 1, User: user1, PayedAmount: dec.NewFromFloat(36.66)},
			{ExpPaymID: 2, User: user2, PayedAmount: dec.NewFromFloat(36.66)},
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
