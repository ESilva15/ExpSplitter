package api

import (
	"net/http"
	"strings"

	expauth "github.com/ESilva15/expenses/expenses/auth"

	"github.com/gin-gonic/gin"
)

type tokenData struct {
	Token string `json:"JWTToken"`
}

// AuthMiddleWare will handle this API authentication.
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		claims, err := expauth.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
			c.Abort()
			return
		}

		// Set the user from the claims
		userID := int32(claims["sub"].(float64))
		err = expauth.SetUser(c, userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
			c.Abort()
			return
		}

		c.Next()
	}
}
