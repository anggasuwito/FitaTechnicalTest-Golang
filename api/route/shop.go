package route

import (
	"FitaTechnicalTest-Golang/api/controller"
	"github.com/gin-gonic/gin"
)

func Shop(r *gin.RouterGroup, c controller.Controller) {
	shopAPI := r.Group("/shop")
	shopAPI.GET("/item-list", c.ShopController.ItemList)
	shopAPI.GET("/cart/:userID", c.ShopController.CartList)
	shopAPI.POST("/cart", c.ShopController.CartInsert)
	shopAPI.POST("/checkout/:userID", c.ShopController.Checkout)
}
