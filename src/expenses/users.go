package expenses

import (
	"expenses/config"

	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UserID   int64
	UserName string
}

func NewUser() User {
	return User{
		UserID:   -1,
		UserName: "",
	}
}

func GetAllUsers() ([]User, error) {
	cfg := config.GetInstance()

	db, err := sql.Open(cfg.DBSys, cfg.DBPath)
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
