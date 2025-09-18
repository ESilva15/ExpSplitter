package pages

import (
	"log"

	exp "github.com/ESilva15/expenses/expenses"
	mod "github.com/ESilva15/expenses/expenses/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type ExpenseDebtOverview struct {
	Exp     *mod.Expense
	Debtors mod.Debts
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
	ctx, err := getLoggedInUserCTX(c)
	if err != nil {
		ServerErrorView(c, "Could not get logged in user")
		return
	}

	filter, err := expenseFilterFromQuery(c)
	if err != nil {
		log.Println("error ")
	}

	queriedExpenses, err := exp.App.GetAllExpenses(ctx, filter)
	if err != nil {
		log.Printf("getting expenses: %v", err)
		return
	}

	// Get the expenses and shares for each expense
	for k := range queriedExpenses {
		err = exp.App.LoadExpenseShares(&queriedExpenses[k])
		if err != nil {
			log.Printf("failed to get shares: %v", err)
			return
		}

		err = exp.App.LoadExpensePayments(&queriedExpenses[k])
		if err != nil {
			log.Printf("failed to get payments: %v", err)
			return
		}

		err = exp.App.LoadExpenseDebts(&queriedExpenses[k])
		if err != nil {
			log.Printf("failed to get debts: %v", err)
			return
		}
	}

	userDebtSummary := make(map[mod.User]decimal.Decimal)
	expensesWithDebts := []*mod.Expense{}
	for _, e := range queriedExpenses {
		if len(e.Shares) <= 1 || len(e.Debts) <= 0 {
			continue
		}

		expensesWithDebts = append(expensesWithDebts, &e)
		for _, debt := range e.Debts {
			userDebtSummary[debt.Debtor] = userDebtSummary[debt.Debtor].Add(debt.Sum)
		}
	}

	c.HTML(http.StatusOK, "overviewResults", gin.H{
		"expenses": expensesWithDebts,
		"summary":  userDebtSummary,
	})
}

func RouteOverview(router *gin.RouterGroup) {
	router.GET("/overview", overviewPage)
	router.POST("/overview", overviewPartialPage)

	router.GET("/overview/ranged", getResults)
}
