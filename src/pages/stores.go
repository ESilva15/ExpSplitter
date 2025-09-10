package pages

import (
	"database/sql"
	exp "expenses/expenses"
	experr "expenses/expenses/errors"
	"fmt"
	"log"
	"net/http"

	fatqr "github.com/ESilva15/gofatqr"
	"github.com/gin-gonic/gin"
)

const (
	StoresPath = "/stores"
)

func storesGlobalPage(c *gin.Context) {
	stores, err := exp.App.GetAllStores()
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
	storeID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		NotFoundView(c, "didn't find request store")
		return
	}

	store, err := exp.App.GetStore(storeID)
	if err == sql.ErrNoRows {
		NotFoundView(c, fmt.Sprintf("didn't find store with ID `%d`", storeID))
		return
	}
	if err != nil {
		ServerErrorView(c, fmt.Sprintf("error getting store: %+v", err.Error()))
		return
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "store",
		"renderNavBar": false,
		"content":      "store",
		"store":        store,
		"method":       "put",
	})
}

func newStorePage(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "storeNew",
		"renderNavBar": false,
		"content":      "store",
		"method":       "post",
	})
}

func createStore(c *gin.Context) {
	storeName := c.PostForm("store-name")
	storeNIF := c.PostForm("store-nif")

	err := exp.App.NewStore(storeName, storeNIF)
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func updateStore(c *gin.Context) {
	storeID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}
	newName := c.PostForm("store-name")
	newNIF := c.PostForm("store-nif")

	err = exp.App.UpdateStore(storeID, newName, newNIF)
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func deleteStore(c *gin.Context) {
	storeID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}

	err = exp.App.DeleteStore(storeID)
	if err == experr.ErrNotFound {
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

func getNIF(c *gin.Context) {
	qr := c.PostForm("store-qr")
	log.Println(qr)

	var fat fatqr.FatQR
	err := fat.Scan(qr, 0)
	if err != nil {
		// TODO
		// handle the error here for the user
		return
	}

	response := map[string]string{
		"nif": fat.TaxRegistrationNumber,
	}

	c.JSON(200, response)
}

func RouteStores(router *gin.RouterGroup) {
	router.GET(StoresPath, storesGlobalPage)
	router.GET(StoresPath+"/:id", storePage)
	router.POST(StoresPath, storesGlobalPage)

	router.GET(StoresPath+"/new", newStorePage)
	router.POST(StoresPath+"/new", createStore)
	router.PUT(StoresPath+"/:id", updateStore)
	router.DELETE(StoresPath+"/:id", deleteStore)

	router.POST(StoresPath+"/getNIF", getNIF)
}
