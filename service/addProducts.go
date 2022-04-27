package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	offer_id  int
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

func tableToSlice(filepath string) ([]Product, error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Can't open file: %v", err)
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, fmt.Errorf("Can't get rowa: %v", err)
	}

	var rawData = make([]string, 0, 5)
	var products []Product

	for _, row := range rows {
		rawData = rawData[:0]
		for _, colCell := range row {
			rawData = append(rawData, colCell)
		}
		offerId, err := strconv.Atoi(rawData[0])
		name := rawData[1]
		price, err := strconv.ParseFloat(rawData[2], 64)
		quantity, err := strconv.Atoi(rawData[3])
		available, err := strconv.ParseBool(rawData[4])
		if err != nil {
			return nil, fmt.Errorf("conversion error: %v", err)
		}

		products = append(products,
			Product{
				offer_id:  offerId,
				name:      name,
				price:     price,
				quantity:  quantity,
				available: available,
			})
	}
	return products, nil
}

func DownloadFile(out *os.File, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Printf("Copy %v to %v\n", url, out.Name())
	return err
}
