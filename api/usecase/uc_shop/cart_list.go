package uc_shop

import "FitaTechnicalTest-Golang/api/models"

func (uc ShopUC) CartListUC(userID string) (res models.CartListResponse, err error) {
	cart, err := uc.Repo.ShopRepo.GetUserCart(userID)
	res.CheckoutTotalPrice = uc.ValidatePromo(cart)
	res.Items = cart
	return res, err
}
