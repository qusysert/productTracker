package test

import (
	"productTracker/internal/app/model"
	"productTracker/internal/app/service"
	"productTracker/internal/pkg"
)

var product = model.Product{}

func TestUpdateTable() {
	db := pkg.OpenConnection()

	service.UpdateQuery()
}
