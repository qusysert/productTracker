package handler

import (
	"net/http"
	"productTracker/internal/app/model"
	"productTracker/internal/app/repository"
)

type GetRequest struct {
	SellerId int    `json:"seller_id"`
	OfferId  int    `json:"offer_id"`
	Name     string `json:"name"`
}

type GetResponse struct {
	Products []GetResponseProduct `json:"products"`
}

type GetResponseProduct struct {
	ID        int     `json:"id"`
	SellerId  int     `json:"seller_id"`
	OfferId   int     `json:"offer_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Available bool    `json:"available"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	handle(w, r, func(req GetRequest) (GetResponse, error) {
		filtered, err := repository.GetProductList(r.Context(), model.ProductFilter{SellerId: req.SellerId, OfferId: req.OfferId, Name: req.Name})
		if err != nil {
			return GetResponse{}, err
		}

		products := make([]GetResponseProduct, 0, len(filtered))
		for _, p := range filtered {
			products = append(products, GetResponseProduct{
				ID:        p.ID,
				SellerId:  p.SellerId,
				OfferId:   p.OfferId,
				Name:      p.Name,
				Price:     p.Price,
				Quantity:  p.Quantity,
				Available: p.Available,
			})
		}
		return GetResponse{products}, nil
	})

}
