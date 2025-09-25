package models

// User defines the model for the a User in the application
type User struct {
	UserID   int32  `json:"UserID"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

// Users defines a list of User
type Users []User

// NewUser returns a new empty user
func NewUser() User {
	return User{
		UserID:   -1,
		UserName: "",
	}
}
