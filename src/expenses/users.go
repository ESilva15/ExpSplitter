package expenses

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

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

func GetAllUsers() ([]User, error) {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT UserID,UserName " +
		"FROM users"

	var userList []User
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.UserID, &user.UserName)
		if err != nil {
			log.Fatalf("Failed to parse data from db: %v", err)
		}
		userList = append(userList, *user)
	}

	return userList, nil
}
