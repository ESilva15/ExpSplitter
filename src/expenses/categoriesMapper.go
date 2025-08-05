package expenses

import (
	repo "expenses/expenses/db/repository"
)

func mapRepoCategory(rc repo.Category) Category {
	return Category{
		CategoryID:   rc.CategoryID,
		CategoryName: rc.CategoryName,
	}
}

func mapRepoCategories(rc []repo.Category) []Category {
	categories := make([]Category, len(rc))
	for k, cat := range rc {
		categories[k] = mapRepoCategory(cat)
	}
	return categories
}
