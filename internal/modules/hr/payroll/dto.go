package payroll

type CreatePayrollPeriodRequest struct {
	Month int `json:"month" binding:"required,min=1,max=12" example:"8"`
	Year  int `json:"year" binding:"required,min=2020" example:"2026"`
}

type ProcessPayrollRequest struct {
	AllowancesDefault  float64 `json:"allowancesDefault" example:"0.00"`  // Fallback allowance if none configured
	DeductionsDefault  float64 `json:"deductionsDefault" example:"50.00"` // Standard tax/insurance deduction per employee
	ConsiderAttendance *bool   `json:"considerAttendance" example:"true"` // Calculate salary from attendance, overtime & leaves
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
	ID                   string  `json:"id"`
	PayrollPeriodID      string  `json:"payrollPeriodId"`
	EmployeeID           string  `json:"employeeId"`
	EmployeeName         string  `json:"employeeName,omitempty"`
	EmployeeDesignation  string  `json:"employeeDesignation,omitempty"`
	BasicSalary          float64 `json:"basicSalary"`
	Allowances           float64 `json:"allowances"`
	OvertimeHours        float64 `json:"overtimeHours"`
	OvertimePay          float64 `json:"overtimePay"`
	Deductions           float64 `json:"deductions"`
	AbsentDays           int     `json:"absentDays"`
	AbsentDeduction      float64 `json:"absentDeduction"`
	UnpaidLeaveDays      int     `json:"unpaidLeaveDays"`
	UnpaidLeaveDeduction float64 `json:"unpaidLeaveDeduction"`
	PresentDays          int     `json:"presentDays"`
	WorkingDays          int     `json:"workingDays"`
	GrossSalary          float64 `json:"grossSalary"`
	NetSalary            float64 `json:"netSalary"`
	Status               string  `json:"status"`
	PaymentDate          string  `json:"paymentDate,omitempty"`
	Notes                string  `json:"notes,omitempty"`
	CreatedAt            int64   `json:"createdAt"`
}
