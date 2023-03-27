package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx   *gin.Context
	Data  interface{}
	Meta  interface{}
	Error error
}

func JSON(res Response) {
	if res.Error != nil {
		res.Ctx.JSON(http.StatusBadRequest, gin.H{
			"status":       false,
			"message":      "error",
			"data":         nil,
			"meta":         nil,
			"err_response": res.Error.Error(),
		})
	} else {
		res.Ctx.JSON(http.StatusOK, gin.H{
			"status":       true,
			"message":      "success",
			"data":         res.Data,
			"meta":         res.Meta,
			"err_response": nil,
		})
	}
}

func JSONUnauthorized(res Response) {
	res.Ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"status":       false,
		"message":      "unauthenticate",
		"data":         nil,
		"meta":         nil,
		"err_response": res.Error.Error(),
	})
}
