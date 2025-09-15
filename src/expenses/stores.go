package expenses

import (
	"context"
	mod "expenses/expenses/models"
)

func (a *ExpApp) GetAllStores(ctx context.Context) ([]mod.Store, error) {
	return a.StoreRepo.GetAll(ctx)
}

func (a *ExpApp) GetStore(id int32) (mod.Store, error) {
	ctx := context.Background()
	return a.StoreRepo.Get(ctx, id)
}

func (a *ExpApp) NewStore(name string, nif string) error {
	ctx := context.Background()

	newStore := mod.Store{
		StoreName: name,
		StoreNIF:  nif,
	}

	return a.StoreRepo.Insert(ctx, newStore)
}

func (a *ExpApp) UpdateStore(id int32, name string, nif string) error {
	ctx := context.Background()

	store := mod.Store{
		StoreID:   id,
		StoreName: name,
		StoreNIF:  nif,
	}

	return a.StoreRepo.Update(ctx, store)
}

func (a *ExpApp) DeleteStore(id int32) error {
	ctx := context.Background()
	return a.StoreRepo.Delete(ctx, id)
}

func (a *ExpApp) GetStoreIDFromNIF(nif string) (int32, error) {
	ctx := context.Background()

	store, err := a.StoreRepo.GetByNIF(ctx, nif)
	if err != nil {
		return -1, err
	}

	return store.StoreID, nil
}
