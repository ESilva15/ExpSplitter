package repo

import (
	mod "expenses/expenses/models"
	pgsqlc "expenses/expenses/repo/pgdb/pgsqlc"
)

func MapRepoCategory(rc pgsqlc.Category) mod.Category {
	return mod.Category{
		CategoryID:   rc.CategoryID,
		CategoryName: rc.CategoryName,
	}
}

func MapRepoCategories(rc []pgsqlc.Category) mod.Categories {
	categories := make(mod.Categories, len(rc))
	for k, cat := range rc {
		categories[k] = MapRepoCategory(cat)
	}
	return categories
}
