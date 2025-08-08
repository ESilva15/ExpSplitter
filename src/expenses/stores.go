package expenses

import (
	mod "expenses/expenses/models"
)

func GetAllStores() ([]mod.Store, error) {
	return mod.GetAllStores()
}

func GetStore(id int64) (mod.Store, error) {
	return mod.GetStore(id)
}

func NewStore(name string) error {
	newStore := mod.Store{
		StoreName: name,
	}
	return newStore.Insert()
}

func UpdateStore(id int64, name string) error {
	store := mod.Store{
		StoreID:   id,
		StoreName: name,
	}

	return store.Update()
}

func DeleteStore(id int64) error {
	store := mod.Store{
		StoreID: id,
	}
	return store.Delete()
}
