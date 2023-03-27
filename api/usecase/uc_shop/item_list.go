package uc_shop

import "FitaTechnicalTest-Golang/api/models"

func (uc ShopUC) ItemListUC() (res []models.Item, err error) {
	res, err = uc.Repo.ShopRepo.GetAllItem()
	return res, err
}
