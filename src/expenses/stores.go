package expenses

type Store struct {
	StoreID   int
	StoreName string
}

func NewStore() Store {
	return Store{
		StoreID:   -1,
		StoreName: "",
	}
}
