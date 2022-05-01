package service

import (
	"fmt"
	"os"
)

type Product struct {
	offerId   int
	name      string
	price     float64
	quantity  int
	available bool
}

func AddProducts(url string, sellerId int) {
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

	products, e := tableToSlice(filepath)
	if e != nil {
		panic(e)
	}
	fmt.Println(products[0].name)
}
