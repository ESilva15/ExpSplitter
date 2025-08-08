package expenses

import (
	mod "expenses/expenses/models"
)

func GetAllTypes() ([]mod.Type, error) {
	return mod.GetAllTypes()
}

func GetType(id int64) (mod.Type, error) {
	return mod.GetType(id)
}

func NewType(name string) error {
	newTyp := mod.Type{
		TypeName: name,
	}
	return newTyp.Insert()
}

func DeleteType(id int64) error {
	typ := mod.Type{
		TypeID: id,
	}
	return typ.Delete()
}

func UpdateType(id int64, name string) error {
	newTyp := mod.Type{
		TypeID:   id,
		TypeName: name,
	}
	return newTyp.Update()
}
