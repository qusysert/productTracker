package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"productTracker/internal/app/pkg/config"
)

func OpenConnection() *sql.DB {
	// Loading cfg
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load cfg:", err)
	}

	// Configuring connection
	psqlInfo := fmt.Sprintf("host=%v port=%d user=%v "+
		"password=%v dbname=%v sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Opening a connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	db.SetMaxOpenConns(5)

	return db
}
