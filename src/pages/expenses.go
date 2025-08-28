package pages

import (
	"database/sql"
	"encoding/json"
	exp "expenses/expenses"
	experr "expenses/expenses/errors"
	mod "expenses/expenses/models"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ExpensesPath = "/expenses"
)

func expensesPartial(c *gin.Context) {
	expenses, err := exp.App.GetAllExpenses()
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}

	c.HTML(http.StatusOK, "expenses", gin.H{
		"expenses": expenses,
	})
}

func ExpensesGlobalPage(c *gin.Context) {
	expenses, err := exp.App.GetAllExpenses()
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "expenses",
		"renderNavBar": true,
		"content":      "expenses",
		"expenses":     expenses,
	})
}

func expensePage(c *gin.Context) {
	expenseID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		ServerErrorView(c, "failed to parse requested expenses id")
		return
	}

	expense, err := exp.App.GetExpense(expenseID)
	if err == sql.ErrNoRows {
		NotFoundView(c, "didn't find requested expense")
		return
	}

	err = exp.App.LoadExpenseShares(&expense)
	if err != nil {
		ServerErrorView(c, "failed to fetch shares")
		return
	}

	err = exp.App.LoadExpensePayments(&expense)
	if err != nil {
		ServerErrorView(c, "failed to fetch payments")
		return
	}

	err = exp.App.LoadExpenseDebts(&expense)
	if err != nil {
		ServerErrorView(c, "failed to fetch debts")
		return
	}

	categories, err := exp.GetAllCategories()
	if err != nil {
		ServerErrorView(c, "failed to fetch categories")
		return
	}

	stores, err := exp.App.GetAllStores()
	if err != nil {
		ServerErrorView(c, "failed to fetch stores")
		return
	}

	types, err := exp.App.GetAllTypes()
	if err != nil {
		ServerErrorView(c, "failed to fetch types")
		return
	}

	users, err := exp.App.GetAllUsers()
	if err != nil {
		ServerErrorView(c, "failed to fetch users")
		return
	}

	// summary, err := expense.GetSummary()
	// if err != nil {
	// 	ServerErrorView(c, "failed to get summary")
	// 	return
	// }

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "expense",
		"renderNavBar": false,
		"content":      "expense",
		"method":       "put",
		"expense":      expense,
		"categories":   categories,
		"stores":       stores,
		"types":        types,
		"users":        users,
		// "summary":      summary,
	})
}

func newExpensePage(c *gin.Context) {
	categories, err := exp.GetAllCategories()
	if err != nil {
		ServerErrorView(c, "failed to fetch categories")
		return
	}

	stores, err := exp.App.GetAllStores()
	if err != nil {
		ServerErrorView(c, "failed to fetch stores")
		return
	}

	types, err := exp.App.GetAllTypes()
	if err != nil {
		ServerErrorView(c, "failed to fetch types")
		return
	}

	users, err := exp.App.GetAllUsers()
	if err != nil {
		ServerErrorView(c, "failed to fetch users")
		return
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "expenseNew",
		"renderNavBar": false,
		"content":      "newExpense",
		"method":       "post",
		"expense":      mod.NewExpense(),
		"categories":   categories,
		"stores":       stores,
		"types":        types,
		"users":        users,
	})
}

func createExpense(c *gin.Context) {
	newExp, err := expenseFromForm(c)
	if err != nil {
		// TODO
		// Change this to something the user can see
		c.Header("HX-Redirect", "/500")
	}

	err = exp.App.NewExpense(*newExp)
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func updateExpense(c *gin.Context) {
	expenseID, err := exp.ParseID(c.Param("id"))
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

	err = exp.App.UpdateExpense(*newExp)
	if err != nil {
		errMsg, _ := json.Marshal(err.Error())
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":%s}", errMsg))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func deleteExpense(c *gin.Context) {
	expenseID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		c.Header("HX-Redirect", "/404")
		return
	}

	err = exp.App.DeleteExpense(expenseID)
	if err == experr.ErrNotFound {
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
