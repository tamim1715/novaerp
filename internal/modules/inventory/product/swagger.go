package product

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateProduct godoc
// @Summary Create a product
// @Description Add a new product to inventory
// @Tags Products
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "Product creation details"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/products [post]
func _() {}

// FindAllProducts godoc
// @Summary List all products
// @Description Retrieve products with pagination, sorting, and search
// @Tags Products
// @Produce json
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Items per page (default 10)"
// @Param search query string false "Search query"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/products [get]
func _() {}

// FindProductByID godoc
// @Summary Get product by ID
// @Description Retrieve details of a specific product
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/products/{id} [get]
func _() {}

// UpdateProduct godoc
// @Summary Update product
// @Description Update product details by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body UpdateProductRequest true "Product update fields"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/products/{id} [put]
func _() {}

// DeleteProduct godoc
// @Summary Delete product
// @Description Remove a product by ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /inventories/products/{id} [delete]
func _() {}
