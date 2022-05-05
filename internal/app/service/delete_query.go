package service

import (
	"database/sql"
	"productTracker/internal/app/model"
)

func DeleteQuery(product model.Product, db *sql.DB, sellerId int) {
	_, err := db.Exec(
		deleteStatement,
		sellerId,
		product.OfferId,
	)
	if err != nil {
		panic(err)
	}
}
