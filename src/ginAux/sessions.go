package ginaux

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetLoggedInUserCTX returns the contextual user
func GetLoggedInUserCTX(c *gin.Context) (context.Context, error) {
	if _, ok := c.Get("user"); !ok {
		return nil, fmt.Errorf("unable to get user from gin context")
	}

	user, _ := c.Get("user")
	ctx := context.WithValue(c.Request.Context(), "user", user)

	return ctx, nil
}
