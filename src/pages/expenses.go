package pages

import (
	"expenses/config"
	"expenses/expenses"
	"expenses/templating"
	"fmt"
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

func expensesPartial(c *gin.Context) {
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

	c.String(http.StatusOK, string(content))
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

func expensePage(c *gin.Context) {
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

	if expense.ExpID == -1 {
		NotFoundView(c, fmt.Sprintf("ID `%d` doesn't exist", expenseID))
		return
	}

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expense",
		"renderNavBar": false,
		"content":      content,
	})
}

func newExpensePage(c *gin.Context) {
	cfg := config.GetInstance()

	categories, err := expenses.GetAllCategories()
	if err != nil {
		ServerErrorView(c, "failed to fetch categories")
		return
	}

	stores, err := expenses.GetAllStores()
	if err != nil {
		ServerErrorView(c, "failed to fetch stores")
		return
	}

	types, err := expenses.GetAllTypes()
	if err != nil {
		ServerErrorView(c, "failed to fetch types")
		return
	}

	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/expenseNew.html"), map[string]any{
			"expense":    expenses.NewExpense(),
			"categories": categories,
			"stores":     stores,
			"types":      types,
		},
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expenseNew",
		"renderNavBar": false,
		"content":      content,
	})
}

func createExpense(c *gin.Context) {
	// newCatName := c.PostForm("category-name")

	// newExp := expenses.Expense{
	// 	CategoryName: newCatName,
	// }
	//
	// err := newExp.Insert()
	// if err != nil {
	// 	c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
	// }

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func RouteExpenses(router *gin.Engine) {
	router.GET(ExpensesPath, ExpensesGlobalPage)
	router.GET(ExpensesPath+"/:id", expensePage)
	router.POST(ExpensesPath, expensesPartial)

	router.GET(ExpensesPath+"/new", newExpensePage)
	router.POST(ExpensesPath+"/new", createExpense)
}
