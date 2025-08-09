package expenses

import (
	"database/sql"
	"expenses/config"
	"log"
)

var (
	Serv *Service
)

type Service struct {
	DB *sql.DB
}

func NewExpenseService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func StartApp() error {
	config.SetConfig("./config.yaml")
	cfg := config.GetInstance()

	db, err := sql.Open(cfg.DBSys, "file:"+cfg.DBPath+"?_foreign_keys=on")
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}

	Serv = NewExpenseService(db)
	return nil
}
