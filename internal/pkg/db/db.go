package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const ctxName = "db"

func AddToContext(ctx context.Context, conn *sql.DB) context.Context {
	return context.WithValue(ctx, ctxName, conn)
}

func FromContext(ctx context.Context) *sql.DB {
	conn, ok := ctx.Value(ctxName).(*sql.DB)
	if !ok {
		panic("Not found db in context")
	}
	return conn
}

func New(host string, port int, user, password, name string) *sql.DB {
	// Configuring connection
	psqlInfo := fmt.Sprintf("host=%v port=%d user=%v "+
		"password=%v dbname=%v sslmode=disable",
		host, port, user, password, name)

	// Opening a connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	db.SetMaxOpenConns(5)

	return db
}
