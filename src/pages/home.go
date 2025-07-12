package pages

import (
	"expenses/config"
	"expenses/utils"
	"html/template"
	"net/http"
	fp "path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	ExpensesPath = "/expenses"
)

func ExpensesContent() template.HTML {
	cfg := config.GetInstance()

	expensesPage := template.HTML(
		utils.LoadFile(fp.Join(cfg.AssetsDir, "htmx/expenses.html")),
	)

	return expensesPage
}

func GetExpensesPage(c *gin.Context) {
	c.String(http.StatusOK, string(ExpensesContent()))
}

func ExpensesPage(c *gin.Context) {
	content := ExpensesContent()

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expenses",
		"renderNavBar": true,
		"content":      content,
	})
}

func RouteHome(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, ExpensesPath)
	})
	router.GET(ExpensesPath, ExpensesPage)
	router.POST(ExpensesPath, ExpensesPage)
}
