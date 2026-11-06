package overtime

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreatePolicy godoc
// @Summary Create an overtime policy
// @Description Configure overtime thresholds, daily maximums, and rate multipliers
// @Tags Overtime
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateOvertimePolicyRequest true "Overtime policy payload"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /hr/overtimes [post]
func _() {}

// FindAllPolicies godoc
// @Summary List overtime policies
// @Description Retrieve overtime policies with pagination
// @Tags Overtime
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.APIResponse
// @Router /hr/overtimes [get]
func _() {}
