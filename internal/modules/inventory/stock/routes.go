package stock

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.GET("", handler.FindAllStocks)
	router.POST("/adjust", handler.AdjustStock)
}
