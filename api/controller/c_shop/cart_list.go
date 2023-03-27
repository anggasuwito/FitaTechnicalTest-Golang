package c_shop

import (
	"FitaTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c ShopController) CartList(g *gin.Context) {
	userID := g.Param("userID")
	res, err := c.UC.ShopUC.CartListUC(userID)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
