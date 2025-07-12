package pages

import (
	"expenses/config"
	"expenses/utils"

	"github.com/gin-gonic/gin"
)

func ServerErrorView(c *gin.Context, msg string) {
	cfg := config.GetInstance()

	errHtml := utils.HtmlTemplate(
		cfg.AssetsDir+"htmx/panic.tpl", map[string]any{
			"msg": msg,
		},
	)

	c.HTML(500, "terminal.gotempl", gin.H{
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
