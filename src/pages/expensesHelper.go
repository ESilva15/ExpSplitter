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
	}
	return &newExp, nil
}
