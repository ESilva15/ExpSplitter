package expenses

import (
	mod "expenses/expenses/models"
)

func (a *ExpensesApp) GetAllStores() ([]mod.Store, error) {
	tx, err := a.DB.Begin()
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

func (a *ExpensesApp) GetStore(id int64) (mod.Store, error) {
	tx, err := a.DB.Begin()
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

func (a *ExpensesApp) NewStore(name string) error {
	tx, err := a.DB.Begin()
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

func (a *ExpensesApp) UpdateStore(id int64, name string) error {
	tx, err := a.DB.Begin()
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

func (a *ExpensesApp) DeleteStore(id int64) error {
	tx, err := a.DB.Begin()
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
