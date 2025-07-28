package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundView(c *gin.Context, msg string) {
	c.HTML(404, "terminal", gin.H{
		"warning":      true,
		"renderNavBar": false,
		"content":      "panic",
		"msg":          msg,
	})
}

func RouteNotFound(router *gin.Engine) {
	// 404 Page
	router.NoRoute(func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			c.Header("HX-Redirect", "/404")
			return
		}

		NotFoundView(c, "Not found")
	})
}
