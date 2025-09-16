package repo

import (
	mod "github.com/ESilva15/expenses/expenses/models"
	"github.com/ESilva15/expenses/expenses/repo/pgdb/pgsqlc"
)

func mapRepoType(rt pgsqlc.ExpenseType) mod.Type {
	return mod.Type{
		TypeID:   rt.TypeID,
		TypeName: rt.TypeName,
	}
}

func mapRepoTypes(rt []pgsqlc.ExpenseType) mod.Types {
	types := make(mod.Types, len(rt))
	for k, typ := range rt {
		types[k] = mapRepoType(typ)
	}
	return types
}
