package main

import (
	"database/sql"
	"log"
	"net/http"
	"productTracker/internal/app/handler"
	"productTracker/internal/app/pkg/config"
	"productTracker/internal/app/repository"
	"productTracker/internal/pkg/db"
)

var conn *sql.DB

func main() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load cfg:", err)
	}

	conn = db.New(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err := repository.Migrate(conn); err != nil {
		log.Fatal("Can't migrate: %w", err.Error())
	}

	for _, rec := range [...]struct {
		route   string
		handler http.HandlerFunc
	}{
		{route: "/import", handler: handler.ImportHandler},
		{route: "/get", handler: handler.GetHandler},
	} {
		http.HandleFunc(rec.route, DbMiddleware(rec.handler))
	}

	log.Printf("Server started on port %s \n", cfg.HttpPort)
	err = http.ListenAndServe(":"+cfg.HttpPort, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func DbMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r.WithContext(db.AddToContext(r.Context(), &db.Db{conn})))
	}
}
