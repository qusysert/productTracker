package handler

import (
	"net/http"
	"productTracker/internal/app/service"
)

type ImportRequest struct {
	Url      string `json:"url"`
	SellerId int    `json:"seller_id"`
}

type ImportResponse struct {
	Added   int `json:"added"`
	Updated int `json:"updated"`
	Deleted int `json:"deleted"`
	Failed  int `json:"failed"`
}

func ImportHandler(w http.ResponseWriter, r *http.Request) {
	handle(w, r, func(req ImportRequest) (ImportResponse, error) {
		stat, err := service.ProcessProductsFromURL(r.Context(), req.Url, req.SellerId)
		return ImportResponse{stat.Added, stat.Updated, stat.Deleted, stat.Failed}, err
	})
}
