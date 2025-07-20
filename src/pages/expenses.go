package pages

import (
	"encoding/json"
	"expenses/config"
	"expenses/expenses"
	"expenses/templating"
	"fmt"
	"log"
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
	err = expense.GetPayments()
	if err != nil {
		ServerErrorView(c, "failed to fetch payments")
		return
	}

	err = expense.GetShares()
	if err != nil {
		ServerErrorView(c, "failed to fetch payments")
		return
	}

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

	users, err := expenses.GetAllUsers()
	if err != nil {
		ServerErrorView(c, "failed to fetch users")
		return
	}

	summary, err := expense.GetSummary()
	if err != nil {
		ServerErrorView(c, "failed to get summary")
		return
	}

	log.Println("Users:")
	log.Println(users)

	log.Println("Payments:")
	log.Println(expense.Payments)

	log.Println("Shares:")
	log.Println(expense.Shares)

	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/expense.html"),
		map[string]any{
			"method":     "put",
			"expense":    expense,
			"categories": categories,
			"stores":     stores,
			"types":      types,
			"summary":    summary,
			"users":      users,
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

	users, err := expenses.GetAllUsers()
	if err != nil {
		ServerErrorView(c, "failed to fetch users")
		return
	}

	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/expense.html"), map[string]any{
			"method":     "post",
			"expense":    expenses.NewExpense(),
			"categories": categories,
			"stores":     stores,
			"types":      types,
			"users":      users,
		},
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expenseNew",
		"renderNavBar": false,
		"content":      content,
	})
}

func createExpense(c *gin.Context) {
	newExp, err := expenseFromForm(c)
	if err != nil {
		// TODO
		// Change this to something the user can see
		c.Header("HX-Redirect", "/500")
	}

	err = newExp.Insert()
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func updateExpense(c *gin.Context) {
	expenseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Header("HX-Redirect", "/404")
		return
	}

	newExp, err := expenseFromForm(c)
	if err != nil {
		// TODO
		// Change this to something the user can see
		c.Header("HX-Redirect", "/500")
		return
	}
	newExp.ExpID = expenseID

	err = newExp.Update()
	if err != nil {
		errMsg, _ := json.Marshal(err.Error())
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":%s}", errMsg))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func RouteExpenses(router *gin.Engine) {
	router.GET(ExpensesPath, ExpensesGlobalPage)
	router.GET(ExpensesPath+"/:id", expensePage)
	router.POST(ExpensesPath, expensesPartial)

	router.GET(ExpensesPath+"/new", newExpensePage)
	router.POST(ExpensesPath+"/new", createExpense)
	router.PUT(ExpensesPath+"/:id", updateExpense)
}
