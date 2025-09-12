package api

import (
	"context"
	exp "expenses/expenses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllCategories(c *gin.Context) {
	// TODO
	// What user is logged in?
	ctx := context.Background()

	// TODO - only get the categories of the user
	// create roles tho, so an admin or someone can get way more categories
	categories, err := exp.App.GetAllCategories(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func RouteCategories(router *gin.RouterGroup) {
	router.GET("categories", getAllCategories)
}
