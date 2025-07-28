package pages

import (
	"expenses/config"
	"expenses/templating"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundView(c *gin.Context, msg string) {
	cfg := config.GetInstance()

	errHtml := templating.HtmlTemplate(
		cfg.AssetsDir+"htmx/panic.html", map[string]any{
			"msg": msg,
		},
	)

	c.HTML(404, "terminal.html", gin.H{
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
