package shift

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateShift godoc
// @Summary Create a work shift configuration
// @Description Configure a work shift with start/end time, grace period, and work hours
// @Tags Shift
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateShiftRequest true "Shift configuration payload"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /hr/shifts [post]
func _() {}

// FindAllShifts godoc
// @Summary List work shifts
// @Description Retrieve work shifts with pagination
// @Tags Shift
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.APIResponse
// @Router /hr/shifts [get]
func _() {}

// AssignShift godoc
// @Summary Assign a shift to an employee
// @Description Assign work shift with effective dates to an employee
// @Tags Shift
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body AssignShiftRequest true "Shift assignment payload"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /hr/shifts/assign [post]
func _() {}
