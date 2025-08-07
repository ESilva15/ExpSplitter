package expenses

import (
	mod "expenses/expenses/models"
)

func GetAllCategories() ([]mod.Category, error) {
	return mod.GetAllCategories()
}

func GetCategory(id int64) (mod.Category, error) {
	return mod.GetCategory(id)
}

func CreateCategory(name string) error {
	newCat := mod.Category{
		CategoryName: name,
	}

	return newCat.Insert()
}

func UpdateCategory(id int64, name string) error {
	cat := mod.Category{
		CategoryID:   id,
		CategoryName: name,
	}

	return cat.Update()
}

func DeleteCategory(id int64) error {
	cat := mod.Category{
		CategoryID: id,
	}

	return cat.Delete()
}
