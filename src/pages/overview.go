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

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expenses",
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

	// Now we are going to try and get how much each user paid and how much 
	// it owes
	summary := make(map[int]float32)
	expensesWithDebts := []expenses.Expense{}
	for _, exp := range queriedExpenses {
		if len(exp.Shares) <= 1 {
			continue
		}

		log.Printf("ID: %d, Value: %.2f", exp.ExpID, exp.Value)
		if len(exp.Shares) > 1 {
			userShares := userShares(&exp)
			userPayments := userPayments(&exp)

			for userID, payment := range userPayments {
				owed := (userShares[userID] * exp.Value) - userPayments[userID]
				log.Printf("  [%d] %.2f - %.2f", userID, payment, owed)
				if owed > 0 {
					summary[userID] += owed
				}
			}
		}

		expensesWithDebts = append(expensesWithDebts, exp)
	}

	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "htmx/overviewResults.html"),
		map[string]any{
			"expenses": expensesWithDebts,
			"summary": summary,
		},
	)

	log.Println("Summary:")
	log.Println(summary)
	c.String(http.StatusOK, string(content))
}

func userPayments(e *expenses.Expense) map[int]float32 {
	userPayments := make(map[int]float32)

	for _, payment := range e.Payments {
		userPayments[payment.User.UserID] += payment.PayedAmount
	}

	return userPayments
}

func userShares(e *expenses.Expense) map[int]float32 {
	userShares := make(map[int]float32)

	for _, share := range e.Shares {
		userShares[share.User.UserID] = share.Share
	}

	return userShares
}

func RouteOverview(router *gin.Engine) {
	router.GET("/overview", overviewPage)
	router.POST("/overview", overviewPartialPage)

	router.POST("/overview/ranged", getResults)
}
