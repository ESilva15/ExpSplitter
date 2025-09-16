package pages

import (
	"net/http"

	exp "github.com/ESilva15/expenses/expenses"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	LoginPath = "/login"
)

func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "expenses",
		"renderNavBar": false,
		"content":      "login",
	})
}

func loginUser(c *gin.Context) {
	username := c.PostForm("login-user")
	password := c.PostForm("login-pass")

	user, err := exp.App.ValidateCredentials(username, password)
	if err != nil {
		// TODO - go to an error page or login failed or something
		return
	}

	// Store new session
	session := sessions.Default(c)
	session.Set("user_id", user.UserID)
	if err := session.Save(); err != nil {
		// TODO - error storing session
		return
	}

	c.Header("Hx-Redirect", "/")
	c.Status(http.StatusOK)
}

func RouteLogin(router *gin.Engine) {
	router.GET(LoginPath, loginPage)
	router.POST(LoginPath, loginUser)
}
