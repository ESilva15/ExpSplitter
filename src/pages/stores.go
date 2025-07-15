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
	StoresPath = "/stores"
)

func storesContent() (map[string]any, error) {
	stores, err := expenses.GetAllStores()
	if err != nil {
		return nil, err
	}

	data := map[string]any{
		"size":     len(stores) - 1,
		"stores":   stores,
		"resource": "stores",
	}

	return data, nil
}

func storesGlobalPage(c *gin.Context) {
	cfg := config.GetInstance()

	stores, err := storesContent()
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "htmx/stores.html"),
		stores,
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "stores",
		"renderNavBar": true,
		"content":      content,
	})
}

func storePage(c *gin.Context) {
	cfg := config.GetInstance()

	storeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}

	store, err := expenses.GetCategory(storeID)
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/category.html"),
		map[string]any{
			"store": store,
		},
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "store",
		"renderNavBar": false,
		"content":      content,
	})
}

func newStorePage(c *gin.Context) {
	cfg := config.GetInstance()

	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/storeNew.html"), map[string]any{},
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "storeNew",
		"renderNavBar": false,
		"content":      content,
	})
}

func createStore(c *gin.Context) {
	newStoreName := c.PostForm("store-name")

	newStore := expenses.Store{
		StoreName: newStoreName,
	}

	err := newStore.Insert()
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func deleteStore(c *gin.Context) {
	storeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}

	store := expenses.Store{
		StoreID: storeID,
	}
	err = store.Delete()
	if err == expenses.ErrNotFound {
		errMsg := fmt.Sprintf("category %d not found", storeID)
		c.String(http.StatusNotFound, errMsg)
		return
	}

	if err != nil {
		errMsg := fmt.Sprintf("error deleting category %d", storeID)
		c.String(http.StatusInternalServerError, errMsg)
		return
	}

	c.Status(http.StatusNoContent)
}

func RouteStores(router *gin.Engine) {
	router.GET(StoresPath, storesGlobalPage)
	router.GET(StoresPath+"/:id", storePage)
	router.POST(StoresPath, storesGlobalPage)

	router.GET(StoresPath+"/new", newStorePage)
	router.POST(StoresPath+"/new", createStore)
	router.DELETE(StoresPath+"/:id", deleteStore)
}
