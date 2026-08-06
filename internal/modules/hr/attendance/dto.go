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
	EmployeeID string  `json:"employeeId" binding:"required" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`
	Date       string  `json:"date" binding:"required" example:"2026-08-06"`
	CheckIn    string  `json:"checkIn" example:"2026-08-06T09:00:00Z"`
	CheckOut   string  `json:"checkOut" example:"2026-08-06T17:00:00Z"`
	Status     string  `json:"status" example:"PRESENT"` // PRESENT, ABSENT, LATE, HALF_DAY, ON_LEAVE
	Notes      string  `json:"notes" example:"Manual attendance log"`
}

type AttendanceResponse struct {
	ID            string  `json:"id"`
	EmployeeID    string  `json:"employeeId"`
	EmployeeName  string  `json:"employeeName,omitempty"`
	Date          string  `json:"date"`
	CheckIn       string  `json:"checkIn,omitempty"`
	CheckOut      string  `json:"checkOut,omitempty"`
	WorkHours     float64 `json:"workHours"`
	OvertimeHours float64 `json:"overtimeHours"`
	Status        string  `json:"status"`
	Notes         string  `json:"notes"`
	CreatedAt     int64   `json:"createdAt"`
}
