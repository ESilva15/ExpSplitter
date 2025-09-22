package pages

import (
	"context"
	"log"
	"time"

	expenses "github.com/ESilva15/expenses/expenses"
	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo"
	ginaux "github.com/ESilva15/expenses/ginAux"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func expenseFilterFromQuery(c *gin.Context) (repo.ExpFilter, error) {
	var err error
	var start, end time.Time
	filter := repo.NewExpFilter()

	startDate := c.Query("range-start")
	if startDate != "" {
		startDate += " 00:00:00"
		start, err = time.ParseInLocation("02-Jan-2006 15:04:05", startDate, time.UTC)
		if err != nil {
			log.Printf("error startDate: %v", err)
			return repo.ExpFilter{}, nil
		}
		filter.Start = &start
	}

	endDate := c.Query("range-end")
	if endDate != "" {
		endDate += " 00:00:00"
		end, err = time.ParseInLocation("02-Jan-2006 15:04:05", endDate, time.UTC)
		if err != nil {
			log.Printf("error endDate: %v", err)
			return repo.ExpFilter{}, nil
		}
		filter.End = &end
	}

	filter.CatIDs, err = ginaux.QueryIntArray(c, "cat-dropdown[]")
	if err != nil {
		log.Printf("Failed to retrieve IDs from cat-dropdown[] %+v", err)
		return repo.ExpFilter{}, nil
	}

	filter.StoreIDs, err = ginaux.QueryIntArray(c, "store-dropdown[]")
	if err != nil {
		log.Printf("Failed to retrieve IDs from store-dropdown[] %+v", err)
		return repo.ExpFilter{}, nil
	}

	filter.TypeIDs, err = ginaux.QueryIntArray(c, "type-dropdown[]")
	if err != nil {
		log.Printf("Failed to retrieve IDs from type-dropdown[] %+v", err)
		return repo.ExpFilter{}, nil
	}

	return filter, nil
}

func processFormShares(c *gin.Context) ([]mod.Share, error) {
	userIDs := c.PostFormArray("shares-user-ids[]")
	shareIDs := c.PostFormArray("shares-ids[]")
	shares := c.PostFormArray("shares-percent[]")

	return expenses.ParseFormShares(userIDs, shares, shareIDs)
}

func processFormPayments(c *gin.Context) ([]mod.Payment, error) {
	userIDs := c.PostFormArray("payments-user-ids[]")
	paymentIDs := c.PostFormArray("payments-ids[]")
	values := c.PostFormArray("payments-payment[]")

	return expenses.ParseFormPayments(userIDs, paymentIDs, values)
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
