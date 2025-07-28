package main

import (
	"expenses/config"
	"expenses/pages"

	"fmt"
	"html/template"
	"log"
	"net/http"
	fp "path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// This can be set in the compilation command - in the Makefile
var ginMode = "debug"

var tmplFuncMap = template.FuncMap{
	"formatDate": func(ts int64) string {
		return time.Unix(ts, 0).Format("02-Jan-2006")
	},
	"formatPrice": func(v float32) string {
		return fmt.Sprintf("%.2f", v)
	},
}

func main() {
	config.SetConfig("./config.yaml")
	cfg := config.GetInstance()

	gin.SetMode(ginMode)
	templates := template.Must(
		template.New("").
			Funcs(tmplFuncMap).
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
