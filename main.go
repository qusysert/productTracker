package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"productTracker/internal/app/pkg/config"
	"productTracker/internal/app/service"
)

func main() {
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

	service.AddProducts("https://drive.google.com/uc?export=download&id=1DMETpkF1UKgYsZ_KYRomea-4BMKxN4NK", db, 1333)
}
