package expenses

import (
	"context"
	mod "expenses/expenses/models"
)

func (a *ExpensesApp) GetAllTypes() (mod.Types, error) {
	ctx := context.Background()
	return a.TypeRepo.GetAll(ctx)
}

func (a *ExpensesApp) GetType(id int32) (mod.Type, error) {
	ctx := context.Background()
	return a.TypeRepo.Get(ctx, id)
}

func (a *ExpensesApp) NewType(name string) error {
	ctx := context.Background()

	newTyp := mod.Type{
		TypeName: name,
	}

	return a.TypeRepo.Insert(ctx, newTyp)
}

func (a *ExpensesApp) DeleteType(id int32) error {
	ctx := context.Background()
	return a.TypeRepo.Delete(ctx, id)
}

func (a *ExpensesApp) UpdateType(id int32, name string) error {
	ctx := context.Background()

	updatedType := mod.Type{
		TypeID:   id,
		TypeName: name,
	}

	return a.TypeRepo.Update(ctx, updatedType)
}
