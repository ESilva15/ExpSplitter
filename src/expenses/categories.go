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

func (s *Service) CreateCategory(name string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	newCat := mod.Category{
		CategoryName: name,
	}

	err = newCat.Insert(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Service) UpdateCategory(id int64, name string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	cat := mod.Category{
		CategoryID:   id,
		CategoryName: name,
	}
	err = cat.Update()
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Service) DeleteCategory(id int64) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	cat := mod.Category{
		CategoryID: id,
	}

	err = cat.Delete()
	if err != nil {
		return err
	}

	return tx.Commit()
}
