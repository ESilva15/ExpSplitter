package repo

import (
	mod "github.com/ESilva15/expenses/expenses/models"
	pgsqlc "github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
)

func mapRepoCategory(rc pgsqlc.Category) mod.Category {
	return mod.Category{
		CategoryID:   rc.CategoryID,
		CategoryName: rc.CategoryName,
	}
}

func mapRepoCategories(rc []pgsqlc.Category) mod.Categories {
	categories := make(mod.Categories, len(rc))
	for k, cat := range rc {
		categories[k] = mapRepoCategory(cat)
	}

	return categories
}
