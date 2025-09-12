package auth

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	exp "expenses/expenses"
	mod "expenses/expenses/models"
)

func AuthMiddleware() gin.HandlerFunc {
	log.Println("Are we authing?")
	return func(c *gin.Context) {
		var user *mod.User
		var err error

		// Check if we have a cookie for the session
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err = exp.App.GetUser(uid.(int32))
			if err != nil {
				// TODO - what do we return here?
				return
			}
		}

		if user == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			// Just to be sure I recon
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
