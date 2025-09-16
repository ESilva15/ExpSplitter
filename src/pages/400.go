package pages

import (
	"net/http"
	"strings"

	"github.com/ESilva15/expenses/api"

	"github.com/gin-gonic/gin"
)

func NotFoundView(c *gin.Context, msg string) {
	if strings.HasPrefix(c.Request.URL.Path, api.APIPath) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "not_found",
			"message": "endpoint doesn't exist",
			"path":    c.Param("path"),
		})
		return
	}

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
