package pages

import (
	"database/sql"
	"expenses/expenses"
	experr "expenses/expenses/errors"

	"encoding/json"
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
	categoryID, err := strconv.ParseInt(c.Param("id"), 10, 16)
	if err != nil {
		NotFoundView(c, "No such category")
		return
	}

	category, err := expenses.GetCategory(categoryID)
	if err == sql.ErrNoRows {
		NotFoundView(c, fmt.Sprintf("Could not find category `%d`", categoryID))
		return
	}
	if err != nil {
		ServerErrorView(c, "The server too makes mistakes")
		return
	}

	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "category",
		"renderNavBar": false,
		"content":      "category",
		"category":     category,
		"method":       "put",
	})
}

func newCategoryPage(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal", gin.H{
		"page":         "categoryNew",
		"renderNavBar": false,
		"content":      "category",
		"method":       "post",
	})
}

func createCategory(c *gin.Context) {
	newCatName := c.PostForm("category-name")
	err := expenses.CreateCategory(newCatName)

	if err != nil {
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":\"%s\"}", err.Error()))
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func updateCategory(c *gin.Context) {
	categoryID, err := expenses.ParseID(c.Param("id"))
	if err != nil {
		NotFoundView(c, "No such category")
		return
	}
	categoryName := c.PostForm("category-name")

	err = expenses.UpdateCategory(categoryID, categoryName)
	if err != nil {
		errMsg, _ := json.Marshal(err.Error())
		c.Header("HX-Trigger", fmt.Sprintf("{\"formState\":%s}", errMsg))
		return
	}

	c.Header("HX-Trigger", "{\"formState\":\"Success\"}")
}

func deleteCategory(c *gin.Context) {
	categoryID, err := expenses.ParseID(c.Param("id"))
	if err != nil {
		NotFoundView(c, "No such category")
		return
	}

	err = expenses.DeleteCategory(categoryID)
	if err == experr.ErrNotFound {
		errMsg := fmt.Sprintf("category %d not found", categoryID)
		c.String(http.StatusNotFound, errMsg)
		return
	}

	if err != nil {
		errMsg := fmt.Sprintf("error deleting category %d: %+v", categoryID,
			err.Error())
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
	router.PUT(CategoriesPath+"/:id", updateCategory)
}
