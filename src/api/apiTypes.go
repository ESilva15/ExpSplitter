package api

import (
	"context"
	exp "expenses/expenses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllTypes(c *gin.Context) {
	// TODO - we need to set the current user from the JWT token we got
	ctx := context.Background()

	// TODO - only get the types of the user
	// create roles tho, so an admin or someone can get way more types
	types, err := exp.App.GetAllTypes(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"users": types})
}

// RouteTypes will route the endpoints for the types resources
func RouteTypes(router *gin.RouterGroup) {
	router.GET("types", getAllTypes)
}
