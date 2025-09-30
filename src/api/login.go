package api

import (
	"net/http"

	"github.com/ESilva15/expenses/expenses/auth"

	exp "github.com/ESilva15/expenses/expenses"
	"github.com/gin-gonic/gin"
)

type loginData struct {
	Name string `json:"username"`
	Pass string `json:"password"`
}

func login(c *gin.Context) {
	var loginData loginData
	err := c.ShouldBindJSON(&loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request - " + err.Error()})
		return
	}

	user, err := exp.App.ValidateCredentials(loginData.Name, loginData.Pass)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "bad credentials"})
		return
	}

	token, err := auth.GenerateToken(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func loginHelp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"help": "to login send a json like username, password",
	})
}

// RouteLogin routes the required API login endpoints
func RouteLogin(router *gin.RouterGroup) {
	router.GET("login", loginHelp)
	router.POST("login", login)
}
