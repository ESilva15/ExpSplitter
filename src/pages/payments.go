package pages

import (
	"expenses/expenses"
	"strconv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func deletePayment(c *gin.Context) {
	paymentID, err := strconv.ParseInt(c.Param("id"), 10, 16)
	if err != nil {
		NotFoundView(c, "failed to fetch payment")
		return
	}

	payment := expenses.ExpensePayment{
		ExpPaymID: paymentID,
	} 

	err = payment.Delete()
	if err == expenses.ErrNotFound {
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
