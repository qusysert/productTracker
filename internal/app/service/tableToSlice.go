package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"productTracker/internal/app/model"
	"strconv"
)

func tableToSlice(filepath string) ([]model.Product, error) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %v", err)
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, fmt.Errorf("can't get row: %v", err)
	}

	var rawData = make([]string, 0, 5)
	var products []model.Product

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
			model.Product{
				OfferId:   offerId,
				Name:      name,
				Price:     price,
				Quantity:  quantity,
				Available: available,
			})
	}
	return products, nil
}
