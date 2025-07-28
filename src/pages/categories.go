package pages

import (
	"expenses/expenses"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	CategoriesPath = "/categories"
)

func categoriesGlobalPage(c *gin.Context) {
	categories, err := expenses.GetAllCategories()
	if err != nil {
		ServerErrorView(c, "The server too makes mistakes")
		return
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "categories",
		"renderNavBar": true,
		"content":      "categories",
		"categories":   categories,
	})
}

func categoryPage(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NotFoundView(c, "No such category")
		return
	}

	category, err := expenses.GetCategory(categoryID)
	if err != nil {
		ServerErrorView(c, "The server too makes mistakes")
		return
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "category",
		"renderNavBar": false,
		"content":      "category",
		"category":     category,
	})
}

func newCategoryPage(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "expenseNew",
		"renderNavBar": false,
		"content":      "categoryNew",
	})
}

func createCategory(c *gin.Context) {
	newCatName := c.PostForm("category-name")

	newCat := expenses.Category{
		CategoryName: newCatName,
	}

	err := newCat.Insert()
	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func deleteCategory(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NotFoundView(c, "No such category")
		return
	}

	category := expenses.Category{
		CategoryID: categoryID,
	}
	err = category.Delete()
	if err == expenses.ErrNotFound {
		errMsg := fmt.Sprintf("category %d not found", categoryID)
		c.String(http.StatusNotFound, errMsg)
		return
	}

	if err != nil {
		errMsg := fmt.Sprintf("error deleting category %d", categoryID)
		c.String(http.StatusInternalServerError, errMsg)
		return
	}

	c.Status(http.StatusNoContent)
}

func RouteCategories(router *gin.Engine) {
	router.GET(CategoriesPath, categoriesGlobalPage)
	router.GET(CategoriesPath+"/:id", categoryPage)
	router.POST(CategoriesPath, categoriesGlobalPage)

	router.GET(CategoriesPath+"/new", newCategoryPage)
	router.POST(CategoriesPath+"/new", createCategory)
	router.DELETE(CategoriesPath+"/:id", deleteCategory)
}
