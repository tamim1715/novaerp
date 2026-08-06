package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateDoc godoc
// @Summary Create User
// @Description Create a new user account
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateUserRequest true "User"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /users [post]
func (h *Handler) CreateDoc(c *gin.Context) {
	h.Create(c)
}

// FindAllDoc godoc
// @Summary Get Users
// @Description Get paginated list of users
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param search query string false "Search"
// @Param sortBy query string false "Sort By"
// @Param order query string false "asc|desc"
// @Success 200 {object} response.APIResponse
// @Router /users [get]
func (h *Handler) FindAllDoc(c *gin.Context) {
	h.FindAll(c)
}

// FindByIDDoc godoc
// @Summary Get User
// @Description Get user by ID
// @Tags User
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /users/{id} [get]
func (h *Handler) FindByIDDoc(c *gin.Context) {
	h.FindByID(c)
}

// UpdateDoc godoc
// @Summary Update User
// @Description Update user details
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param request body UpdateUserRequest true "User"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /users/{id} [put]
func (h *Handler) UpdateDoc(c *gin.Context) {
	h.Update(c)
}

// DeleteDoc godoc
// @Summary Delete User
// @Description Soft delete user account
// @Tags User
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /users/{id} [delete]
func _(h *Handler, c *gin.Context) {
	h.Delete(c)
}
