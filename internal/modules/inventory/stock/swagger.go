package stock

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// FindAllStocks godoc
// @Summary List all stock entries
// @Description Retrieve inventory stock levels across warehouses
// @Tags Stocks
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Items per page (default 10)"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/stocks [get]
func _() {}

// AdjustStock godoc
// @Summary Adjust product stock in warehouse
// @Description Increase or decrease quantity of a product in a warehouse
// @Tags Stocks
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body AdjustStockRequest true "Stock adjustment details"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/stocks/adjust [post]
func _() {}
