package account

import (
	"errors"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/common/response"
	"gorm.io/gorm"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var req CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	acc, err := h.service.CreateAccount(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(c, "Account created successfully", ToAccountResponse(acc))
}

func (h *Handler) FindAllAccounts(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	accountType := c.Query("type")

	accounts, total, err := h.service.FindAllAccounts(c.Request.Context(), req, accountType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	req.Normalize()
	pageResponse := pagination.PageResponse{
		Page:       req.Page,
		Size:       req.Size,
		TotalItems: total,
		TotalPages: int(math.Ceil(float64(total) / float64(req.Size))),
		Data:       ToAccountResponseList(accounts),
	}

	response.Success(c, "Accounts fetched successfully", pageResponse)
}

func (h *Handler) GetAccountTree(c *gin.Context) {
	roots, err := h.service.GetAccountTree(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	tree := make([]AccountTreeResponse, len(roots))
	for i, r := range roots {
		tree[i] = ToAccountTreeResponse(&r)
	}

	response.Success(c, "Account tree hierarchy fetched successfully", tree)
}

func (h *Handler) FindAccountByID(c *gin.Context) {
	id := c.Param("id")
	acc, err := h.service.FindAccountByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Account not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Account fetched successfully", ToAccountResponse(acc))
}

func (h *Handler) UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	var req UpdateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	acc, err := h.service.UpdateAccount(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Account not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Account updated successfully", ToAccountResponse(acc))
}

func (h *Handler) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteAccount(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Account not found")
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Account deleted successfully", nil)
}

func (h *Handler) SeedAccounts(c *gin.Context) {
	if err := h.service.SeedChartOfAccounts(c.Request.Context()); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Standard Chart of Accounts seeded successfully", nil)
}
