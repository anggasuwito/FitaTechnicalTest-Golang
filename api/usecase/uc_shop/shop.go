package uc_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"FitaTechnicalTest-Golang/api/repository"
)

type ShopUC struct {
	Repo repository.Repository
}

type ShopUCInterface interface {
	ItemListUC() (res []models.Item, err error)
	CartListUC(userID string) (res models.CartListResponse, err error)
	CartInsertUC(req models.CartInsertRequest) (res string, err error)
	CheckoutUC(userID string) (res models.CartListResponse, err error)
}

func NewShopUC(repo repository.Repository) ShopUCInterface {
	return &ShopUC{Repo: repo}
}
