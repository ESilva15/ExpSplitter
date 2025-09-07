package models

type User struct {
	UserID   int32  `json:"UserID"`
	UserName string `json:"UserName"`
}
type Users []User

func NewUser() User {
	return User{
		UserID:   -1,
		UserName: "",
	}
}
