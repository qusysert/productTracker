package service

import (
	"database/sql"
	"fmt"
	"os"
)

func ChangeTable(url string, db *sql.DB, sellerId int) {
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
		if product.Available {
			if isInTable(product, db, sellerId) {
				UpdateQuery(product, db, sellerId)
			} else {
				InsertQuery(product, db, sellerId)
			}
		} else {

			DeleteQuery(product, db, sellerId)
		}
	}
}
