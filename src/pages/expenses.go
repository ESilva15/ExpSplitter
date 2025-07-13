package pages

import (
	"expenses/config"
	"expenses/expenses"
	"expenses/templating"
	"strconv"

	"net/http"
	fp "path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	ExpensesPath = "/expenses"
)

func expensesContent() (map[string]any, error) {
	expenses, err := expenses.GetAllExpenses()
	if err != nil {
		return nil, err
	}

	data := map[string]any{
		"size":     len(expenses) - 1,
		"expenses": expenses,
		"resource": "expenses",
	}

	return data, nil
}

func ExpensesGlobalPage(c *gin.Context) {
	cfg := config.GetInstance()

	expenses, err := expensesContent()
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "htmx/expenses.html"),
		expenses,
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expenses",
		"renderNavBar": true,
		"content":      content,
	})
}

func ExpensePage(c *gin.Context) {
	cfg := config.GetInstance()

	expenseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}

	expense, err := expenses.GetExpense(expenseID)
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/expense.html"),
		map[string]any{
			"expense": expense,
		},
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expense",
		"renderNavBar": false,
		"content":      content,
	})
}

func RouteExpenses(router *gin.Engine) {
	router.GET(ExpensesPath, ExpensesGlobalPage)
	router.POST(ExpensesPath, ExpensesGlobalPage)

	router.GET(ExpensesPath+"/:id", ExpensePage)
}
