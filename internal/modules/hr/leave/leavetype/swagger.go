package leavetype

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateLeaveType godoc
// @Summary Create a leave type
// @Description Define a new leave policy category
// @Tags Leave Types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateLeaveTypeRequest true "Leave type details"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /hr/leaves/types [post]
func _() {}

// FindAllLeaveTypes godoc
// @Summary List leave types
// @Description Retrieve leave policy categories
// @Tags Leave Types
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.APIResponse
// @Router /hr/leaves/types [get]
func _() {}
