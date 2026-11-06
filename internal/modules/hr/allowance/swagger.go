package allowance

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateType godoc
// @Summary Create an allowance type
// @Description Configure an allowance type (e.g. HRA, Medical) with fixed amount or percentage of basic
// @Tags Allowance
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateAllowanceTypeRequest true "Allowance type payload"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /hr/allowances/types [post]
func _() {}

// FindAllTypes godoc
// @Summary List allowance types
// @Description Retrieve allowance types with pagination
// @Tags Allowance
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.APIResponse
// @Router /hr/allowances/types [get]
func _() {}

// AssignAllowance godoc
// @Summary Assign allowance to employee
// @Description Allocate fixed or percentage allowance to an employee
// @Tags Allowance
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body AssignAllowanceRequest true "Allowance assignment payload"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /hr/allowances/assign [post]
func _() {}
