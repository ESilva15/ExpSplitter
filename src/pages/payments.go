package pages

import (
	"fmt"
	"net/http"

	exp "github.com/ESilva15/expenses/expenses"
	experr "github.com/ESilva15/expenses/expenses/errors"
	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/gin-gonic/gin"
	dec "github.com/shopspring/decimal"
)

func deletePayment(c *gin.Context) {
	paymentID, err := exp.ParseID(c.Param("id"))
	if err != nil {
		NotFoundView(c, "failed to fetch payment")
		return
	}

	err = exp.App.DeletePayment(paymentID)
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

func addPayment(c *gin.Context) {
	expID, err := exp.ParseID(c.PostForm("ExpID"))
	if err != nil {
	}

	debtorID, err := exp.ParseID(c.PostForm("DebtorID"))
	if err != nil {
	}

	creditorID, err := exp.ParseID(c.PostForm("CreditorID"))
	if err != nil {
	}

	sum, err := dec.NewFromString(c.PostForm("Sum"))
	if err != nil {
	}

	debt := mod.Debt{
		Debtor: mod.User{
			UserID: debtorID,
		},
		Creditor: mod.User{
			UserID: creditorID,
		},
		Sum: sum,
	}

	err = exp.App.ProcessDebt(expID, debt)
	if err != nil {
	}
}

func RoutePayments(router *gin.RouterGroup) {
	router.DELETE("/payments/:id", deletePayment)

	router.POST("/payments/add", addPayment)
}
