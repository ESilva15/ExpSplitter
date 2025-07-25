package main

import (
	"expenses/config"
	"expenses/pages"

	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// This can be set in the compilation command - in the Makefile
var ginMode = "debug"

func main() {
	gin.SetMode(ginMode)

	config.SetConfig("./config.yaml")

	cfg := config.GetInstance()
	router := gin.Default()

	router.StaticFS("/assets", http.Dir(cfg.AssetsDir))
	router.StaticFile("/favicon.ico", filepath.Join(cfg.AssetsDir, "img/favicon.webp"))
	router.LoadHTMLGlob(filepath.Join(cfg.AssetsDir, "htmx/*.gotempl"))

	pages.RouteHome(router)
	pages.RouteExpenses(router)
	pages.RouteCategories(router)
	pages.RouteStores(router)
	pages.RouteTypes(router)
	pages.RoutePayments(router)
	pages.RouteOverview(router)

	pages.RouteNotFound(router)
	pages.RouteServerError(router)

	err := router.Run(":8081")
	if err != nil {
		log.Printf("Failed to launch application: \n  %s\n", err)
	}
}
