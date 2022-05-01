package service

import (
	"database/sql"
	"fmt"
	"os"
)

func AddProducts(url string, db *sql.DB, sellerId int) {
	var tmpFile, err = os.CreateTemp("/tmp", "*.xlsx")
	if err != nil {
		panic(err)
	}
	filepath := tmpFile.Name()

	err = DownloadFile(tmpFile, url)
	if err != nil {
		fmt.Errorf("can't download the file: %v", err)
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Errorf("can't remove the file: %v", err)
		}
	}(filepath)

	products, err := tableToSlice(filepath)
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		insertStatement := `
			INSERT INTO products (sellerId, offerId, name, price, quantity, available)
			VALUES ($1, $2, $3, $4, $5, $6)`
		_, err = db.Exec(
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

}
