package pages

import (
	"expenses/config"
	"expenses/expenses"
	"expenses/templating"

	"fmt"
	"net/http"
	fp "path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	CategoriesPath = "/categories"
)

func categoriesContent() (map[string]any, error) {
	categories, err := expenses.GetAllCategories()
	if err != nil {
		return nil, err
	}

	data := map[string]any{
		"size":       len(categories) - 1,
		"categories": categories,
		"resource":   "categories",
	}

	return data, nil
}

func categoriesGlobalPage(c *gin.Context) {
	cfg := config.GetInstance()

	categories, err := categoriesContent()
	if err != nil {
		c.Header("HX-Redirect", "/500")
		return
	}
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "htmx/categories.html"),
		categories,
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "categories",
		"renderNavBar": true,
		"content":      content,
	})
}

func categoryPage(c *gin.Context) {
	cfg := config.GetInstance()

	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Redirect(404, "/404")
	}

	category, err := expenses.GetCategory(categoryID)
	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/category.html"),
		map[string]any{
			"category": category,
		},
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "category",
		"renderNavBar": false,
		"content":      content,
	})
}

func newCategoryPage(c *gin.Context) {
	cfg := config.GetInstance()

	content := templating.HtmlTemplate(
		fp.Join(cfg.AssetsDir, "/htmx/categoryNew.html"), map[string]any{},
	)

	c.HTML(http.StatusOK, "terminal.gotempl", gin.H{
		"page":         "expenseNew",
		"renderNavBar": false,
		"content":      content,
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
		c.Redirect(404, "/404")
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
