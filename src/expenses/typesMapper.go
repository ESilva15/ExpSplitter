package expenses

import repo "expenses/expenses/db/repository"

func mapRepoType(rt repo.ExpenseType) Type {
	return Type{
		TypeID:   rt.TypeID,
		TypeName: rt.TypeName,
	}
}

func mapRepoTypes(rt []repo.ExpenseType) []Type {
	types := make([]Type, len(rt))
	for k, typ := range rt {
		types[k] = mapRepoType(typ)
	}
	return types
}
