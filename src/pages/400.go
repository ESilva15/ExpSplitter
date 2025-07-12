package pages

import (
	"expenses/config"
	"expenses/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundView(c *gin.Context, msg string) {
	cfg := config.GetInstance()

	errHtml := utils.HtmlTemplate(
		cfg.AssetsDir+"htmx/panic.tpl", map[string]any{
			"msg": msg,
		},
	)

	c.HTML(404, "terminal.gotempl", gin.H{
		"warning":      true,
		"renderNavBar": false,
		"content":      errHtml,
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
