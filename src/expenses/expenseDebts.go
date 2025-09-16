package expenses

import (
	"context"
	"log"
	"sort"

	mod "github.com/ESilva15/expenses/expenses/models"

	dec "github.com/shopspring/decimal"
)

// userTab describes the tab a given user has.
type userTab struct {
	User mod.User
	Sum  dec.Decimal
}

// userTabs is a list of UserTab.
type userTabs []userTab

// SortBySum sorts the list of UserTab by the attrib Sum.
func (ut userTabs) SortBySum() {
	sort.Slice(ut, func(i, j int) bool {
		return ut[i].Sum.Cmp(ut[j].Sum) > 0
	})
}

// Equal returns the equality of the passed UserTab.
func (ut userTabs) Equal(t userTabs) bool {
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

func filterExpenseParticipants(e *mod.Expense) (userTabs, userTabs) {
	shares := mapShares(e)
	payments := mapPayments(e)

	debtors := userTabs{}
	creditors := userTabs{}

	for user, share := range shares {
		debt := share.Sub(payments[user])
		if debt.LessThan(dec.NewFromFloat(0.0)) {
			creditors = append(creditors, userTab{
				User: user,
				Sum:  debt.Abs(),
			})
		}

		if debt.GreaterThan(dec.NewFromFloat(0.0)) {
			debtors = append(debtors, userTab{
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
// handle doing the accounting while collecting the data instead.
type compoundeKey struct {
	Debtor   mod.User
	Creditor mod.User
}

func resolveDebt(creditor userTab, debtors userTabs) mod.Debts {
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

func resolveDebts(debtors userTabs, creditors userTabs) mod.Debts {
	keyedDebts := make(map[compoundeKey]dec.Decimal)

	for _, creditor := range creditors {
		debt := resolveDebt(creditor, debtors)

		for _, d := range debt {
			key := compoundeKey{Debtor: d.Debtor, Creditor: d.Creditor}
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

// TODO - Are we sure there are no errors in here?
func CalculateDebts(e *mod.Expense) (mod.Debts, error) {
	debtors, creditors := filterExpenseParticipants(e)
	log.Println("Debtors:", debtors)
	log.Println("Creditors:", creditors)

	debts := resolveDebts(debtors, creditors)
	log.Println("Debts:", debts)

	return debts, nil
}

// settleDebt updates a debt settlement on the database
// it takes away the payed debt by payment from the creditor payments
// TODO
// This won't work if the payment was made with multiple payments - but for my
// use case that won't be a thing. Right now every expense I introduce has been
// fully payed in one go
func (a *ExpApp) settleDebt(payment mod.Payment,
	credit mod.Payment, eID int32) error {
	ctx := context.Background()

	credit.PayedAmount = credit.PayedAmount.Sub(payment.PayedAmount)

	return a.ExpRepo.SettleDebt(ctx, eID, payment, credit)
}

func (a *ExpApp) ProcessDebt(expID int32, debt mod.Debt) error {
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
