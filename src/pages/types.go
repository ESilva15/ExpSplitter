package pages

import (
	"database/sql"
	exp "expenses/expenses"
	experr "expenses/expenses/errors"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	TypesPath = "/types"
)

func typesGlobalPage(c *gin.Context) {
	types, err := exp.App.GetAllTypes()
	if err != nil {
		ServerErrorView(c, "the server failed to get the types")
		return
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "types",
		"renderNavBar": true,
		"content":      "types",
		"types":        types,
	})
}

func typePage(c *gin.Context) {
	typeID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		NotFoundView(c, "requested type doesn't exist")
		return
	}

	typ, err := exp.App.GetType(typeID)
	if err == sql.ErrNoRows {
		NotFoundView(c, fmt.Sprintf("Couldn't find type `%d`", typeID))
		return
	}
	if err != nil {
		ServerErrorView(c, "failed to retrieve requested type")
		return
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "type",
		"renderNavBar": false,
		"content":      "type",
		"type":         typ,
		"method":       "put",
	})
}

func newTypePage(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "typeNew",
		"renderNavBar": false,
		"content":      "type",
		"method":       "post",
	})
}

func createType(c *gin.Context) {
	newTypName := c.PostForm("type-name")

	err := exp.App.NewType(newTypName)
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func deleteType(c *gin.Context) {
	typeID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}

	err = exp.App.DeleteType(typeID)
	if err == experr.ErrNotFound {
		errMsg := fmt.Sprintf("category %d not found", typeID)
		c.String(http.StatusNotFound, errMsg)
		return
	}

	if err != nil {
		errMsg := fmt.Sprintf("error deleting category %d", typeID)
		c.String(http.StatusInternalServerError, errMsg)
		return
	}

	c.Status(http.StatusNoContent)
}

func updateType(c *gin.Context) {
	typeID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}
	newName := c.PostForm("type-name")

	err = exp.App.UpdateType(typeID, newName)
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func RouteTypes(router *gin.RouterGroup) {
	router.GET(TypesPath, typesGlobalPage)
	router.GET(TypesPath+"/:id", typePage)
	router.POST(TypesPath, typesGlobalPage)

	router.GET(TypesPath+"/new", newTypePage)
	router.POST(TypesPath+"/new", createType)
	router.DELETE(TypesPath+"/:id", deleteType)
	router.PUT(TypesPath+"/:id", updateType)
}
