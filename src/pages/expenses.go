package pages

import (
	"expenses/config"
	"expenses/expenses"
	"expenses/templating"
	"fmt"
	"log"
	"strconv"
	"time"

	"net/http"
	fp "path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	ExpensesPath = "/expenses"
)

func expensesContent() (map[string]any, error) {
	expenses, err := expenses.GetAllExpenses()
	if err != nil {
		return nil, err
	}

	data := map[string]any{
		"size":     len(expenses) - 1,
		"expenses": expenses,
		"resource": "expenses",
	}

	return data, nil
}

func expensesPartial(c *gin.Context) {
	cfg := config.GetInstance()

	expenses, err := expensesContent()
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "htmx/expenses.html"),
		expenses,
	)

	c.String(http.StatusOK, string(content))
}

func ExpensesGlobalPage(c *gin.Context) {
	cfg := config.GetInstance()

	expenses, err := expensesContent()
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "htmx/expenses.html"),
		expenses,
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expenses",
		"renderNavBar": true,
		"content":      content,
	})
}

func expensePage(c *gin.Context) {
	cfg := config.GetInstance()

	expenseID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}

	expense, err := expenses.GetExpense(expenseID)
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/expense.html"),
		map[string]any{
			"expense": expense,
		},
	)

	if expense.ExpID == -1 {
		NotFoundView(c, fmt.Sprintf("ID `%d` doesn't exist", expenseID))
		return
	}

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expense",
		"renderNavBar": false,
		"content":      content,
	})
}

func newExpensePage(c *gin.Context) {
	cfg := config.GetInstance()

	categories, err := expenses.GetAllCategories()
	if err != nil {
		ServerErrorView(c, "failed to fetch categories")
		return
	}

	stores, err := expenses.GetAllStores()
	if err != nil {
		ServerErrorView(c, "failed to fetch stores")
		return
	}

	types, err := expenses.GetAllTypes()
	if err != nil {
		ServerErrorView(c, "failed to fetch types")
		return
	}

	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/expenseNew.html"), map[string]any{
			"expense":    expenses.NewExpense(),
			"categories": categories,
			"stores":     stores,
			"types":      types,
		},
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expenseNew",
		"renderNavBar": false,
		"content":      content,
	})
}

func createExpense(c *gin.Context) {
	newDescription := c.PostForm("expense-desc")
	newDate := c.PostForm("expense-date")

	formattedDate, err := time.Parse("02-Jan-2006", newDate)
	if err != nil {
		// TODO
		// Change this to something the user can see
		c.Header("HX-Redirect", "/500")
		return
	}
	date := formattedDate.Unix()

	newValue := c.PostForm("expense-value")
	value, err := strconv.ParseFloat(newValue, 32)
	if err != nil {
		// TODO
		// Change this to something the user can see
		c.Header("HX-Redirect", "/500")
		return
	}

	newTyp := c.PostForm("newexp-type-dropdown")
	typID, err := strconv.Atoi(newTyp)
	if err != nil {
		// TODO
		// Change this to something the user can see
		log.Println("failed to parse newType:", newTyp)
		c.Header("HX-Redirect", "/500")
		return
	}

	newCat := c.PostForm("newexp-cat-dropdown")
	catID, err := strconv.Atoi(newCat)
	if err != nil {
		// TODO
		// Change this to something the user can see
		log.Println("failed to parse catID:", newCat)
		c.Header("HX-Redirect", "/500")
		return
	}

	newStore := c.PostForm("newexp-store-dropdown")
	storeID, err := strconv.Atoi(newStore)
	if err != nil {
		// TODO
		// Change this to something the user can see
		log.Println("failed to parse storeID:", newStore)
		c.Header("HX-Redirect", "/500")
		return
	}

	newExp := expenses.Expense{
		Description: newDescription,
		ExpDate:     date,
		Value:       float32(value),
		ExpType: expenses.Type{
			TypeID: typID,
		},
		ExpCategory: expenses.Category{
			CategoryID: catID,
		},
		ExpStore: expenses.Store{
			StoreID: storeID,
		},
	}
	log.Println(newExp)

	err = newExp.Insert()
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func RouteExpenses(router *gin.Engine) {
	router.GET(ExpensesPath, ExpensesGlobalPage)
	router.GET(ExpensesPath+"/:id", expensePage)
	router.POST(ExpensesPath, expensesPartial)

	router.GET(ExpensesPath+"/new", newExpensePage)
	router.POST(ExpensesPath+"/new", createExpense)
}
