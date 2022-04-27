package service

import (
	"fmt"
	"io"
	"net/http"
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

func DownloadFile(out *os.File, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Errorf("can't close Body after downloading: %v", err)
		}
	}(resp.Body)

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Printf("Copy %v to %v\n", url, out.Name())
	return err
}
