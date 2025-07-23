package expenses

type Debt struct {
	Creditor User
	Debtor   User
	Sum      float32
}

func sortBySum(a, b Debt) int {
	if a.Sum < b.Sum {
		return -1
	}
	if a.Sum > b.Sum {
		return 1
	}
	return 0
}

func (e *Expense) calculateDebts() ([]Debt, error) {
	// debts := []Debt{}
	//
	// if len(e.Shares) <= 1 {
	// 	return debts, nil
	// }
	//
	// if len(e.Shares) > 1 {
	// 	userShares := userShares(e)
	// 	userPayments := userPayments(e)
	//
	// 	for user := range userPayments {
	// 		owed := (userShares[user] * e.Value) - userPayments[user]
	// 		if owed > 0 {
	// 			debts = append(debts, Debt{
	// 				Creditor: 
	// 			})
	// 		}
	// 	}
	// }

	return []Debt{
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
	}, nil
}
