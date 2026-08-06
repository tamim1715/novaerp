package attendance

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CheckIn godoc
// @Summary Record employee check-in
// @Description Clock in for daily work shift
// @Tags Attendance
// @Accept json
// @Produce json
// @Param request body CheckInRequest true "Check-in payload"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /hr/attendances/check-in [post]
func _() {}

// CheckOut godoc
// @Summary Record employee check-out
// @Description Clock out for daily work shift and calculate hours
// @Tags Attendance
// @Accept json
// @Produce json
// @Param request body CheckOutRequest true "Check-out payload"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /hr/attendances/check-out [post]
func _() {}

// FindAll godoc
// @Summary List attendance logs
// @Description Retrieve daily attendance records with pagination
// @Tags Attendance
// @Produce json
// @Success 200 {object} response.APIResponse
// @Router /hr/attendances [get]
func _() {}
