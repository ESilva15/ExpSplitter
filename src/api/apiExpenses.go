package api

import (
	"net/http"
	"time"

	exp "github.com/ESilva15/expenses/expenses"
	mod "github.com/ESilva15/expenses/expenses/models"
	gaux "github.com/ESilva15/expenses/ginAux"

	"github.com/gin-gonic/gin"
)

func getAllExpenses(c *gin.Context) {
	ctx, err := gaux.GetLoggedInUserCTX(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	eFilter, err := gaux.ExpenseFilterFromQuery(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	expenses, err := exp.App.GetAllExpenses(ctx, eFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
}

func createExpense(c *gin.Context) {
	ctx, err := gaux.GetLoggedInUserCTX(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var newExp mod.Expense
	err = c.ShouldBindJSON(&newExp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request - " + err.Error()})
		return
	}
	newExp.CreationDate = time.Now()

	err = exp.App.NewExpense(ctx, newExp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request - " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

// RouteExpenses routes the endpoints for the API.
func RouteExpenses(router *gin.RouterGroup) {
	router.GET("expenses", getAllExpenses)
	router.POST("expenses/create", createExpense)
}
