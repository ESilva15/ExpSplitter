package pages

import (
	"encoding/json"
	exp "expenses/expenses"
	experr "expenses/expenses/errors"
	mod "expenses/expenses/models"
	"fmt"
	"log"
	"net/http"

	fatqr "github.com/ESilva15/gofatqr"
	"github.com/gin-gonic/gin"
)

const (
	ExpensesPath = "/expenses"
)

func expensesPartial(c *gin.Context) {
	ctx, err := getLoggedInUserCTX(c)
	if err != nil {
		ServerErrorView(c, "Could not get logged in user")
		return
	}

	expenses, err := exp.App.GetAllExpenses(ctx)
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}

	c.HTML(http.StatusOK, "expenses", gin.H{
		"expenses": expenses,
	})
}

func ExpensesGlobalPage(c *gin.Context) {
	ctx, err := getLoggedInUserCTX(c)
	if err != nil {
		ServerErrorView(c, "Could not get logged in user")
		return
	}

	expenses, err := exp.App.GetAllExpenses(ctx)
	if err != nil {
		log.Panicln("error:", err)
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
	ctx, err := getLoggedInUserCTX(c)
	if err != nil {
		ServerErrorView(c, "Could not get logged in user")
		return
	}

	expenseID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		ServerErrorView(c, "failed to parse requested expenses id")
		return
	}

	expense, err := exp.App.GetExpense(expenseID)
	if err != nil {
		NotFoundView(c, "didn't find requested expense")
		return
	}

	err = exp.App.LoadExpenseShares(&expense)
	if err != nil {
		log.Println("error:", err)
		ServerErrorView(c, "failed to fetch shares")
		return
	}

	err = exp.App.LoadExpensePayments(&expense)
	if err != nil {
		log.Println("error:", err)
		ServerErrorView(c, "failed to fetch payments")
		return
	}

	err = exp.App.LoadExpenseDebts(&expense)
	if err != nil {
		ServerErrorView(c, "failed to fetch debts")
		return
	}

	categories, err := exp.App.GetAllCategories(ctx)
	if err != nil {
		ServerErrorView(c, "failed to fetch categories")
		return
	}

	stores, err := exp.App.GetAllStores(ctx)
	if err != nil {
		ServerErrorView(c, "failed to fetch stores")
		return
	}

	types, err := exp.App.GetAllTypes(ctx)
	if err != nil {
		ServerErrorView(c, "failed to fetch types")
		return
	}

	users, err := exp.App.GetAllUsers(ctx)
	if err != nil {
		ServerErrorView(c, "failed to fetch users")
		return
	}

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
	ctx, err := getLoggedInUserCTX(c)
	if err != nil {
		log.Println("failed to fetch logged in user -", err.Error())
		ServerErrorView(c, "The server too makes mistakes")
		return
	}

	categories, err := exp.App.GetAllCategories(ctx)
	if err != nil {
		ServerErrorView(c, "failed to fetch categories")
		return
	}

	stores, err := exp.App.GetAllStores(ctx)
	if err != nil {
		ServerErrorView(c, "failed to fetch stores")
		return
	}

	types, err := exp.App.GetAllTypes(ctx)
	if err != nil {
		ServerErrorView(c, "failed to fetch types")
		return
	}

	users, err := exp.App.GetAllUsers(ctx)
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
	ctx, err := getLoggedInUserCTX(c)
	if err != nil {
		log.Println("failed to fetch logged in user -", err.Error())
		ServerErrorView(c, "The server too makes mistakes")
		return
	}

	newExp, err := expenseFromForm(c, ctx)
	if err != nil {
		// TODO
		// Change this to something the user can see
		c.Header("HX-Redirect", "/500")
		return
	}

	err = exp.App.NewExpense(ctx, *newExp)
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func updateExpense(c *gin.Context) {
	ctx, err := getLoggedInUserCTX(c)
	if err != nil {
		log.Println("failed to fetch logged in user -", err.Error())
		ServerErrorView(c, "The server too makes mistakes")
		return
	}

	expenseID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		c.Header("HX-Redirect", "/404")
		return
	}

	newExp, err := expenseFromForm(c, ctx)
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
		log.Println("error:", err)
		errMsg := fmt.Sprintf("category %d not found", expenseID)
		c.String(http.StatusNotFound, errMsg)
		return
	}

	if err != nil {
		log.Println("error:", err)
		errMsg := fmt.Sprintf("error deleting category %d", expenseID)
		c.String(http.StatusInternalServerError, errMsg)
		return
	}

	c.Status(http.StatusNoContent)
}

func qrRequest(c *gin.Context) {
	qr := c.PostForm("expense-qr")
	log.Println(qr)

	var fat fatqr.FatQR
	err := fat.Scan(qr, 0)
	if err != nil {
		log.Println("error scanning QR code:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// try to find storeID
	storeID, err := exp.App.GetStoreIDFromNIF(fat.TaxRegistrationNumber)
	if err != nil {
		log.Println("error getting storeID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := map[string]string{
		"total":   fat.GrossTotal.String(),
		"date":    fat.InvoiceDate.Format("02-Jan-2006"),
		"storeID": fmt.Sprintf("%d", storeID),
	}

	c.JSON(200, response)
}

func RouteExpenses(router *gin.RouterGroup) {
	router.GET(ExpensesPath, ExpensesGlobalPage)
	router.GET(ExpensesPath+"/:id", expensePage)
	router.POST(ExpensesPath, expensesPartial)

	router.GET(ExpensesPath+"/new", newExpensePage)
	router.POST(ExpensesPath+"/new", createExpense)
	router.PUT(ExpensesPath+"/:id", updateExpense)
	router.DELETE(ExpensesPath+"/:id", deleteExpense)

	router.POST(ExpensesPath+"/scanQR", qrRequest)
}
