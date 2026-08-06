package leaverequest

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateLeaveRequest godoc
// @Summary Apply for leave
// @Description Submit a leave application for an employee
// @Tags Leave Requests
// @Accept json
// @Produce json
// @Param request body CreateLeaveRequest true "Leave application details"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /hr/leaves/requests [post]
func _() {}

// UpdateLeaveStatus godoc
// @Summary Approve or reject leave request
// @Description Update leave status to APPROVED or REJECTED
// @Tags Leave Requests
// @Accept json
// @Produce json
// @Param id path string true "Leave Request ID"
// @Param request body UpdateLeaveStatusRequest true "Approval details"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /hr/leaves/requests/{id}/status [put]
func _() {}
