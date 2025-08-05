package expenses

type UserExpenseSummary struct {
	User       User
	Share      float64
	PayedTotal float64
}

func NewUserExpenseSummary() UserExpenseSummary {
	return UserExpenseSummary{
		User:       NewUser(),
		Share:      0.0,
		PayedTotal: 0.0,
	}
}

func (e *Expense) GetSummary() ([]UserExpenseSummary, error) {
	summary := []UserExpenseSummary{}

	for _, share := range e.Shares {
		curSummary := NewUserExpenseSummary()
		curSummary.User = share.User
		curSummary.Share = share.Share

		for _, payment := range e.Payments {
			if payment.User.UserID == curSummary.User.UserID {
				curSummary.PayedTotal += payment.PayedAmount
			}
		}

		summary = append(summary, curSummary)
	}

	return summary, nil
}
