package leaverequest

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/hr/leave/leavetype"
)

func TestToLeaveRequestResponse(t *testing.T) {
	empID := uuid.New()
	ltID := uuid.New()
	now := time.Now()

	lr := &LeaveRequest{
		EmployeeID:  empID,
		LeaveTypeID: ltID,
		StartDate:   now,
		EndDate:     now.AddDate(0, 0, 2),
		TotalDays:   3,
		Reason:      "Vacation",
		Status:      StatusPending,
		Employee:    employee.Employee{FirstName: "Alex", LastName: "Morgan"},
		LeaveType:   leavetype.LeaveType{Name: "Annual Leave"},
	}

	res := ToLeaveRequestResponse(lr)
	if res.Status != "PENDING" {
		t.Errorf("expected Status PENDING, got %s", res.Status)
	}
	if res.TotalDays != 3 {
		t.Errorf("expected TotalDays 3, got %d", res.TotalDays)
	}
	if res.EmployeeName != "Alex Morgan" {
		t.Errorf("expected EmployeeName Alex Morgan, got %s", res.EmployeeName)
	}
}
