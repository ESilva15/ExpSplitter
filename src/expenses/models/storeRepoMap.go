package models

import repo "expenses/expenses/db/repository"

func mapRepoStore(rs repo.Store) Store {
	return Store{
		StoreID:   rs.StoreID,
		StoreName: rs.StoreName,
	}
}

func mapRepoStores(rs []repo.Store) []Store {
	stores := make([]Store, len(rs))
	for k, store := range rs {
		stores[k] = mapRepoStore(store)
	}
	return stores
}
