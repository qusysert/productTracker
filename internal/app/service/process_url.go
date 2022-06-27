package service

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"productTracker/internal/app/model"
	"productTracker/internal/app/repository"
	"productTracker/internal/pkg/db"
	"productTracker/internal/pkg/downloader"
	"strconv"
)

func ProcessProductsFromURL(ctx context.Context, url string, sellerId int) (model.ProductListStat, error) {
	reader, err := downloader.DownloadReader(url)
	ret := model.ProductListStat{}
	if err != nil {
		fmt.Errorf("can't download the file: %v", err)
		return ret, err
	}
	defer reader.Close()

	products, err := TableToSlice(reader, sellerId)
	if err != nil {
		return ret, err
	}

	if err := repository.CreateSellerIfNotExists(ctx, sellerId); err != nil {
		return ret, err
	}
	if err := db.ExecTx(ctx, func(ctx context.Context) error {
		if err := repository.LockSeller(ctx, sellerId); err != nil {
			return err
		}
		for _, product := range products {
			if product.Available {
				exist, err := repository.IsProductExists(ctx, product)
				if err != nil {
					return err
				}
				if exist {
					if err := repository.UpdateProduct(ctx, product); err != nil {
						return err
					}
					ret.Updated++

				} else {
					if err := repository.InsertProduct(ctx, product); err != nil {
						return err
					}
					ret.Added++
				}
			} else {
				if err := repository.DeleteProduct(ctx, product); err != nil {
					return err
				}
				ret.Deleted++
			}
		}
		return nil
	}); err != nil {
		return ret, err
	}

	return ret, nil
}

func TableToSlice(reader io.Reader, sellerId int) ([]model.Product, error) {
	r, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %v", err)
	}

	// Get all the rows in the Sheet1.
	rows, err := r.GetRows("Sheet1")
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
				SellerId:  sellerId,
				OfferId:   offerId,
				Name:      name,
				Price:     price,
				Quantity:  quantity,
				Available: available,
			})
	}
	return products, nil
}
