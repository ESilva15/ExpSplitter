package pages

import (
	"log"
	"net/http"

	exp "github.com/ESilva15/expenses/expenses"
	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/gin-gonic/gin"
)

func getSummaryData(c *gin.Context) ([]mod.MonthlySummary, error) {
	ctx, err := getLoggedInUserCTX(c)
	if err != nil {
		log.Printf("error getting user from context: %+v", err)
		ServerErrorView(c, "Could not get logged in user")
		return nil, nil
	}

	filter, err := filterFromQuery(c)
	if err != nil {
		log.Printf("error parsing filters from query: %+v", err)
		ServerErrorView(c, "Could not get logged in user")
		return nil, nil
	}
	filter.UserIDs = *&ctx.Value("user").(*mod.User).UserID
	summary, err := exp.App.GetMonthlySummary(ctx, &filter)
	if err != nil {
		log.Printf("error retrieving monthly summary from DB: %+v", err)
		c.Header("HX-Redirect", "/500")
		return nil, err
	}

	return summary, nil
}

func analysisDashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "analysis-filtered",
		"renderNavBar": true,
		"content":      "analysis-filtered",
	})
}

func analysisDashboardPagePartial(c *gin.Context) {
	c.HTML(http.StatusOK, "analysis-filtered", gin.H{
		// "plotData": data,
	})
}

func analysisDashboardQuery(c *gin.Context) {
	data, err := getSummaryData(c)
	if err != nil {
		log.Printf("error: %+v", err)
		ServerErrorView(c, "failed to compile analysis data")
		return
	}

	c.HTML(http.StatusOK, "analysis-filtered-results", data)
}

// RouteAnalysis routes the endpoints for analysis stuff.
func RouteAnalysisDashboard(router *gin.RouterGroup) {
	router.GET("/filtered-analysis", analysisDashboardPage)
	router.GET("/filtered-analysis/filters", analysisDashboardQuery)
	router.POST("/filtered-analysis", analysisDashboardPagePartial)
}
