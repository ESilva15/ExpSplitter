package pages

import (
	"expenses/config"
	"expenses/expenses"
	"expenses/templating"

	"fmt"
	"net/http"
	fp "path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	TypesPath = "/types"
)

func typesContent() (map[string]any, error) {
	types, err := expenses.GetAllTypes()
	if err != nil {
		return nil, err
	}

	data := map[string]any{
		"size":     len(types) - 1,
		"types":    types,
		"resource": "types",
	}

	return data, nil
}

func typesGlobalPage(c *gin.Context) {
	cfg := config.GetInstance()

	types, err := typesContent()
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "htmx/types.html"),
		types,
	)

	c.HTML(http.StatusOK, "terminal.html", gin.H{
		"page":         "types",
		"renderNavBar": true,
		"content":      content,
	})
}

func typePage(c *gin.Context) {
	cfg := config.GetInstance()

	typeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}

	typ, err := expenses.GetType(typeID)
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/type.html"),
		map[string]any{
			"type": typ,
		},
	)

	c.HTML(http.StatusOK, "terminal.html", gin.H{
		"page":         "type",
		"renderNavBar": false,
		"content":      content,
	})
}

func newTypePage(c *gin.Context) {
	cfg := config.GetInstance()

	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/typeNew.html"), map[string]any{},
	)

	c.HTML(http.StatusOK, "terminal.html", gin.H{
		"page":         "typeNew",
		"renderNavBar": false,
		"content":      content,
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
