package pages

import (
	"github.com/gin-gonic/gin"
)

func ServerErrorView(c *gin.Context, msg string) {
	c.HTML(500, "terminal", gin.H{
		"warning":      true,
		"renderNavBar": false,
		"content":      "panic",
		"msg":          msg,
	})
}

func RouteServerError(router *gin.Engine) {
	// 500 Page
	router.GET("/500", func(c *gin.Context) {
		ServerErrorView(c, "The server too makes mistakes")
	})
}
