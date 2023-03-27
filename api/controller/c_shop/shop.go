package c_shop

import (
	"FitaTechnicalTest-Golang/api/usecase"
	"github.com/gin-gonic/gin"
)

type ShopController struct {
	UC usecase.UseCase
}

type ShopControllerInterface interface {
	ItemList(g *gin.Context)
	CartList(g *gin.Context)
	CartInsert(g *gin.Context)
	Checkout(g *gin.Context)
}

func NewShopController(uc usecase.UseCase) ShopControllerInterface {
	return &ShopController{UC: uc}
}
