package pages

import (
	"expenses/expenses"
	experr "expenses/expenses/errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func deletePayment(c *gin.Context) {
	paymentID, err := expenses.ParseID(c.Param("id"))
	if err != nil {
		NotFoundView(c, "failed to fetch payment")
		return
	}

	err = expenses.DeletePayment(paymentID)
	if err == experr.ErrNotFound {
		errMsg := fmt.Sprintf("category %d not found", paymentID)
		c.String(http.StatusNotFound, errMsg)
		return
	}

	if err != nil {
		errMsg := fmt.Sprintf("error deleting category %d", paymentID)
		c.String(http.StatusInternalServerError, errMsg)
		return
	}

	c.Status(http.StatusNoContent)
}

func RoutePayments(router *gin.Engine) {
	router.DELETE("/payments/:id", deletePayment)
}
