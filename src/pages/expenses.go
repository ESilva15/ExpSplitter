package pages

import (
	"encoding/json"
	"expenses/expenses"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ExpensesPath = "/expenses"
)

func expensesPartial(c *gin.Context) {
	expenses, err := expenses.GetAllExpenses()
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}

	c.HTML(http.StatusOK, "expenses", gin.H{
		"expenses": expenses,
	})
}

func ExpensesGlobalPage(c *gin.Context) {
	expenses, err := expenses.GetAllExpenses()
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}

	c.HTML(http.StatusOK, "terminal.html", gin.H{
		"page":         "expenses",
		"renderNavBar": true,
		"content":      "expenses",
		"expenses":     expenses,
	})
}

func expensePage(c *gin.Context) {
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

	if expense.ExpID == -1 {
		NotFoundView(c, fmt.Sprintf("ID `%d` doesn't exist", expenseID))
		return
	}

	c.HTML(http.StatusOK, "terminal.html", gin.H{
		"page":         "expense",
		"renderNavBar": false,
		"content":      "expense",
		"method":       "put",
		"expense":      expense,
		"categories":   categories,
		"stores":       stores,
		"types":        types,
		"summary":      summary,
		"users":        users,
	})
}

func newExpensePage(c *gin.Context) {
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

	c.HTML(http.StatusOK, "terminal.html", gin.H{
		"page":         "expenseNew",
		"renderNavBar": false,
		"content":      "newExpense",
		"method":     "post",
		"expense":    expenses.NewExpense(),
		"categories": categories,
		"stores":     stores,
		"types":      types,
		"users":      users,
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

func deleteExpense(c *gin.Context) {
	expenseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Header("HX-Redirect", "/404")
		return
	}

	payment := expenses.Expense{
		ExpID: expenseID,
	}

	err = payment.Delete()
	if err == expenses.ErrNotFound {
		errMsg := fmt.Sprintf("category %d not found", expenseID)
		c.String(http.StatusNotFound, errMsg)
		return
	}

	if err != nil {
		errMsg := fmt.Sprintf("error deleting category %d", expenseID)
		c.String(http.StatusInternalServerError, errMsg)
		return
	}

	c.Status(http.StatusNoContent)
}

func RouteExpenses(router *gin.Engine) {
	router.GET(ExpensesPath, ExpensesGlobalPage)
	router.GET(ExpensesPath+"/:id", expensePage)
	router.POST(ExpensesPath, expensesPartial)

	router.GET(ExpensesPath+"/new", newExpensePage)
	router.POST(ExpensesPath+"/new", createExpense)
	router.PUT(ExpensesPath+"/:id", updateExpense)
	router.DELETE(ExpensesPath+"/:id", deleteExpense)
}
