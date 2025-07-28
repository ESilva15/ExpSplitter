package pages

import (
	"expenses/config"
	"expenses/expenses"
	"expenses/templating"
	"log"
	"time"

	"html/template"
	"net/http"
	fp "path/filepath"

	"github.com/gin-gonic/gin"
)

type ExpenseDebtOverview struct {
	Exp     *expenses.Expense
	Debtors []expenses.Debt
}

type UserDebtSummary struct {
	Debtor expenses.User
	Total  float32
}

func overviewContent() template.HTML {
	cfg := config.GetInstance()

	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "htmx/overview.html"),
		map[string]any{},
	)

	return content
}

func overviewPartialPage(c *gin.Context) {
	content := overviewContent()
	c.String(http.StatusOK, string(content))
}

func overviewPage(c *gin.Context) {
	content := overviewContent()

	c.HTML(http.StatusOK, "terminal.html", gin.H{
		"page":         "overview",
		"renderNavBar": true,
		"content":      content,
	})
}

func getResults(c *gin.Context) {
	cfg := config.GetInstance()

	startDateStr := c.PostForm("range-start") + " 00:00:00"
	endDateStr := c.PostForm("range-end") + " 23:59:59"

	startDateTime, err := time.ParseInLocation("02-Jan-2006 15:04:05", startDateStr, time.UTC)
	if err != nil {
		log.Printf("error startDate: %v", err)
		return
	}
	startDate := startDateTime.Unix()

	endDateTime, err := time.ParseInLocation("02-Jan-2006 15:04:05", endDateStr, time.UTC)
	if err != nil {
		log.Printf("error endDate: %v", err)
		return
	}
	endDate := endDateTime.Unix()

	queriedExpenses, err := expenses.GetExpensesRange(startDate, endDate)
	if err != nil {
		log.Printf("getting expenses: %v", err)
		return
	}

	// Get the expenses and shares for each expense
	for k := range queriedExpenses {
		err = queriedExpenses[k].GetShares()
		if err != nil {
			log.Printf("failed to get shares: %v", err)
			return
		}
		err = queriedExpenses[k].GetPayments()
		if err != nil {
			log.Printf("failed to get payments: %v", err)
			return
		}
	}

	userDebtSummary := make(map[expenses.User]float32)
	expensesWithDebts := []ExpenseDebtOverview{}
	for _, exp := range queriedExpenses {
		if len(exp.Shares) <= 1 {
			continue
		}

		expenseDebts, err := exp.CalculateDebts()
		if err != nil {
			// whatevs yo
			continue
		}
		if len(expenseDebts) <= 0 {
			continue
		}

		expensesWithDebts = append(expensesWithDebts, ExpenseDebtOverview{
			Exp: &exp, Debtors: expenseDebts,
		})

		for _, debt := range expenseDebts {
			userDebtSummary[debt.Debtor] += debt.Sum
		}
	}

	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "htmx/overviewResults.html"),
		map[string]any{
			"expenses": expensesWithDebts,
			"summary":  userDebtSummary,
		},
	)

	c.String(http.StatusOK, string(content))
}

func RouteOverview(router *gin.Engine) {
	router.GET("/overview", overviewPage)
	router.POST("/overview", overviewPartialPage)

	router.POST("/overview/ranged", getResults)
}
