package pages

import (
	exp "expenses/expenses"
	mod "expenses/expenses/models"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type ExpenseDebtOverview struct {
	Exp     *mod.Expense
	Debtors []exp.Debt
}

type UserDebtSummary struct {
	Debtor mod.User
	Total  float64
}

func overviewPartialPage(c *gin.Context) {
	c.HTML(http.StatusOK, "overview", gin.H{})
}

func overviewPage(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "overview",
		"renderNavBar": true,
		"content":      "overview",
	})
}

func getResults(c *gin.Context) {
	startDate := c.PostForm("range-start") + " 00:00:00"
	endDate := c.PostForm("range-end") + " 23:59:59"

	queriedExpenses, err := exp.Serv.GetExpensesRanged(startDate, endDate)
	if err != nil {
		log.Printf("getting expenses: %v", err)
		return
	}

	// Get the expenses and shares for each expense
	for k := range queriedExpenses {
		err = exp.Serv.LoadExpenseShares(&queriedExpenses[k])
		if err != nil {
			log.Printf("failed to get shares: %v", err)
			return
		}
		err = exp.Serv.LoadExpensePayments(&queriedExpenses[k])
		if err != nil {
			log.Printf("failed to get payments: %v", err)
			return
		}
	}

	userDebtSummary := make(map[mod.User]decimal.Decimal)
	expensesWithDebts := []ExpenseDebtOverview{}
	for _, e := range queriedExpenses {
		if len(e.Shares) <= 1 {
			continue
		}

		expenseDebts, err := exp.CalculateDebts(&e)
		if err != nil {
			// whatevs yo
			continue
		}
		if len(expenseDebts) <= 0 {
			continue
		}

		expensesWithDebts = append(expensesWithDebts, ExpenseDebtOverview{
			Exp: &e, Debtors: expenseDebts,
		})

		for _, debt := range expenseDebts {
			userDebtSummary[debt.Debtor] = userDebtSummary[debt.Debtor].Add(debt.Sum)
		}
	}

	c.HTML(http.StatusOK, "overviewResults", gin.H{
		"expenses": expensesWithDebts,
		"summary":  userDebtSummary,
	})
}

func RouteOverview(router *gin.Engine) {
	router.GET("/overview", overviewPage)
	router.POST("/overview", overviewPartialPage)

	router.POST("/overview/ranged", getResults)
}
