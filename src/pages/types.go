package pages

import (
	"expenses/expenses"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	TypesPath = "/types"
)

func typesGlobalPage(c *gin.Context) {
	types, err := expenses.GetAllTypes()
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
	typeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NotFoundView(c, "requested type doesn't exist")
		return
	}

	typ, err := expenses.GetType(typeID)
	if err != nil {
		ServerErrorView(c, "failed to retrieve requested type")
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "type",
		"renderNavBar": false,
		"content":      "type",
		"type":         typ,
	})
}

func newTypePage(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "typeNew",
		"renderNavBar": false,
		"content":      "typeNew",
	})
}

func createType(c *gin.Context) {
	newTypName := c.PostForm("type-name")

	newTyp := expenses.Type{
		TypeName: newTypName,
	}

	err := newTyp.Insert()
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func deleteType(c *gin.Context) {
	typeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}

	typ := expenses.Type{
		TypeID: typeID,
	}
	err = typ.Delete()
	if err == expenses.ErrNotFound {
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
	typeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}
	newName := c.PostForm("type-name")

	newTyp := expenses.Type{
		TypeID:   typeID,
		TypeName: newName,
	}

	err = newTyp.Update()
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func RouteTypes(router *gin.Engine) {
	router.GET(TypesPath, typesGlobalPage)
	router.GET(TypesPath+"/:id", typePage)
	router.POST(TypesPath, typesGlobalPage)

	router.GET(TypesPath+"/new", newTypePage)
	router.POST(TypesPath+"/new", createType)
	router.DELETE(TypesPath+"/:id", deleteType)
	router.PUT(TypesPath+"/:id", updateType)
}
