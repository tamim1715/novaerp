package product

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("", handler.CreateProduct)
	router.GET("", handler.FindAllProducts)
	router.GET("/:id", handler.FindProductByID)
	router.PUT("/:id", handler.UpdateProduct)
	router.DELETE("/:id", handler.DeleteProduct)
}
