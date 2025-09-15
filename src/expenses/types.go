package expenses

import (
	"context"
	mod "expenses/expenses/models"
)

// GetAllTypes returns all types
func (a *ExpApp) GetAllTypes(ctx context.Context) (mod.Types, error) {
	return a.TypeRepo.GetAll(ctx)
}

// GetType returns a type by its `id`
func (a *ExpApp) GetType(id int32) (mod.Type, error) {
	ctx := context.Background()
	return a.TypeRepo.Get(ctx, id)
}

// NewType creates a new type in the DB, or fails with error
func (a *ExpApp) NewType(name string) error {
	ctx := context.Background()

	newTyp := mod.Type{
		TypeName: name,
	}

	return a.TypeRepo.Insert(ctx, newTyp)
}

// DeleteType deletes a type by its `id`, or fails with error
func (a *ExpApp) DeleteType(id int32) error {
	ctx := context.Background()
	return a.TypeRepo.Delete(ctx, id)
}

// UpdateType updates a type by its `id`, or fails with error
func (a *ExpApp) UpdateType(id int32, name string) error {
	ctx := context.Background()

	updatedType := mod.Type{
		TypeID:   id,
		TypeName: name,
	}

	return a.TypeRepo.Update(ctx, updatedType)
}
