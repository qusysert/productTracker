package repository

import (
	"context"
	"fmt"
	"productTracker/internal/app/model"
	"productTracker/internal/pkg/db"
	"strings"
)

func DeleteProduct(ctx context.Context, product model.Product) error {
	_, err := db.FromContext(ctx).Exec(
		`DELETE FROM product WHERE seller_id = $1 AND offer_id = $2;`,
		product.SellerId,
		product.OfferId,
	)
	return err
}

func InsertProduct(ctx context.Context, product model.Product) error {

	_, err := db.FromContext(ctx).Exec(
		`
			INSERT INTO product (seller_id, offer_id, name, price, quantity, available)
			VALUES ($1, $2, $3, $4, $5, $6);`,
		product.SellerId,
		product.OfferId,
		product.Name,
		product.Price,
		product.Quantity,
		product.Available,
	)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(ctx context.Context, product model.Product) error {
	_, err := db.FromContext(ctx).Exec(
		`UPDATE product SET name = $1, price = $2, quantity = $3, available = $4
			WHERE seller_id = $5 AND offer_id = $6`,
		product.Name,
		product.Price,
		product.Quantity,
		product.Available,
		product.SellerId,
		product.OfferId,
	)
	return err
}

func IsProductExists(ctx context.Context, product model.Product) (bool, error) {
	rows, err := db.FromContext(ctx).Query(`SELECT 1 FROM product WHERE seller_id = $1 AND offer_id = $2`, product.SellerId, product.OfferId)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return rows.Next(), nil
}

func GetProductList(ctx context.Context, filter model.ProductFilter) ([]model.Product, error) {
	q := "SELECT id, seller_id, offer_id, name, price, quantity, available FROM product"
	where := make([]string, 0)

	if filter.SellerId != 0 {
		where = append(where, fmt.Sprintf("seller_id = %d", filter.SellerId))
	}
	if filter.OfferId != 0 {
		where = append(where, fmt.Sprintf("offer_id = %d", filter.OfferId))
	}
	if filter.Name != "" {
		where = append(where, fmt.Sprintf("name LIKE %s%%", filter.Name))
	}

	if len(where) != 0 {
		q += " WHERE " + strings.Join(where, " AND ")
	}

	res, err := db.FromContext(ctx).Query(q)
	if err != nil {
		return nil, err
	}

	ret := make([]model.Product, 0)

	for res.Next() {
		var product model.Product
		err := res.Scan(&product.ID, &product.SellerId, &product.OfferId, &product.Name, &product.Price, &product.Quantity, &product.Available)
		if err != nil {
			return nil, err
		}
		ret = append(ret, product)
	}
	return ret, nil
}
