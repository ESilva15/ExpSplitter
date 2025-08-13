package expenses

import (
	// mod "expenses/expenses/models"
	// "slices"
	// "testing"
	//
	// "github.com/shopspring/decimal"
	// "github.com/stretchr/testify/assert"
)

/*
The debt calculation is pretty much the core functionality of this thing
Formula to calculate how much a given user owes:
(share * total) - payedAmount

TestCase1
A simple expense of 150 with:
Shares:
User1 0.5
User2 0.5

Payments  | Debts
User1 135 | (0.5 * 150) - 135 = -60 (is owed 60)
User2  15 | (0.5 * 150) - 15  =  60 (owes 60)

Debtors:   [User2: 60]
Credtiors: [User1: -60]
Kill the debts with:
User1.Debt + User2.Credit = 0 {creditor: User1, debtor: User2, debt: 60}
State Update:
Debtors:   [User2: 0]
Credtiors: [User1: 0]

Summary:
User2 owes User1 60

TestCase2
An expense of 160 with:
Shares:
User1 0.3
User2 0.3
User3 0.3
User4 0.1

Payments  | Debts
User1  80 | (0.5 * 160) - 80 = -32 (is owed 60)
User2   5 | (0.5 * 160) -  5 =  43 (owes 43)
User3  70 | (0.5 * 160) - 70 = -22 (is owed 22)
User4   5 | (0.5 * 160) -  5 =  11 (owes 11)

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

// func TestCalculateDebtsTestCase1(t *testing.T) {
// 	expense := mod.Expense{
// 		ExpID:       1,
// 		Description: "Test Exp",
// 		Value:       decimal.NewFromFloat(150.0),
// 		Type: mod.Type{
// 			TypeID:   1,
// 			TypeName: "Despesa",
// 		},
// 		Category: mod.Category{
// 			CategoryID:   0,
// 			CategoryName: "Groceries",
// 		},
// 		Owner: mod.User{
// 			UserID:   1,
// 			UserName: "ESilva",
// 		},
// 		Date: 1753307982,
// 		Payments: []mod.ExpensePayment{
// 			{
// 				ExpPaymID: 0,
// 				User: mod.User{
// 					UserID:   1,
// 					UserName: "ESilva",
// 				},
// 				PayedAmount: decimal.NewFromFloat(135.0),
// 			},
// 			{
// 				ExpPaymID: 1,
// 				User: mod.User{
// 					UserID:   2,
// 					UserName: "Kika",
// 				},
// 				PayedAmount: decimal.NewFromFloat(15.0),
// 			},
// 		},
// 		Shares: []mod.ExpenseShare{
// 			{
// 				ExpShareID: 0,
// 				User: mod.User{
// 					UserID:   1,
// 					UserName: "ESilva",
// 				},
// 				Share: decimal.NewFromFloat(0.5),
// 			},
// 			{
// 				ExpShareID: 1,
// 				User: mod.User{
// 					UserID:   2,
// 					UserName: "Kika",
// 				},
// 				Share: decimal.NewFromFloat(0.5),
// 			},
// 		},
// 	}
//
// 	expectedDebts := []Debt{
// 		{
// 			Debtor: mod.User{
// 				UserID:   2,
// 				UserName: "Kika",
// 			},
// 			Sum: decimal.NewFromInt(60),
// 		},
// 	}
// 	slices.SortFunc(expectedDebts, sortBySum)
//
// 	resultDebts, err := calculateDebts()
// 	if err != nil {
// 		t.Fatalf("Debt calculation failed: %+v", err)
// 	}
// 	slices.SortFunc(resultDebts, sortBySum)
//
// 	assert.Equal(t, expectedDebts, resultDebts)
// }
