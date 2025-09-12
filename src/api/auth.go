package api

import (
	"net/http"
	"strings"

	expauth "expenses/expenses/auth"
	"github.com/gin-gonic/gin"
)

type TokenData struct {
	Token string `json:"JWTToken"`
}

func APIAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		validateToken, err := expauth.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !validateToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
			c.Abort()
			return
		}

		c.Next()
	}
}
