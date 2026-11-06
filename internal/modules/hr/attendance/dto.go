package attendance

type CheckInRequest struct {
	EmployeeID string `json:"employeeId" binding:"required" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`
	Notes      string `json:"notes" example:"Arrived at main office"`
}

type CheckOutRequest struct {
	EmployeeID string `json:"employeeId" binding:"required" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`
	Notes      string `json:"notes" example:"Completed shift"`
}

type CreateAttendanceRequest struct {
	EmployeeID string `json:"employeeId" binding:"required" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`
	ShiftID    string `json:"shiftId" example:"5e53303c-8ab5-46eb-a1f9-9ea2eb60b37f"`
	Date       string `json:"date" binding:"required" example:"2026-08-06"`
	CheckIn    string `json:"checkIn" example:"2026-08-06T09:00:00Z"`
	CheckOut   string `json:"checkOut" example:"2026-08-06T17:00:00Z"`
	Status     string `json:"status" example:"PRESENT"` // PRESENT, ABSENT, LATE, HALF_DAY, ON_LEAVE, HOLIDAY, WEEKEND
	Notes      string `json:"notes" example:"Manual attendance log"`
}

type AttendanceResponse struct {
	ID            string  `json:"id"`
	EmployeeID    string  `json:"employeeId"`
	EmployeeName  string  `json:"employeeName,omitempty"`
	ShiftID       string  `json:"shiftId,omitempty"`
	ShiftName     string  `json:"shiftName,omitempty"`
	Date          string  `json:"date"`
	CheckIn       string  `json:"checkIn,omitempty"`
	CheckOut      string  `json:"checkOut,omitempty"`
	WorkHours     float64 `json:"workHours"`
	OvertimeHours float64 `json:"overtimeHours"`
	LateMinutes   int     `json:"lateMinutes"`
	Status        string  `json:"status"`
	Notes         string  `json:"notes"`
	CreatedAt     int64   `json:"createdAt"`
}

type ProcessDailyAttendanceRequest struct {
	Date         string `json:"date" binding:"required" example:"2026-08-06"`
	DepartmentID string `json:"departmentId" example:"d290f1ee-6c54-4b01-90e6-d701748f0851"`
}

type DailyAttendanceProcessResult struct {
	Date           string `json:"date"`
	TotalEmployees int    `json:"totalEmployees"`
	ProcessedCount int    `json:"processedCount"`
	PresentCount   int    `json:"presentCount"`
	AbsentCount    int    `json:"absentCount"`
	LateCount      int    `json:"lateCount"`
	LeaveCount     int    `json:"leaveCount"`
}

type MonthlyAttendanceSummary struct {
	EmployeeID         string  `json:"employeeId"`
	EmployeeName       string  `json:"employeeName,omitempty"`
	StartDate          string  `json:"startDate"`
	EndDate            string  `json:"endDate"`
	TotalWorkingDays   int     `json:"totalWorkingDays"`
	PresentDays        int     `json:"presentDays"`
	LateDays           int     `json:"lateDays"`
	HalfDays           int     `json:"halfDays"`
	AbsentDays         int     `json:"absentDays"`
	LeaveDays          int     `json:"leaveDays"`
	TotalWorkHours     float64 `json:"totalWorkHours"`
	TotalOvertimeHours float64 `json:"totalOvertimeHours"`
}
