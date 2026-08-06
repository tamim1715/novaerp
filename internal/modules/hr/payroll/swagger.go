package payroll

import (
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreatePeriod godoc
// @Summary Create a payroll period
// @Description Initialize a monthly payroll period batch
// @Tags Payroll
// @Accept json
// @Produce json
// @Param request body CreatePayrollPeriodRequest true "Payroll period month/year"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /hr/payrolls [post]
func _() {}

// ProcessPayroll godoc
// @Summary Process monthly payroll
// @Description Calculate salary slips for all active employees for the period
// @Tags Payroll
// @Accept json
// @Produce json
// @Param id path string true "Payroll Period ID"
// @Param request body ProcessPayrollRequest true "Default allowances and deductions"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Router /hr/payrolls/{id}/process [post]
func _() {}

// GetPayslips godoc
// @Summary List payslips for period
// @Description Retrieve generated payslips for a payroll period
// @Tags Payroll
// @Produce json
// @Param id path string true "Payroll Period ID"
// @Success 200 {object} response.APIResponse
// @Router /hr/payrolls/{id}/payslips [get]
func _() {}
