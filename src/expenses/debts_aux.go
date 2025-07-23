package expenses

func userPayments(e *Expense) map[User]float32 {
	userPayments := make(map[User]float32)

	for _, payment := range e.Payments {
		userPayments[payment.User] += payment.PayedAmount
	}

	return userPayments
}

func userShares(e *Expense) map[User]float32 {
	userShares := make(map[User]float32)

	for _, share := range e.Shares {
		userShares[share.User] = share.Share
	}

	return userShares
}
