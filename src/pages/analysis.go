package pages

import (
	"net/http"

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

	for i := range queriedExpenses {
		_ = exp.App.LoadExpensePayments(&queriedExpenses[i])
	}

	mBalance := make([]MonthlyTotal, val.MonthsPerYear)
	for k := range val.MonthsPerYear {
		mBalance[k].Month = k
	}

	for _, e := range queriedExpenses {
		month := int(e.Date.Month()) - 1 // Jan = 0
		t := e.Type.TypeName

		value := exp.App.PaidByUser(e, user.UserID)
		if t == "Despesa" {
			value = value.Neg()
		}

		mBalance[month].Total = mBalance[month].Total.Add(value)
	}

	mTotal := make([]MonthlyTotal, len(mBalance))
	copy(mTotal, mBalance)
	for i := 1; i < len(mTotal); i++ {
		mTotal[i].Total = mTotal[i].Total.Add(mTotal[i-1].Total)
	}

	plotData := map[string]any{
		"mBalance": mBalance,
		"mTotal":   mTotal,
	}

	return plotData, nil
}

func analysisPartialPage(c *gin.Context) {
	data, err := getAnalysisDataWG(c)
	if err != nil {
		ServerErrorView(c, "failed to compile analysis data")
		return
	}

	c.HTML(http.StatusOK, "analysis", gin.H{
		"plotData": data,
	})
}

func analysisPage(c *gin.Context) {
	data, err := getAnalysisDataWG(c)
	if err != nil {
		ServerErrorView(c, "failed to compile analysis data")
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
