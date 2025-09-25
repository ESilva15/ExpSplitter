package pages

import (
	"net/http"
	"sync"

	exp "github.com/ESilva15/expenses/expenses"
	mod "github.com/ESilva15/expenses/expenses/models"
	val "github.com/ESilva15/expenses/expenses/values"
	dec "github.com/shopspring/decimal"

	"github.com/gin-gonic/gin"
)

// MonthlyTotal holds the data for the MonthlyTotal plot.
type MonthlyTotal struct {
	Month int         `json:"month"`
	Total dec.Decimal `json:"total"`
}

func getMonthlyBalance(expenses mod.Expenses, uID int32) ([]MonthlyTotal, error) {
	monthlyBalance := make([]MonthlyTotal, val.MonthsPerYear)
	for k := range val.MonthsPerYear {
		monthlyBalance[k].Month = k
	}

	for _, e := range expenses {
		month := e.Date.Month() - 1
		t := e.Type.TypeName

		// Get the current expenses value and make it negative if its an expense
		// It's not the value, its what we paid on it (our payments)
		// I don't like this like this - takes too long, we should load it from DB
		// at the start, and cache it too
		_ = exp.App.LoadExpensePayments(&e)
		value := exp.App.PaidByUser(e, uID)
		// value := e.Value
		if t == "Despesa" {
			value = value.Neg()
		}

		// Store the data on the montly balance list
		monthlyBalance[month].Total = monthlyBalance[month].Total.Add(value)
	}

	return monthlyBalance, nil
}

func getMonthlyTotal(expenses mod.Expenses, uID int32) ([]MonthlyTotal, error) {
	mTotal := make([]MonthlyTotal, val.MonthsPerYear)
	for k := range mTotal {
		mTotal[k].Month = k
	}

	// Step 1: net total for each month
	for _, e := range expenses {
		month := int(e.Date.Month()) - 1 // Jan = 0
		// Same as on the other one here
		_ = exp.App.LoadExpensePayments(&e)
		value := exp.App.PaidByUser(e, uID)
		if e.Type.TypeName == "Despesa" {
			value = value.Neg()
		}
		mTotal[month].Total = mTotal[month].Total.Add(value)
	}

	// Step 2: cumulative balance
	for i := 1; i < len(mTotal); i++ {
		mTotal[i].Total = mTotal[i].Total.Add(mTotal[i-1].Total)
	}

	return mTotal, nil
}

func getAnalysisDataWG(c *gin.Context) (map[string]any, error) {
	ctx, err := getLoggedInUserCTX(c)
	if err != nil {
		return map[string]any{}, err
	}
	user := *ctx.Value("user").(*mod.User)

	filter, err := expenseFilterFromQuery(c)
	if err != nil {
		return map[string]any{}, err
	}

	queriedExpenses, err := exp.App.GetAllExpenses(ctx, filter)
	if err != nil {
		return map[string]any{}, err
	}

	var (
		mBalance []MonthlyTotal
		mTotal   []MonthlyTotal
		bErr     error
		tErr     error
		wg       sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		mBalance, bErr = getMonthlyBalance(queriedExpenses, user.UserID)
	}()
	go func() {
		defer wg.Done()
		mTotal, tErr = getMonthlyTotal(queriedExpenses, user.UserID)
	}()
	wg.Wait()

	if tErr != nil {
		return map[string]any{}, tErr
	}
	if bErr != nil {
		return map[string]any{}, bErr
	}

	plotData := map[string]any{
		"mBalance": mBalance,
		"mTotal":   mTotal,
	}

	return plotData, nil
}

func analysisPartialPage(c *gin.Context) {
	// TODO - this is repeated code with the analysisPage one, figure something
	data, err := getAnalysisDataWG(c)
	if err != nil {
		ServerErrorView(c, "failed to compiled analysis data")
		return
	}

	c.HTML(http.StatusOK, "analysis", gin.H{
		"plotData": data,
	})
}

func analysisPage(c *gin.Context) {
	data, err := getAnalysisDataWG(c)
	if err != nil {
		ServerErrorView(c, "failed to compiled analysis data")
		return
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "analysis",
		"renderNavBar": true,
		"content":      "analysis",
		"plotData":     data,
	})
}

// RouteAnalysis routes the endpoints for analysis stuff.
func RouteAnalysis(router *gin.RouterGroup) {
	router.GET("/analysis", analysisPage)
	router.POST("/analysis", analysisPartialPage)
}
