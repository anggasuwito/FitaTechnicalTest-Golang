package uc_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"errors"
)

func (uc ShopUC) CartInsertUC(req models.CartInsertRequest) (res string, err error) {
	if req.Quantity == 0 {
		return res, errors.New("please insert quantity")
	} else if req.UserID == "" {
		return res, errors.New("please insert your user id")
	} else if req.ItemSKU == "" {
		return res, errors.New("please insert item SKU")
	}

	item, err := uc.Repo.ShopRepo.GetItemBySKU(req.ItemSKU)
	if err != nil {
		return res, err
	}

	if req.Quantity > item.Quantity {
		return res, errors.New("item current quantity is not enough")
	}

	cart, err := uc.CartListUC(req.UserID)
	if err != nil {
		return res, err
	}
	var currentItem models.Cart
	for _, c := range cart.Items {
		if req.ItemSKU == c.ItemSKU {
			currentItem = c
			break
		}
	}

	if currentItem.CartID == "" {
		err = uc.Repo.ShopRepo.AddtoUserCart(req)
		if err != nil {
			return res, err
		}
	} else {
		err = uc.Repo.ShopRepo.UpdateUserCart(currentItem.CartID, req)
		if err != nil {
			return res, err
		}
	}

	res = "Success add to your cart"
	return res, err
}
