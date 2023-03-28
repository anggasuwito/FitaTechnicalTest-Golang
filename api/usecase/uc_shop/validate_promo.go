package uc_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"math"
)

const (
	GoogleHomeSKU   = "120P90"
	MacbookProSKU   = "43N23P"
	AlexaSpeakerSKU = "A304SD"
	RaspberryPiSKU  = "234234"
)

func (uc ShopUC) ValidatePromo(item []models.Cart) (finalPrice float64) {
	items := make(map[string]int)
	for _, cart := range item {
		items[cart.ItemSKU] += cart.CartQuantity
	}
	for _, cart := range item {
		switch cart.ItemSKU {
		case RaspberryPiSKU:
			cart.CartQuantity -= items[MacbookProSKU]
			finalPrice += cart.ItemPrice * float64(cart.CartQuantity)
		case GoogleHomeSKU:
			cart.CartQuantity -= items[GoogleHomeSKU] / 3
			finalPrice += cart.ItemPrice * float64(cart.CartQuantity)
		case AlexaSpeakerSKU:
			if items[AlexaSpeakerSKU] >= 3 {
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
