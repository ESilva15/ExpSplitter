package models

type Category struct {
	CategoryID   int32  `json:"CategoryID"`
	CategoryName string `json:"CategoryName"`
}
type Categories []Category

func NewCategory() Category {
	return Category{
		CategoryID:   -1,
		CategoryName: "",
	}
}
