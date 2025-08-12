package pages

import (
	exp "expenses/expenses"
	experr "expenses/expenses/errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func deleteShare(c *gin.Context) {
	shareID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		NotFoundView(c, "failed to fetch payment")
		return
	}

	err = exp.Serv.DeleteShare(shareID)
	if err == experr.ErrNotFound {
		errMsg := fmt.Sprintf("share %d not found", shareID)
		c.String(http.StatusNotFound, errMsg)
		return
	}

	if err != nil {
		errMsg := fmt.Sprintf("error deleting share %d", shareID)
		c.String(http.StatusInternalServerError, errMsg)
		return
	}

	c.Status(http.StatusNoContent)
}

func RouteShares(router *gin.Engine) {
	router.DELETE("/shares/:id", deleteShare)
}
