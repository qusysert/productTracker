package repository

import (
	"context"
	"fmt"
	"productTracker/internal/app/model"
	"productTracker/internal/pkg/db"
)

func DeleteProduct(ctx context.Context, product model.Product) {
	_, err := db.FromContext(ctx).Exec(
		`DELETE FROM products WHERE sellerId = $1 AND offerid = $2;`,
		product.SellerId,
		product.OfferId,
	)
	if err != nil {
		panic(err)
	}
}

func InsertProduct(ctx context.Context, product model.Product) {
	_, err := db.FromContext(ctx).Exec(
		`
			INSERT INTO products (sellerid, offerid, name, price, quantity, available)
			VALUES ($1, $2, $3, $4, $5, $6);`,
		product.SellerId,
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

func UpdateProduct(ctx context.Context, product model.Product) {
	res, err := db.FromContext(ctx).Exec(
		`UPDATE products SET name = $1, price = $2, quantity = $3, available = $4
			WHERE sellerid = $5 AND offerid = $6;`,
		product.Name,
		product.Price,
		product.Quantity,
		product.Available,
		product.SellerId,
		product.OfferId,
	)
	fmt.Println(res)
	if err != nil {
		panic(err)
	}
}

func IsProductExists(ctx context.Context, product model.Product) bool {
	var isExists bool
	_, err := db.FromContext(ctx).Query(`SELECT EXISTS(SELECT 1 FROM products WHERE sellerid = $1 AND offerid = $2)`, &isExists)
	if err != nil {
		panic(err)
	}

	return isExists
}
