package repo

import (
	mod "expenses/expenses/models"
	"expenses/expenses/repo/pgdb/pgsqlc"
)

func mapRepoStore(rs pgsqlc.Store) mod.Store {
	return mod.Store{
		StoreID:   rs.StoreID,
		StoreName: rs.StoreName,
		StoreNIF:  rs.NIF,
	}
}

func mapRepoStores(rs []pgsqlc.Store) mod.Stores {
	stores := make(mod.Stores, len(rs))
	for k, store := range rs {
		stores[k] = mapRepoStore(store)
	}
	return stores
}
