package models

type Type struct {
	TypeID   int32  `json:"TypeID"`
	TypeName string `json:"TypeName"`
}
type Types []Type

func NewType() Type {
	return Type{
		TypeID:   -1,
		TypeName: "",
	}
}
