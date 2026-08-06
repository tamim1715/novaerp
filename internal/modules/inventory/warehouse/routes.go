package warehouse

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("", handler.CreateWarehouse)
	router.GET("", handler.FindAllWarehouses)
	router.GET("/:id", handler.FindWarehouseByID)
	router.PUT("/:id", handler.UpdateWarehouse)
	router.DELETE("/:id", handler.DeleteWarehouse)
}
