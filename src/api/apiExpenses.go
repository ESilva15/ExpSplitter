package api

import (
	"context"
	exp "expenses/expenses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllExpenses(c *gin.Context) {
	ctx := context.Background()
	// TODO - only get the expenses of the user
	// create roles tho, so an admin or someone can get way more expenses
	expenses, err := exp.App.GetAllExpenses(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
}

func RouteExpenses(router *gin.RouterGroup) {
	router.GET("expenses", getAllExpenses)
}
