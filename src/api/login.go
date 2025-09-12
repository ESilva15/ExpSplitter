package api

import (
	exp "expenses/expenses"
	"expenses/expenses/auth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type LoginData struct {
	Name string `json:"username"`
	Pass string `json:"password"`
}

func login(c *gin.Context) {
	var loginData LoginData
	err := c.ShouldBindJSON(&loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request - " + err.Error()})
		return
	}

	log.Println("user:", loginData.Name)
	log.Println("pass:", loginData.Pass)

	user, err := exp.App.ValidateCredentials(loginData.Name, loginData.Pass)
	if err != nil {
		log.Println("Error validating user: ", err)
		// TODO - go to an error page or login failed or something
		return
	}

	token, err := auth.GenerateToken(user.UserName)
	if err != nil {
		log.Println("Error generating token: ", err)
		// TODO - get error return
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func RouteLogin(eng *gin.Engine) {
	eng.POST("login", login)
}
