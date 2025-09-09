package expenses

import (
	"context"
	mod "expenses/expenses/models"
)

func (a *ExpensesApp) GetAllStores() ([]mod.Store, error) {
	ctx := context.Background()
	return a.StoreRepo.GetAll(ctx)
}

func (a *ExpensesApp) GetStore(id int32) (mod.Store, error) {
	ctx := context.Background()
	return a.StoreRepo.Get(ctx, id)
}

func (a *ExpensesApp) NewStore(name string, nif string) error {
	ctx := context.Background()

	newStore := mod.Store{
		StoreName: name,
		StoreNIF:  nif,
	}

	return a.StoreRepo.Insert(ctx, newStore)
}

func (a *ExpensesApp) UpdateStore(id int32, name string, nif string) error {
	ctx := context.Background()

	store := mod.Store{
		StoreID:   id,
		StoreName: name,
		StoreNIF:  nif,
	}

	return a.StoreRepo.Update(ctx, store)
}

func (a *ExpensesApp) DeleteStore(id int32) error {
	ctx := context.Background()
	return a.StoreRepo.Delete(ctx, id)
}

func (a *ExpensesApp) GetStoreIDFromNIF(nif string) (int32, error) {
	ctx := context.Background()

	store, err := a.StoreRepo.GetByNIF(ctx, nif)
	if err != nil {
		return -1, err
	}

	return store.StoreID, nil
}
