package service

import (
	"database/sql"
	"productTracker/internal/app/model"
)

func InsertQuery(product model.Product, db *sql.DB, sellerId int) {
	_, err := db.Exec(
		insertStatement,
		sellerId,
		product.OfferId,
		product.Name,
		product.Price,
		product.Quantity,
		product.Available,
	)
	if err != nil {
		panic(err)
	}
}
