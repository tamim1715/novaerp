package warehouse

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateWarehouse godoc
// @Summary Create a warehouse
// @Description Add a new warehouse to inventory
// @Tags Warehouses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateWarehouseRequest true "Warehouse creation details"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/warehouses [post]
func _() {}

// FindAllWarehouses godoc
// @Summary List all warehouses
// @Description Retrieve warehouses with pagination, sorting, and search
// @Tags Warehouses
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Items per page (default 10)"
// @Param search query string false "Search query"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/warehouses [get]
func _() {}

// FindWarehouseByID godoc
// @Summary Get warehouse by ID
// @Description Retrieve details of a specific warehouse
// @Tags Warehouses
// @Produce json
// @Security BearerAuth
// @Param id path string true "Warehouse ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/warehouses/{id} [get]
func _() {}

// UpdateWarehouse godoc
// @Summary Update warehouse
// @Description Update warehouse details by ID
// @Tags Warehouses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Warehouse ID"
// @Param request body UpdateWarehouseRequest true "Warehouse update fields"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/warehouses/{id} [put]
func _() {}

// DeleteWarehouse godoc
// @Summary Delete warehouse
// @Description Remove a warehouse by ID
// @Tags Warehouses
// @Produce json
// @Security BearerAuth
// @Param id path string true "Warehouse ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/warehouses/{id} [delete]
func _() {}
