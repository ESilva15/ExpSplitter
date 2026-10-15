package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func analysisPage(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "analysis",
		"renderNavBar": true,
		"content":      "analysis",
	})
}

func analysisPartialPage(c *gin.Context) {
	c.HTML(http.StatusOK, "analysis", gin.H{})
}

// RouteAnalysis routes the endpoints for analysis stuff.
func RouteAnalysis(router *gin.RouterGroup) {
	router.GET("/analysis", analysisPage)
	router.POST("/analysis", analysisPartialPage)
}
