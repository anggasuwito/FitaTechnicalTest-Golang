package c_shop

import (
	"FitaTechnicalTest-Golang/api/models"
	"FitaTechnicalTest-Golang/helper/response"
	"github.com/gin-gonic/gin"
)

func (c ShopController) CartInsert(g *gin.Context) {
	var req models.CartInsertRequest
	g.ShouldBindJSON(&req)
	res, err := c.UC.ShopUC.CartInsertUC(req)
	response.JSON(response.Response{
		Ctx:   g,
		Data:  res,
		Meta:  nil,
		Error: err,
	})
}
