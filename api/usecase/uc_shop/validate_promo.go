package uc_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"math"
)

func (uc ShopUC) ValidatePromo(item []models.Cart) (finalPrice float64) {
	googleHomeSKU := "120P90"
	macbookProSKU := "43N23P"
	alexaSpeakerSKU := "A304SD"
	raspberryPiSKU := "234234"

	items := make(map[string]int)
	for _, cart := range item {
		items[cart.ItemSKU] += cart.CartQuantity
	}
	for _, cart := range item {
		switch cart.ItemSKU {
		case raspberryPiSKU:
			cart.CartQuantity -= items[macbookProSKU]
			finalPrice += cart.ItemPrice * float64(cart.CartQuantity)
		case googleHomeSKU:
			cart.CartQuantity -= items[googleHomeSKU] / 3
			finalPrice += cart.ItemPrice * float64(cart.CartQuantity)
		case alexaSpeakerSKU:
			if items[alexaSpeakerSKU] >= 3 {
				finalPrice += (cart.ItemPrice - cart.ItemPrice*10/100) * float64(cart.CartQuantity)
			} else {
				finalPrice += cart.ItemPrice * float64(cart.CartQuantity)
			}
		default:
			finalPrice += cart.ItemPrice * float64(cart.CartQuantity)
		}
	}
	finalPrice = math.Round(finalPrice*100) / 100
	return finalPrice
}
