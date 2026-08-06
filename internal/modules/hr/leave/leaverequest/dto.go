package leaverequest

type CreateLeaveRequest struct {
	EmployeeID  string `json:"employeeId" binding:"required" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`
	LeaveTypeID string `json:"leaveTypeId" binding:"required" example:"8a393bbd-9790-4045-853e-26a2c732ee09"`
	StartDate   string `json:"startDate" binding:"required" example:"2026-08-10"` // YYYY-MM-DD
	EndDate     string `json:"endDate" binding:"required" example:"2026-08-12"`   // YYYY-MM-DD
	Reason      string `json:"reason" example:"Family vacation"`
}

type UpdateLeaveStatusRequest struct {
	Status     string `json:"status" binding:"required,oneof=APPROVED REJECTED" example:"APPROVED"`
	ApprovedBy string `json:"approvedBy" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`
	Remarks    string `json:"remarks" example:"Approved by manager"`
}

type LeaveRequestResponse struct {
	ID            string `json:"id"`
	EmployeeID    string `json:"employeeId"`
	EmployeeName  string `json:"employeeName,omitempty"`
	LeaveTypeID   string `json:"leaveTypeId"`
	LeaveTypeName string `json:"leaveTypeName,omitempty"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	TotalDays     int    `json:"totalDays"`
	Reason        string `json:"reason"`
	Status        string `json:"status"`
	ApprovedBy    string `json:"approvedBy,omitempty"`
	Remarks       string `json:"remarks,omitempty"`
	CreatedAt     int64  `json:"createdAt"`
}
