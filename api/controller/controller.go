package controller

import (
	"FitaTechnicalTest-Golang/api/controller/c_shop"
	"FitaTechnicalTest-Golang/api/usecase"
	"database/sql"
)

type Controller struct {
	ShopController c_shop.ShopControllerInterface
}

func NewController(db *sql.DB) Controller {
	uc := usecase.NewUseCase(db)
	return Controller{
		ShopController: c_shop.NewShopController(uc),
	}
}
