package pages

import (
	"expenses/expenses"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func expenseFromForm(c *gin.Context) (*expenses.Expense, error) {
	newDescription := c.PostForm("expense-desc")
	newDate := c.PostForm("expense-date")

	formattedDate, err := time.Parse("02-Jan-2006", newDate)
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
	typID, err := strconv.Atoi(newTyp)
	if err != nil {
		return nil, err
	}

	newCat := c.PostForm("newexp-cat-dropdown")
	catID, err := strconv.Atoi(newCat)
	if err != nil {
		return nil, err
	}

	newStore := c.PostForm("newexp-store-dropdown")
	storeID, err := strconv.Atoi(newStore)
	if err != nil {
		return nil, err
	}

	shares := []expenses.ExpenseShare{}
	payments := []expenses.ExpensePayment{}

	sharesUserIDs := c.PostFormArray("shares-user-ids[]")
	formShares := c.PostFormArray("shares-percent[]")
	formPayments := c.PostFormArray("payments-payment[]")
	
	for i := range sharesUserIDs {
		userID, err := strconv.Atoi(sharesUserIDs[i])
		if err != nil {
			return nil, err
		}

		share, err := strconv.ParseFloat(formShares[i], 32)
		if err != nil {
			return nil, err
		}

		payed, err := strconv.ParseFloat(formPayments[i], 32)
		if err != nil {
			return nil, err
		}

		newShare := expenses.ExpenseShare{
			User: expenses.User{
				UserID: userID,
			},
			Share: float32(share),
		}	
		shares = append(shares, newShare)

		newPayment := expenses.ExpensePayment{
			User: expenses.User{
				UserID: userID,
			},
			PayedAmount: float32(payed),
		} 
		payments = append(payments, newPayment)
	}

	newExp := expenses.Expense{
		Description: newDescription,
		ExpDate:     date,
		Value:       float32(value),
		ExpType: expenses.Type{
			TypeID: typID,
		},
		ExpCategory: expenses.Category{
			CategoryID: catID,
		},
		ExpStore: expenses.Store{
			StoreID: storeID,
		},
		Shares: shares,
		Payments: payments,
	}
	return &newExp, nil
}
