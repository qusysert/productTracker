package service

import (
	"database/sql"
	"productTracker/internal/app/model"
)

func UpdateQuery(product model.Product, db *sql.DB, sellerId int) {
	_, err := db.Exec(
		updateStatement,
		product.Name,
		product.Price,
		product.Quantity,
		product.Available,
		sellerId,
		product.OfferId,
	)
	if err != nil {
		panic(err)
	}
}
