package department

import (
	"github.com/gin-gonic/gin"

	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateDoc Create godoc
// @Summary Create Department
// @Description Create a new department
// @Tags Department
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateDepartmentRequest true "Department"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /departments [post]
func (h *Handler) CreateDoc(c *gin.Context) {
	h.Create(c)
}

// FindAllDoc FindAll godoc
// @Summary Get Departments
// @Description Get all departments
// @Tags Department
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param search query string false "Search"
// @Param sortBy query string false "Sort By"
// @Param order query string false "asc|desc"
// @Success 200 {object} response.APIResponse
// @Router /departments [get]
func (h *Handler) FindAllDoc(c *gin.Context) {
	h.FindAll(c)
}

// FindByIDDoc FindByID godoc
// @Summary Get Department
// @Description Get department by ID
// @Tags Department
// @Produce json
// @Security BearerAuth
// @Param id path string true "Department ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /departments/{id} [get]
func (h *Handler) FindByIDDoc(c *gin.Context) {
	h.FindByID(c)
}

// UpdateDoc Update godoc
// @Summary Update Department
// @Description Update department
// @Tags Department
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Department ID"
// @Param request body UpdateDepartmentRequest true "Department"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /departments/{id} [put]
func (h *Handler) UpdateDoc(c *gin.Context) {
	h.Update(c)
}

// DeleteDoc Delete godoc
// @Summary Delete Department
// @Description Soft delete department
// @Tags Department
// @Produce json
// @Security BearerAuth
// @Param id path string true "Department ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /departments/{id} [delete]
func (h *Handler) DeleteDoc(c *gin.Context) {
	h.Delete(c)
}
