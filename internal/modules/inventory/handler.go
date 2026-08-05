package inventory

import (
	"errors"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/common/response"
	"github.com/tamim1715/novaerp/internal/common/validator"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// Warehouse Handlers
func (h *Handler) CreateWarehouse(c *gin.Context) {
	var req CreateWarehouseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate.Struct(req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	wh, err := h.service.CreateWarehouse(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(c, "Warehouse created successfully", ToWarehouseResponse(wh))
}

func (h *Handler) FindAllWarehouses(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	warehouses, total, err := h.service.FindAllWarehouses(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	req.Normalize()
	pageResp := pagination.PageResponse{
		Page:       req.Page,
		Size:       req.Size,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(req.Size))),
		Data:       ToWarehouseResponses(warehouses),
	}

	response.Success(c, "Warehouses fetched successfully", pageResp)
}

func (h *Handler) FindWarehouseByID(c *gin.Context) {
	id := c.Param("id")
	wh, err := h.service.FindWarehouseByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Warehouse not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Warehouse fetched successfully", ToWarehouseResponse(wh))
}

func (h *Handler) UpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	var req UpdateWarehouseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	wh, err := h.service.UpdateWarehouse(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Warehouse not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Warehouse updated successfully", ToWarehouseResponse(wh))
}

func (h *Handler) DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteWarehouse(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Warehouse not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Warehouse deleted successfully", nil)
}

// Product Handlers
func (h *Handler) CreateProduct(c *gin.Context) {
	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate.Struct(req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	prd, err := h.service.CreateProduct(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(c, "Product created successfully", ToProductResponse(prd))
}

func (h *Handler) FindAllProducts(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	products, total, err := h.service.FindAllProducts(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	req.Normalize()
	pageResp := pagination.PageResponse{
		Page:       req.Page,
		Size:       req.Size,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(req.Size))),
		Data:       ToProductResponses(products),
	}

	response.Success(c, "Products fetched successfully", pageResp)
}

func (h *Handler) FindProductByID(c *gin.Context) {
	id := c.Param("id")
	prd, err := h.service.FindProductByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Product not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Product fetched successfully", ToProductResponse(prd))
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	prd, err := h.service.UpdateProduct(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Product not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Product updated successfully", ToProductResponse(prd))
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteProduct(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Product not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Product deleted successfully", nil)
}

// Stock Handlers
func (h *Handler) FindAllStocks(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	stocks, total, err := h.service.FindAllStocks(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	req.Normalize()
	pageResp := pagination.PageResponse{
		Page:       req.Page,
		Size:       req.Size,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(req.Size))),
		Data:       ToStockResponses(stocks),
	}

	response.Success(c, "Stocks fetched successfully", pageResp)
}

func (h *Handler) AdjustStock(c *gin.Context) {
	var req AdjustStockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := validator.Validate.Struct(req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	stk, err := h.service.AdjustStock(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Stock adjusted successfully", ToStockResponse(stk))
}
