package cmd

import (
	"encoding/json"

	"github.com/ESilva15/expenses/api"
	"github.com/ESilva15/expenses/config"
	"github.com/ESilva15/expenses/expenses/auth"
	val "github.com/ESilva15/expenses/expenses/values"
	"github.com/ESilva15/expenses/pages"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"

	"html/template"
	"log"
	"net/http"
	fp "path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// TODO MOVE THIS REGISTERED FUNCTIONS SOMEWHER ELSE
// This can be set in the compilation command - in the Makefile
var (
	ginMode     = "debug"
	tmplFuncMap = template.FuncMap{
		"formatDate": func(t time.Time) string {
			return t.Format("02-Jan-2006")
		},
		"formatPrice": func(v decimal.Decimal) string {
			return v.Round(val.DecRange).StringFixed(val.DecRange)
		},
		"toJson": func(v any) template.JS {
			a, err := json.Marshal(v)
			if err != nil {
				log.Println("failed to marshal json data:", err)
				return template.JS("{}")
			}

			return template.JS(a)
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

func setRoutes(router *gin.RouterGroup, eng *gin.Engine) {
	pages.RouteHome(router)
	pages.RouteExpenses(router)
	pages.RouteCategories(router)
	pages.RouteStores(router)
	pages.RouteTypes(router)
	pages.RoutePayments(router)
	pages.RouteShares(router)
	pages.RouteOverview(router)
	pages.RouteAnalysis(router)

	// This are for the engine, I really need to rethink this
	pages.RouteLogin(eng)
	pages.RouteNotFound(eng)
	pages.RouteServerError(eng)
}

func setAPI(router *gin.RouterGroup) {
	api.RouteExpenses(router)
	api.RouteUsers(router)
	api.RouteStores(router)
	api.RouteTypes(router)
	api.RouteCategories(router)
}

func startWebserver() {
	cfg := config.GetInstance()

	router := configGin(cfg)
	// setRoutes(router)

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   int(time.Hour * 24),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	router.Use(sessions.Sessions("mysessions", store))

	// Set auth groups
	authGroup := router.Group("/")
	authGroup.Use(auth.AuthMiddleware())
	{
		setRoutes(authGroup, router)
	}

	// TODO - this also needs authentication
	apiGroup := router.Group(api.APIPath)
	{
		apiAuth := apiGroup.Group("/")
		apiAuth.Use(api.APIAuthMiddleWare())
		setAPI(apiAuth)
	}

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
