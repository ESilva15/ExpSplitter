package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func expensesGlobalPagePartial(c *gin.Context) {
	data, err := fetchExpensesData(c)
	if err != nil {
		ServerErrorView(c, err.Error())
		return
	}

	c.HTML(http.StatusOK, "mainPage", data)
}

func expensesGlobalPage(c *gin.Context) {
	data, err := fetchExpensesData(c)
	if err != nil {
		ServerErrorView(c, err.Error())
		return
	}

	data["page"] = "mainPage"
	data["renderNavBar"] = true
	data["content"] = "mainPage"

	c.HTML(http.StatusOK, "terminal", data)
}

// RouteHome routes the home pages.
func RouteHome(router *gin.RouterGroup) {
	router.GET("/", expensesGlobalPage)
	router.POST("/", expensesGlobalPagePartial)
}
