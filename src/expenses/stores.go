package expenses

import (
	mod "expenses/expenses/models"
)

func (s *Service) GetAllStores() ([]mod.Store, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return []mod.Store{}, err
	}
	defer tx.Rollback()

	stores, err := mod.GetAllStores(tx)
	if err != nil {
		return []mod.Store{}, err
	}

	return stores, tx.Commit()
}

func (s *Service) GetStore(id int64) (mod.Store, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return mod.Store{}, err
	}
	defer tx.Rollback()

	store, err := mod.GetStore(tx, id)
	if err != nil {
		return mod.Store{}, err
	}

	return store, tx.Commit()
}

func (s *Service) NewStore(name string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	newStore := mod.Store{
		StoreName: name,
	}

	err = newStore.Insert(tx)
	return tx.Commit()
}

func (s *Service) UpdateStore(id int64, name string) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	store := mod.Store{
		StoreID:   id,
		StoreName: name,
	}

	err = store.Update(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Service) DeleteStore(id int64) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	store := mod.Store{
		StoreID: id,
	}

	err = store.Delete(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
