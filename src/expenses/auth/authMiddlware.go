package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	exp "github.com/ESilva15/expenses/expenses"
	mod "github.com/ESilva15/expenses/expenses/models"
)

// SetUser sets the context data for the currently logged in user.
func SetUser(c *gin.Context, uID int32) error {
	var user *mod.User

	user, err := exp.App.GetUser(uID)
	if err != nil {
		return err
	}

	c.Set("user", user)

	return nil
}

// TODO - use the SetUser function here too somehow
// Middleware is the middleware we will use to verify if a user is
// authenticated or not.
func Middleware() gin.HandlerFunc {
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
