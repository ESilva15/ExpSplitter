package cmd

import (
	"expenses/config"
	"expenses/pages"
	"github.com/spf13/cobra"

	"fmt"
	"html/template"
	"log"
	"net/http"
	fp "path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// This can be set in the compilation command - in the Makefile
var (
	ginMode     = "debug"
	tmplFuncMap = template.FuncMap{
		"formatDate": func(ts int64) string {
			return time.Unix(ts, 0).Format("02-Jan-2006")
		},
		"formatPrice": func(v float64) string {
			return fmt.Sprintf("%.2f", v)
		},
	}
)

func configGin(cfg *config.Configuration) *gin.Engine {
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

	return router
}

func setRoutes(router *gin.Engine) {
	pages.RouteHome(router)
	pages.RouteExpenses(router)
	pages.RouteCategories(router)
	pages.RouteStores(router)
	pages.RouteTypes(router)
	pages.RoutePayments(router)
	pages.RouteOverview(router)
	pages.RouteNotFound(router)
	pages.RouteServerError(router)
}

func startWebserver() {
	cfg := config.GetInstance()

	router := configGin(cfg)
	setRoutes(router)

	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Printf("Failed to launch application: \n  %s\n", err)
	}
}

func server(cmd *cobra.Command, args []string) {
	startWebserver()
}

// shelf
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	Args:  nil,
	Run:   server,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
