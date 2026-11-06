package shift

type CreateShiftRequest struct {
	Name               string  `json:"name" binding:"required" example:"General Morning Shift"`
	Code               string  `json:"code" binding:"required" example:"GEN_MORN"`
	StartTime          string  `json:"startTime" binding:"required" example:"09:00"`
	EndTime            string  `json:"endTime" binding:"required" example:"17:00"`
	GracePeriodMinutes int     `json:"gracePeriodMinutes" example:"15"`
	WorkHours          float64 `json:"workHours" example:"8.0"`
	HalfDayHours       float64 `json:"halfDayHours" example:"4.0"`
	IsNightShift       bool    `json:"isNightShift" example:"false"`
	Description        string  `json:"description" example:"Default corporate morning shift"`
}

type UpdateShiftRequest struct {
	Name               string  `json:"name" example:"General Morning Shift"`
	StartTime          string  `json:"startTime" example:"09:00"`
	EndTime            string  `json:"endTime" example:"17:00"`
	GracePeriodMinutes int     `json:"gracePeriodMinutes" example:"15"`
	WorkHours          float64 `json:"workHours" example:"8.0"`
	HalfDayHours       float64 `json:"halfDayHours" example:"4.0"`
	IsNightShift       bool    `json:"isNightShift" example:"false"`
	IsActive           *bool   `json:"isActive" example:"true"`
	Description        string  `json:"description" example:"Updated corporate morning shift"`
}

type ShiftResponse struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	Code               string  `json:"code"`
	StartTime          string  `json:"startTime"`
	EndTime            string  `json:"endTime"`
	GracePeriodMinutes int     `json:"gracePeriodMinutes"`
	WorkHours          float64 `json:"workHours"`
	HalfDayHours       float64 `json:"halfDayHours"`
	IsNightShift       bool    `json:"isNightShift"`
	IsActive           bool    `json:"isActive"`
	Description        string  `json:"description"`
	CreatedAt          int64   `json:"createdAt"`
}

type AssignShiftRequest struct {
	EmployeeID    string `json:"employeeId" binding:"required" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`
	ShiftID       string `json:"shiftId" binding:"required" example:"5e53303c-8ab5-46eb-a1f9-9ea2eb60b37f"`
	EffectiveFrom string `json:"effectiveFrom" binding:"required" example:"2026-08-01"`
	EffectiveTo   string `json:"effectiveTo" example:"2026-12-31"`
}

type EmployeeShiftResponse struct {
	ID            string         `json:"id"`
	EmployeeID    string         `json:"employeeId"`
	EmployeeName  string         `json:"employeeName,omitempty"`
	ShiftID       string         `json:"shiftId"`
	Shift         *ShiftResponse `json:"shift,omitempty"`
	EffectiveFrom string         `json:"effectiveFrom"`
	EffectiveTo   string         `json:"effectiveTo,omitempty"`
	IsActive      bool           `json:"isActive"`
	CreatedAt     int64          `json:"createdAt"`
}
