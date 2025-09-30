package pages

import (
	"context"
	"fmt"
	"time"

	exp "github.com/ESilva15/expenses/expenses"
	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo"
	gaux "github.com/ESilva15/expenses/ginAux"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func processFormShares(c *gin.Context) ([]mod.Share, error) {
	userIDs := c.PostFormArray("shares-user-ids[]")
	shareIDs := c.PostFormArray("shares-ids[]")
	shares := c.PostFormArray("shares-percent[]")

	return exp.ParseFormShares(userIDs, shares, shareIDs)
}

func processFormPayments(c *gin.Context) ([]mod.Payment, error) {
	userIDs := c.PostFormArray("payments-user-ids[]")
	paymentIDs := c.PostFormArray("payments-ids[]")
	values := c.PostFormArray("payments-payment[]")

	return exp.ParseFormPayments(userIDs, paymentIDs, values)
}

func expenseFromForm(ctx context.Context, c *gin.Context) (*mod.Expense, error) {
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

	typID, err := exp.ParseID(c.PostForm("newexp-type-dropdown"))
	if err != nil {
		return nil, err
	}

	catID, err := exp.ParseID(c.PostForm("newexp-cat-dropdown"))
	if err != nil {
		return nil, err
	}

	// TODO:
	// Move the names of the dropdowns to a variable that can be passed to the
	// html page
	storeID, err := exp.ParseID(c.PostForm("newexp-store-dropdown"))
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

func fetchExpensesData(c *gin.Context) (map[string]any, error) {
	ctx, err := gaux.GetLoggedInUserCTX(c)
	if err != nil {
		return nil, fmt.Errorf("could not get logged in user - %v", err)
	}

	eFilter := repo.NewExpFilter()
	expenses, err := exp.App.GetAllExpenses(ctx, eFilter)
	if err != nil {
		return nil, fmt.Errorf("could not get expenses - %v", err)
	}

	categories, err := exp.App.GetAllCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get categories - %v", err)
	}

	stores, err := exp.App.GetAllStores(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get stores - %v", err)
	}

	types, err := exp.App.GetAllTypes(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get types - %v", err)
	}

	return map[string]any{
		"expenses":   expenses,
		"categories": categories,
		"stores":     stores,
		"types":      types,
	}, nil
}
