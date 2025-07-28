package pages

import (
	"expenses/config"
	"expenses/templating"

	"github.com/gin-gonic/gin"
)

func ServerErrorView(c *gin.Context, msg string) {
	cfg := config.GetInstance()

	errHtml := templating.HtmlTemplate(
		cfg.AssetsDir+"htmx/panic.html", map[string]any{
			"msg": msg,
		},
	)

	c.HTML(500, "terminal.html", gin.H{
		"warning":      true,
		"renderNavBar": false,
		"content":      errHtml,
	})
}

func RouteServerError(router *gin.Engine) {
	// 500 Page
	router.GET("/500", func(c *gin.Context) {
		ServerErrorView(c, "The server too makes mistakes")
	})
}
