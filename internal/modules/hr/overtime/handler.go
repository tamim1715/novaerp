package overtime

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

func (h *Handler) CreatePolicy(c *gin.Context) {
	var req CreateOvertimePolicyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	policy, err := h.service.CreatePolicy(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Created(c, "Overtime policy created successfully", ToOvertimePolicyResponse(policy))
}

func (h *Handler) FindAllPolicies(c *gin.Context) {
	var req pagination.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	list, total, err := h.service.FindAllPolicies(c.Request.Context(), req)
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
		Data:       ToOvertimePolicyResponseList(list),
	}

	response.Success(c, "Overtime policies fetched successfully", pageResp)
}

func (h *Handler) FindPolicyByID(c *gin.Context) {
	id := c.Param("id")
	policy, err := h.service.FindPolicyByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, "Overtime policy not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Overtime policy fetched successfully", ToOvertimePolicyResponse(policy))
}

func (h *Handler) UpdatePolicy(c *gin.Context) {
	id := c.Param("id")
	var req UpdateOvertimePolicyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	policy, err := h.service.UpdatePolicy(c.Request.Context(), id, req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "Overtime policy updated successfully", ToOvertimePolicyResponse(policy))
}

func (h *Handler) DeletePolicy(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeletePolicy(c.Request.Context(), id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, "Overtime policy deleted successfully", nil)
}
