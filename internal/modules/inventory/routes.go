package inventory

import "github.com/gin-gonic/gin"

func RegisterRoutes(api *gin.RouterGroup, handler *Handler) {
	// Warehouse routes
	warehouses := api.Group("/warehouses")
	{
		warehouses.POST("", handler.CreateWarehouse)
		warehouses.GET("", handler.FindAllWarehouses)
		warehouses.GET("/:id", handler.FindWarehouseByID)
		warehouses.PUT("/:id", handler.UpdateWarehouse)
		warehouses.DELETE("/:id", handler.DeleteWarehouse)
	}

	// Product routes
	products := api.Group("/products")
	{
		products.POST("", handler.CreateProduct)
		products.GET("", handler.FindAllProducts)
		products.GET("/:id", handler.FindProductByID)
		products.PUT("/:id", handler.UpdateProduct)
		products.DELETE("/:id", handler.DeleteProduct)
	}

	// Stock routes
	stocks := api.Group("/stocks")
	{
		stocks.GET("", handler.FindAllStocks)
		stocks.POST("/adjust", handler.AdjustStock)
	}
}
