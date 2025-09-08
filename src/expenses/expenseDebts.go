package expenses

import (
	"context"
	mod "expenses/expenses/models"
	"sort"

	dec "github.com/shopspring/decimal"
)

type UserTab struct {
	User mod.User
	Sum  dec.Decimal
}
type UserTabs []UserTab

func (ut UserTabs) SortBySum() {
	sort.Slice(ut, func(i, j int) bool {
		return ut[i].Sum.Cmp(ut[j].Sum) > 0
	})
}

func (ut UserTabs) Equal(t UserTabs) bool {
	if len(ut) != len(t) {
		return false
	}

	for k := range ut {
		if ut[k].User != t[k].User || !ut[k].Sum.Equal(t[k].Sum) {
			return false
		}
	}

	return true
}

func filterExpenseParticipants(e *mod.Expense) (UserTabs, UserTabs) {
	shares := mapShares(e)
	payments := mapPayments(e)

	debtors := UserTabs{}
	creditors := UserTabs{}

	for user, share := range shares {
		debt := share.Sub(payments[user])
		if debt.LessThan(dec.NewFromFloat(0.0)) {
			creditors = append(creditors, UserTab{
				User: user,
				Sum:  debt.Abs(),
			})
		}

		if debt.GreaterThan(dec.NewFromFloat(0.0)) {
			debtors = append(debtors, UserTab{
				User: user,
				Sum:  debt.Abs(),
			})
		}
	}

	debtors.SortBySum()
	creditors.SortBySum()

	return debtors, creditors
}

// TODO
// Instead of doing this impererively (I think?) create a struct that can
// handle doing the accounting while collecting the data instead
type CompoundKey struct {
	Debtor   mod.User
	Creditor mod.User
}

func resolveDebt(creditor UserTab, debtors UserTabs) mod.Debts {
	debts := mod.Debts{}

	credit := creditor.Sum
	for k := range debtors {
		if credit.IsZero() {
			break
		}

		debt := mod.Debt{
			Creditor: creditor.User,
			Debtor:   debtors[k].User,
			Sum:      dec.NewFromFloat(0.0),
		}

		if debtors[k].Sum.GreaterThanOrEqual(credit) {
			debt.Sum = credit
		} else {
			debt.Sum = debtors[k].Sum
		}
		credit = credit.Sub(debt.Sum)
		debtors[k].Sum = debtors[k].Sum.Sub(debt.Sum)

		debts = append(debts, debt)
	}

	return debts
}

func resolveDebts(debtors UserTabs, creditors UserTabs) mod.Debts {
	keyedDebts := make(map[CompoundKey]dec.Decimal)

	for _, creditor := range creditors {
		debt := resolveDebt(creditor, debtors)

		for _, d := range debt {
			key := CompoundKey{Debtor: d.Debtor, Creditor: d.Creditor}
			if _, ok := keyedDebts[key]; ok {
				keyedDebts[key] = keyedDebts[key].Add(d.Sum)
			} else {
				keyedDebts[key] = d.Sum
			}
		}
	}

	debts := mod.Debts{}
	for key, debt := range keyedDebts {
		newDebt := mod.Debt{
			Debtor:   key.Debtor,
			Creditor: key.Creditor,
			Sum:      debt.Truncate(2), // Removes fractional cents from the debt
		}

		// Only appends the debt if its not zero
		if !newDebt.Sum.IsZero() {
			debts = append(debts, newDebt)
		}
	}

	return debts
}

// TODO
// Are we sure there are no errors in here?
func CalculateDebts(e *mod.Expense) (mod.Debts, error) {
	debtors, creditors := filterExpenseParticipants(e)
	debts := resolveDebts(debtors, creditors)

	return debts, nil
}

// settleDebt updates a debt settlement on the database
// it takes away the payed debt by payment from the creditor payments
// TODO
// This won't work if the payment was made with multiple payments - but for my
// use case that won't be a thing. Right now every expense I introduce has been
// fully payed in one go
func (a *ExpensesApp) settleDebt(payment mod.Payment,
	credit mod.Payment, eId int32) error {
	ctx := context.Background()

	credit.PayedAmount = credit.PayedAmount.Sub(payment.PayedAmount)

	return a.ExpRepo.SettleDebt(ctx, eId, payment, credit)
}

func (a *ExpensesApp) ProcessDebt(expID int32, debt mod.Debt) error {
	// We have to balance the debtor and creditor payments
	payment := mod.Payment{
		User: mod.User{
			UserID: debt.Debtor.UserID,
		},
		PayedAmount: debt.Sum,
	}

	creditorPayment, err := a.GetExpensePaymentByUserID(expID, debt.Creditor.UserID)
	if err != nil {
		return err
	}

	err = a.settleDebt(payment, creditorPayment, expID)
	if err != nil {
		return err
	}

	return nil
}
