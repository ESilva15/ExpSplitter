package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouteHome(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, ExpensesPath)
	})
}
