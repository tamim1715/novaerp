package employee

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateDoc Create godoc
//
// @Summary Create Employee
// @Description Create a new employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param request body CreateEmployeeRequest true "Employee"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /employees [post]
func (h *Handler) CreateDoc(c *gin.Context) {
	h.Create(c)
}

// FindAllDoc FindAll godoc
//
// @Summary Get Employees
// @Description Get all employees
// @Tags Employee
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param search query string false "Search"
// @Param sortBy query string false "Sort By"
// @Param order query string false "asc|desc"
// @Success 200 {object} response.APIResponse
// @Router /employees [get]
func (h *Handler) FindAllDoc(c *gin.Context) {
	h.FindAll(c)
}

// FindByIDDoc FindByID godoc
//
// @Summary Get Employee
// @Description Get employee by ID
// @Tags Employee
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /employees/{id} [get]
func (h *Handler) FindByIDDoc(c *gin.Context) {
	h.FindByID(c)
}

// UpdateDoc Update godoc
//
// @Summary Update Employee
// @Description Update employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param request body UpdateEmployeeRequest true "Employee"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /employees/{id} [put]
func (h *Handler) UpdateDoc(c *gin.Context) {
	h.Update(c)
}

// DeleteDoc Delete godoc
//
// @Summary Delete Employee
// @Description Soft delete employee
// @Tags Employee
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /employees/{id} [delete]
func (h *Handler) DeleteDoc(c *gin.Context) {
	h.Delete(c)
}
