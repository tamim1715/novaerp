package leaverequest

import (
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/hr/leave/leavetype"
)

// LeaveRequest represents an employee's application for leave.
type LeaveRequest struct {
	model.BaseModel

	EmployeeID  uuid.UUID  `gorm:"type:uuid;not null;index" json:"employeeId"`
	LeaveTypeID uuid.UUID  `gorm:"type:uuid;not null;index" json:"leaveTypeId"`
	StartDate   time.Time  `gorm:"type:date;not null" json:"startDate"`
	EndDate     time.Time  `gorm:"type:date;not null" json:"endDate"`
	TotalDays   int        `gorm:"not null" json:"totalDays"`
	Reason      string     `gorm:"size:255" json:"reason"`
	Status      string     `gorm:"size:20;default:'PENDING'" json:"status"` // PENDING, APPROVED, REJECTED
	ApprovedBy  *uuid.UUID `gorm:"type:uuid" json:"approvedBy,omitempty"`
	Remarks     string     `gorm:"size:255" json:"remarks"`

	Employee  employee.Employee   `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	LeaveType leavetype.LeaveType `gorm:"foreignKey:LeaveTypeID" json:"leaveType,omitempty"`
}
