package inventory

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// Warehouse Swagger Docs
// @Summary Create Warehouse
// @Description Create a new warehouse location
// @Tags Inventory - Warehouses
// @Accept json
// @Produce json
// @Param request body CreateWarehouseRequest true "Warehouse"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /warehouses [post]
func (h *Handler) CreateWarehouseDoc(c *gin.Context) {
	h.CreateWarehouse(c)
}

// @Summary Get Warehouses
// @Description Get paginated list of warehouses
// @Tags Inventory - Warehouses
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param search query string false "Search"
// @Param sortBy query string false "Sort By"
// @Param order query string false "asc|desc"
// @Success 200 {object} response.APIResponse
// @Router /warehouses [get]
func (h *Handler) FindAllWarehousesDoc(c *gin.Context) {
	h.FindAllWarehouses(c)
}

// @Summary Get Warehouse
// @Description Get warehouse by ID
// @Tags Inventory - Warehouses
// @Produce json
// @Param id path string true "Warehouse ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /warehouses/{id} [get]
func (h *Handler) FindWarehouseByIDDoc(c *gin.Context) {
	h.FindWarehouseByID(c)
}

// @Summary Update Warehouse
// @Description Update warehouse details
// @Tags Inventory - Warehouses
// @Accept json
// @Produce json
// @Param id path string true "Warehouse ID"
// @Param request body UpdateWarehouseRequest true "Warehouse"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /warehouses/{id} [put]
func (h *Handler) UpdateWarehouseDoc(c *gin.Context) {
	h.UpdateWarehouse(c)
}

// @Summary Delete Warehouse
// @Description Delete warehouse
// @Tags Inventory - Warehouses
// @Produce json
// @Param id path string true "Warehouse ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /warehouses/{id} [delete]
func (h *Handler) DeleteWarehouseDoc(c *gin.Context) {
	h.DeleteWarehouse(c)
}

// Product Swagger Docs
// @Summary Create Product
// @Description Create a new product item in catalog
// @Tags Inventory - Products
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "Product"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /products [post]
func (h *Handler) CreateProductDoc(c *gin.Context) {
	h.CreateProduct(c)
}

// @Summary Get Products
// @Description Get paginated product catalog
// @Tags Inventory - Products
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param search query string false "Search"
// @Param sortBy query string false "Sort By"
// @Param order query string false "asc|desc"
// @Success 200 {object} response.APIResponse
// @Router /products [get]
func (h *Handler) FindAllProductsDoc(c *gin.Context) {
	h.FindAllProducts(c)
}

// @Summary Get Product
// @Description Get product details by ID
// @Tags Inventory - Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /products/{id} [get]
func (h *Handler) FindProductByIDDoc(c *gin.Context) {
	h.FindProductByID(c)
}

// @Summary Update Product
// @Description Update product details
// @Tags Inventory - Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body UpdateProductRequest true "Product"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /products/{id} [put]
func (h *Handler) UpdateProductDoc(c *gin.Context) {
	h.UpdateProduct(c)
}

// @Summary Delete Product
// @Description Delete product item
// @Tags Inventory - Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /products/{id} [delete]
func (h *Handler) DeleteProductDoc(c *gin.Context) {
	h.DeleteProduct(c)
}

// Stock Swagger Docs
// @Summary Get Stocks
// @Description Get stock balances across warehouses
// @Tags Inventory - Stocks
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Success 200 {object} response.APIResponse
// @Router /stocks [get]
func (h *Handler) FindAllStocksDoc(c *gin.Context) {
	h.FindAllStocks(c)
}

// @Summary Adjust Stock
// @Description Adjust quantity of a product in a warehouse
// @Tags Inventory - Stocks
// @Accept json
// @Produce json
// @Param request body AdjustStockRequest true "Stock Adjustment"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /stocks/adjust [post]
func (h *Handler) AdjustStockDoc(c *gin.Context) {
	h.AdjustStock(c)
}
