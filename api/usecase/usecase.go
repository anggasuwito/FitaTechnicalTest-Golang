package usecase

import (
	"FitaTechnicalTest-Golang/api/repository"
	"FitaTechnicalTest-Golang/api/usecase/uc_shop"
	"database/sql"
)

type UseCase struct {
	ShopUC uc_shop.ShopUCInterface
}

func NewUseCase(db *sql.DB) UseCase {
	repo := repository.NewRepository(db)
	return UseCase{
		ShopUC: uc_shop.NewShopUC(repo),
	}
}
