package models

type Store struct {
	StoreID   int32  `json:"StoreID"`
	StoreName string `json:"StoreName"`
}
type Stores []Store

func NewStore() Store {
	return Store{
		StoreID:   -1,
		StoreName: "",
	}
}
