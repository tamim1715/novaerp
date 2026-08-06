package payroll

type CreatePayrollPeriodRequest struct {
	Month int `json:"month" binding:"required,min=1,max=12" example:"8"`
	Year  int `json:"year" binding:"required,min=2020" example:"2026"`
}

type ProcessPayrollRequest struct {
	AllowancesDefault float64 `json:"allowancesDefault" example:"200.00"` // Standard bonus/allowance per employee
	DeductionsDefault float64 `json:"deductionsDefault" example:"50.00"`  // Standard tax/insurance deduction per employee
}

type PayrollPeriodResponse struct {
	ID          string  `json:"id"`
	Month       int     `json:"month"`
	Year        int     `json:"year"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
	Status      string  `json:"status"`
	TotalGross  float64 `json:"totalGross"`
	TotalNet    float64 `json:"totalNet"`
	ProcessedAt string  `json:"processedAt,omitempty"`
	CreatedAt   int64   `json:"createdAt"`
}

type PayslipResponse struct {
	ID                  string  `json:"id"`
	PayrollPeriodID     string  `json:"payrollPeriodId"`
	EmployeeID          string  `json:"employeeId"`
	EmployeeName        string  `json:"employeeName,omitempty"`
	EmployeeDesignation string  `json:"employeeDesignation,omitempty"`
	BasicSalary         float64 `json:"basicSalary"`
	Allowances          float64 `json:"allowances"`
	Deductions          float64 `json:"deductions"`
	UnpaidLeaveDeduction float64 `json:"unpaidLeaveDeduction"`
	GrossSalary         float64 `json:"grossSalary"`
	NetSalary           float64 `json:"netSalary"`
	Status              string  `json:"status"`
	PaymentDate         string  `json:"paymentDate,omitempty"`
	Notes               string  `json:"notes,omitempty"`
	CreatedAt           int64   `json:"createdAt"`
}
