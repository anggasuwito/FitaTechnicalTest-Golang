package uc_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func (uc ShopUC) CheckoutUC(userID string) (res models.CartListResponse, err error) {
	if userID == "" {
		return res, errors.New("please insert your user id")
	}

	cart, err := uc.Repo.ShopRepo.GetUserCart(userID)
	if err != nil {
		return res, err
	} else if len(cart) == 0 {
		return res, errors.New("there is no item for checkout")
	}

	tx, err := uc.Repo.ShopRepo.TXBegin()
	if err != nil {
		return res, err
	}

	transaction := models.Transaction{
		ID:         uuid.New().String(),
		UserID:     userID,
		TotalPrice: uc.ValidatePromo(cart),
		Status:     "success",
	}
	if transaction.ID == "" {
		return res, errors.New("error generate transaction id")
	}

	err = uc.Repo.ShopRepo.TXAddTransaction(tx, transaction)
	if err != nil {
		tx.Rollback()
		return res, err
	}

	for _, c := range cart {
		item, err := uc.Repo.ShopRepo.TXGetItemForUpdate(tx, c.ItemSKU)
		if err != nil {
			tx.Rollback()
			return res, err
		} else if c.CartQuantity > item.Quantity {
			tx.Rollback()
			return res, fmt.Errorf("item %v current quantity is not enough with your requested checkout quantity", item.Name)
		}

		err = uc.Repo.ShopRepo.TXUpdateItem(tx, c)
		if err != nil {
			tx.Rollback()
			return res, err
		}

		err = uc.Repo.ShopRepo.TXAddTransactionHistory(tx, c, transaction.ID)
		if err != nil {
			tx.Rollback()
			return res, err
		}

		err = uc.Repo.ShopRepo.TXDeleteExistingCart(tx, userID)
		if err != nil {
			tx.Rollback()
			return res, err
		}
	}

	tx.Commit()

	res.Items = cart
	res.CheckoutTotalPrice = transaction.TotalPrice
	return res, err
}
