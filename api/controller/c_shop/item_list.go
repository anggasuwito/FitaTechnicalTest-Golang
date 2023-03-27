package c_shop

import (
	"FitaTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c ShopController) ItemList(g *gin.Context) {
	res, err := c.UC.ShopUC.ItemListUC()
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
