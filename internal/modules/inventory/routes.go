package inventory

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/inventory/product"
	"github.com/tamim1715/novaerp/internal/modules/inventory/stock"
	"github.com/tamim1715/novaerp/internal/modules/inventory/warehouse"
)

// RegisterRoutes registers all inventory submodules under /inventories (and /inventorys for compatibility).
func RegisterRoutes(api *gin.RouterGroup, application *app.Application) {
	// Initialize Submodule Repositories
	warehouseRepo := warehouse.NewRepository(application.DB)
	productRepo := product.NewRepository(application.DB)
	stockRepo := stock.NewRepository(application.DB)

	// Initialize Submodule Services
	warehouseService := warehouse.NewService(warehouseRepo, application.Logger)
	productService := product.NewService(productRepo, application.Logger)
	stockService := stock.NewService(stockRepo, warehouseRepo, productRepo, application.Logger)

	// Initialize Submodule Handlers
	warehouseHandler := warehouse.NewHandler(warehouseService)
	productHandler := product.NewHandler(productService)
	stockHandler := stock.NewHandler(stockService)

	// Primary REST endpoint: /api/v1/inventories/...
	inventoriesGroup := api.Group("/inventories")
	{
		warehouse.RegisterRoutes(inventoriesGroup.Group("/warehouses"), warehouseHandler)
		product.RegisterRoutes(inventoriesGroup.Group("/products"), productHandler)
		stock.RegisterRoutes(inventoriesGroup.Group("/stocks"), stockHandler)
	}

	// Alias endpoint: /api/v1/inventorys/...
	inventorysGroup := api.Group("/inventorys")
	{
		warehouse.RegisterRoutes(inventorysGroup.Group("/warehouses"), warehouseHandler)
		product.RegisterRoutes(inventorysGroup.Group("/products"), productHandler)
		stock.RegisterRoutes(inventorysGroup.Group("/stocks"), stockHandler)
	}
}
