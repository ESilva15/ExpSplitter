package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouteHome routes the home pages.
func RouteHome(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/main")
	})
}
