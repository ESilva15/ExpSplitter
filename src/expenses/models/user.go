package models

type User struct {
	UserID   int32  `json:"UserID"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}
type Users []User

func NewUser() User {
	return User{
		UserID:   -1,
		UserName: "",
	}
}
