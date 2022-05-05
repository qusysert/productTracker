package service

import (
	"database/sql"
	"productTracker/internal/app/model"
)

func isInTable(product model.Product, db *sql.DB, sellerId int) bool {
	selectStatement := `SELECT * FROM products WHERE sellerId = $1 AND offerId = $2;`
	res, err := db.Exec(selectStatement, sellerId, product.OfferId)
	if err != nil {
		panic(err)
	}

	// TODO: check this
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return false
	}
	return true
}
