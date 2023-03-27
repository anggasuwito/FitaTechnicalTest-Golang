package repository

import (
	"FitaTechnicalTest-Golang/api/repository/r_shop"
	"database/sql"
)

type Repository struct {
	ShopRepo r_shop.ShopRepoInterface
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		ShopRepo: r_shop.NewShopRepo(db),
	}
}
