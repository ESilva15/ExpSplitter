package pages

import (
	"context"
	"expenses/expenses"
	mod "expenses/expenses/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func processFormShares(c *gin.Context) ([]mod.Share, error) {
	userIDS := c.PostFormArray("shares-user-ids[]")
	shareIDs := c.PostFormArray("shares-ids[]")
	shares := c.PostFormArray("shares-percent[]")

	return expenses.ParseFormShares(userIDS, shares, shareIDs)
}

func processFormPayments(c *gin.Context) ([]mod.Payment, error) {
	userIDs := c.PostFormArray("payments-user-ids[]")
	paymentIDs := c.PostFormArray("payments-ids[]")
	values := c.PostFormArray("payments-payment[]")

	return expenses.ParseFormPayments(userIDs, paymentIDs, values)
}

func expenseFromForm(c *gin.Context, ctx context.Context) (*mod.Expense, error) {
	newDescription := c.PostForm("expense-desc")
	newDate := c.PostForm("expense-date")
	newQR := c.PostForm("expense-qr")

	formattedDate, err := time.ParseInLocation("02-Jan-2006", newDate, time.UTC)
	if err != nil {
		return nil, err
	}

	newValue := c.PostForm("expense-value")
	value, err := decimal.NewFromString(newValue)
	if err != nil {
		return nil, err
	}

	typID, err := expenses.ParseID(c.PostForm("newexp-type-dropdown"))
	if err != nil {
		return nil, err
	}

	catID, err := expenses.ParseID(c.PostForm("newexp-cat-dropdown"))
	if err != nil {
		return nil, err
	}

	// TODO:
	// Move the names of the dropdowns to a variable that can be passed to the
	// html page
	storeID, err := expenses.ParseID(c.PostForm("newexp-store-dropdown"))
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

	creator := *ctx.Value("user").(*mod.User)
	return &mod.Expense{
		Description: newDescription,
		Date:        formattedDate,
		Value:       value,
		Type: mod.Type{
			TypeID: typID,
		},
		Category: mod.Category{
			CategoryID: catID,
		},
		Store: mod.Store{
			StoreID: storeID,
		},
		Owner:        creator,
		CreationDate: time.Now(),
		Shares:       shares,
		Payments:     payments,
		QRString:     newQR,
	}, nil
}
