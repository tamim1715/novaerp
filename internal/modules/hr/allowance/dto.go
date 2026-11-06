package allowance

type CreateAllowanceTypeRequest struct {
	Name            string  `json:"name" binding:"required" example:"House Rent Allowance"`
	Code            string  `json:"code" binding:"required" example:"HRA"`
	CalculationType string  `json:"calculationType" binding:"required" example:"PERCENTAGE"` // FIXED or PERCENTAGE
	DefaultAmount   float64 `json:"defaultAmount" example:"40.0"`                            // 40% if PERCENTAGE, or flat amount if FIXED
	IsTaxable       bool    `json:"isTaxable" example:"false"`
	Description     string  `json:"description" example:"Monthly housing allowance"`
}

type UpdateAllowanceTypeRequest struct {
	Name            string   `json:"name" example:"House Rent Allowance"`
	CalculationType string   `json:"calculationType" example:"PERCENTAGE"`
	DefaultAmount   *float64 `json:"defaultAmount" example:"45.0"`
	IsTaxable       *bool    `json:"isTaxable" example:"false"`
	IsActive        *bool    `json:"isActive" example:"true"`
	Description     string   `json:"description" example:"Updated housing allowance"`
}

type AllowanceTypeResponse struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Code            string  `json:"code"`
	CalculationType string  `json:"calculationType"`
	DefaultAmount   float64 `json:"defaultAmount"`
	IsTaxable       bool    `json:"isTaxable"`
	IsActive        bool    `json:"isActive"`
	Description     string  `json:"description"`
	CreatedAt       int64   `json:"createdAt"`
}

type AssignAllowanceRequest struct {
	EmployeeID      string  `json:"employeeId" binding:"required" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`
	AllowanceTypeID string  `json:"allowanceTypeId" binding:"required" example:"5e53303c-8ab5-46eb-a1f9-9ea2eb60b37f"`
	Amount          float64 `json:"amount" example:"500.00"` // Specific override amount; if 0, type default is used
}

type EmployeeAllowanceResponse struct {
	ID              string                 `json:"id"`
	EmployeeID      string                 `json:"employeeId"`
	EmployeeName    string                 `json:"employeeName,omitempty"`
	AllowanceTypeID string                 `json:"allowanceTypeId"`
	AllowanceType   *AllowanceTypeResponse `json:"allowanceType,omitempty"`
	Amount          float64                `json:"amount"`
	IsActive        bool                   `json:"isActive"`
	CreatedAt       int64                  `json:"createdAt"`
}

type CalculatedAllowanceItem struct {
	AllowanceTypeID string  `json:"allowanceTypeId"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	CalculationType string  `json:"calculationType"`
	RateOrAmount    float64 `json:"rateOrAmount"`
	CalculatedTotal float64 `json:"calculatedTotal"`
}
