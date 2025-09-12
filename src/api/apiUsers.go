package api

import (
	"context"
	exp "expenses/expenses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllUsers(c *gin.Context) {
	// TODO - stuff
	ctx := context.Background()

	// TODO - only get the users of the user
	// create roles tho, so an admin or someone can get way more users
	users, err := exp.App.GetAllUsers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func RouteUsers(router *gin.RouterGroup) {
	router.GET("users", getAllUsers)
}
