package c_shop

import (
	"FitaTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c ShopController) Checkout(g *gin.Context) {
	userID := g.Param("userID")
	res, err := c.UC.ShopUC.CheckoutUC(userID)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
