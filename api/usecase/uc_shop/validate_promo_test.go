package uc_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type validatePromoCase struct {
	param        []models.Cart
	expectResult float64
	messageFail  string
}

func TestShopUC_ValidatePromo(t *testing.T) {
	shopUC := ShopUC{}
	testCase := []validatePromoCase{
		{
			param: []models.Cart{
				{
					ItemSKU:      MacbookProSKU,
					ItemPrice:    100,
					CartQuantity: 1,
				},
				{
					ItemSKU:      RaspberryPiSKU,
					ItemPrice:    50,
					CartQuantity: 1,
				},
			},
			expectResult: 100,
			messageFail:  "not pass buy 1 macbook free 1 raspberry",
		},
		{
			param: []models.Cart{
				{
					ItemSKU:      GoogleHomeSKU,
					ItemPrice:    100,
					CartQuantity: 3,
				},
			},
			expectResult: 200,
			messageFail:  "not pass buy 3 google home pay only 2 item",
		},
		{
			param: []models.Cart{
				{
					ItemSKU:      AlexaSpeakerSKU,
					ItemPrice:    100,
					CartQuantity: 3,
				},
			},
			expectResult: 270,
			messageFail:  "not pass buy 3 or more alexa speaker will discount all item 10 percent",
		},
	}

	for _, promoCase := range testCase {
		result := shopUC.ValidatePromo(promoCase.param)
		assert.Equal(t, promoCase.expectResult, result, promoCase.messageFail)
	}
}
