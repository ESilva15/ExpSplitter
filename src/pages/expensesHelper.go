package pages

import (
	"expenses/expenses"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func processFormShares(c *gin.Context) ([]expenses.ExpenseShare, error) {
	shares := []expenses.ExpenseShare{}

	sharesUserIDs := c.PostFormArray("shares-user-ids[]")
	formSharesIDs := c.PostFormArray("shares-ids[]")
	formShares := c.PostFormArray("shares-percent[]")

	for i := range sharesUserIDs {
		userID, err := strconv.ParseInt(sharesUserIDs[i], 10, 16)
		if err != nil {
			return nil, err
		}

		share, err := strconv.ParseFloat(formShares[i], 32)
		if err != nil {
			return nil, err
		}

		newShare := expenses.ExpenseShare{
			ExpShareID: -1,
			User: expenses.User{
				UserID: userID,
			},
			Share: share,
		}

		if formSharesIDs[i] != "" {
			id, err := strconv.ParseInt(formSharesIDs[i], 10, 16)
			if err != nil {
				return nil, err
			}
			newShare.ExpShareID = id
		}

		shares = append(shares, newShare)
	}

	return shares, nil
}

func processFormPayments(c *gin.Context) ([]expenses.ExpensePayment, error) {
	payments := []expenses.ExpensePayment{}

	paymentsUserIDs := c.PostFormArray("payments-user-ids[]")
	formPaymentsIDs := c.PostFormArray("payments-ids[]")
	formPaymentsValues := c.PostFormArray("payments-payment[]")

	for k := range paymentsUserIDs {
		userID, err := strconv.ParseInt(paymentsUserIDs[k], 10, 16)
		if err != nil {
			return nil, err
		}

		payed, err := strconv.ParseFloat(formPaymentsValues[k], 32)
		if err != nil {
			return nil, err
		}

		newPayment := expenses.ExpensePayment{
			ExpPaymID: -1,
			User: expenses.User{
				UserID: userID,
			},
			PayedAmount: payed,
		}

		if formPaymentsIDs[k] != "" {
			id, err := strconv.ParseInt(formPaymentsIDs[k], 10, 16)
			if err != nil {
				return nil, err
			}
			newPayment.ExpPaymID = id
		}

		payments = append(payments, newPayment)
	}

	return payments, nil
}

func expenseFromForm(c *gin.Context) (*expenses.Expense, error) {
	newDescription := c.PostForm("expense-desc")
	newDate := c.PostForm("expense-date")

	formattedDate, err := time.ParseInLocation("02-Jan-2006", newDate, time.UTC)
	if err != nil {
		return nil, err
	}
	date := formattedDate.Unix()

	newValue := c.PostForm("expense-value")
	value, err := strconv.ParseFloat(newValue, 32)
	if err != nil {
		return nil, err
	}

	newTyp := c.PostForm("newexp-type-dropdown")
	typID, err := strconv.ParseInt(newTyp, 10, 16)
	if err != nil {
		return nil, err
	}

	newCat := c.PostForm("newexp-cat-dropdown")
	catID, err := strconv.ParseInt(newCat, 10, 16)
	if err != nil {
		return nil, err
	}

	newStore := c.PostForm("newexp-store-dropdown")
	storeID, err := strconv.ParseInt(newStore, 10, 16)
	if err != nil {
		return nil, err
	}

	payments, err := processFormPayments(c)
	if err != nil {
		return nil, err
	}
	shares, err := processFormShares(c)
	if err != nil {
		return nil, err
	}

	newExp := expenses.Expense{
		Description: newDescription,
		Date:        date,
		Value:       value,
		Type: expenses.Type{
			TypeID: typID,
		},
		Category: expenses.Category{
			CategoryID: catID,
		},
		Store: expenses.Store{
			StoreID: storeID,
		},
		Owner: expenses.User{
			// TODO
			// This needs to be dynamic - to be added once we have logins and whatnot
			UserID: 1,
		},
		CreationDate: time.Now().Unix(),
		Shares:       shares,
		Payments:     payments,
	}
	return &newExp, nil
}
