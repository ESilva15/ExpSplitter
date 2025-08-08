package models

import repo "expenses/expenses/db/repository"

func MapRepoCategory(rc repo.Category) Category {
	return Category{
		CategoryID:   rc.CategoryID,
		CategoryName: rc.CategoryName,
	}
}

func MapRepoCategories(rc []repo.Category) []Category {
	categories := make([]Category, len(rc))
	for k, cat := range rc {
		categories[k] = MapRepoCategory(cat)
	}
	return categories
}
