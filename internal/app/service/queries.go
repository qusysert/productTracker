package service

var insertStatement = `
			INSERT INTO products (sellerId, offerId, name, price, quantity, available)
			VALUES ($1, $2, $3, $4, $5, $6);`
var updateStatement = `
			UPDATE products SET (name = $2, price = $3, quantity = $4, available = $5)
			WHERE sellerId = $6, offerId = $7;`
var deleteStatement = `DELETE FROM products WHERE sellerId = $1, offerId = $2;`
