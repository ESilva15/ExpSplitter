package main

import (
	"expenses/config"
	"expenses/pages"
	"expenses/templating"

	"html/template"
	"log"
	"net/http"
	fp "path/filepath"

	"github.com/gin-gonic/gin"
)

// This can be set in the compilation command - in the Makefile
var ginMode = "debug"

func main() {
	config.SetConfig("./config.yaml")
	cfg := config.GetInstance()

	gin.SetMode(ginMode)
	templates := template.Must(
		template.New("").
		Funcs(templating.TmplFuncMap).
		ParseGlob(fp.Join(cfg.AssetsDir, "htmx/*.html")),
	)

	router := gin.Default()
	router.SetHTMLTemplate(templates)

	router.StaticFS("/assets", http.Dir(cfg.AssetsDir))
	router.StaticFile("/favicon.ico", fp.Join(cfg.AssetsDir, "img/favicon.webp"))

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
