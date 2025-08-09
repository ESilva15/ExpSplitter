package expenses

import (
	"database/sql"
	"expenses/config"
	mod "expenses/expenses/models"
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

func openDB(sys string, path string, extra string) (*sql.DB, error) {
	db, err := sql.Open(sys, "file:"+path+extra)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func StartApp() error {
	config.SetConfig("./config.yaml")
	cfg := config.GetInstance()

	migDB, err := openDB(cfg.DBSys, cfg.DBPath, "")
	if err != nil {
		log.Fatalf("Failed to open migration DB: %v", err)
	}

	err = mod.RunMigrations(migDB)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	migDB.Close()

	db, err := openDB(cfg.DBSys, cfg.DBPath, "?_foreign_keys=on")
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	Serv = NewExpenseService(db)

	return nil
}
