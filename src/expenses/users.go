package expenses

type User struct {
	UserID   int
	UserName string
}

func NewUser() User {
	return User{
		UserID:   -1,
		UserName: "",
	}
}
