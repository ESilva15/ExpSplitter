package expenses

import (
	mod "expenses/expenses/models"
)

func (a *ExpensesApp) GetAllTypes() ([]mod.Type, error) {
	tx, err := a.DB.Begin()
	if err != nil {
		return []mod.Type{}, err
	}
	defer tx.Rollback()

	types, err := mod.GetAllTypes(tx)
	if err != nil {
		return []mod.Type{}, err
	}

	return types, tx.Commit()
}

func (a *ExpensesApp) GetType(id int64) (mod.Type, error) {
	tx, err := a.DB.Begin()
	if err != nil {
		return mod.Type{}, err
	}
	defer tx.Rollback()

	typ, err := mod.GetType(tx, id)
	if err != nil {
		return mod.Type{}, err
	}

	return typ, tx.Commit()
}

func (a *ExpensesApp) NewType(name string) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	newTyp := mod.Type{
		TypeName: name,
	}

	err = newTyp.Insert(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *ExpensesApp) DeleteType(id int64) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	typ := mod.Type{
		TypeID: id,
	}

	err = typ.Delete(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (a *ExpensesApp) UpdateType(id int64, name string) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	newTyp := mod.Type{
		TypeID:   id,
		TypeName: name,
	}

	err = newTyp.Update(tx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
