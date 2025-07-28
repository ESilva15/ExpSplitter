package pages

import (
	"expenses/expenses"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	StoresPath = "/stores"
)

func storesGlobalPage(c *gin.Context) {
	stores, err := expenses.GetAllStores()
	if err != nil {
		ServerErrorView(c, "Failed to fetch stores content")
		return
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "stores",
		"renderNavBar": true,
		"content":      "stores",
		"stores":       stores,
	})
}

func storePage(c *gin.Context) {
	storeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NotFoundView(c, "didn't find request store")
		return
	}

	store, err := expenses.GetStore(storeID)

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "store",
		"renderNavBar": false,
		"content":      "store",
		"store":        store,
	})
}

func newStorePage(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "storeNew",
		"renderNavBar": false,
		"content":      "storeNew",
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
