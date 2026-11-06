package attendance

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/hr/shift"
)

func TestToAttendanceResponse(t *testing.T) {
	empID := uuid.New()
	shiftID := uuid.New()
	now := time.Now()

	att := &Attendance{
		EmployeeID:    empID,
		ShiftID:       &shiftID,
		Date:          now,
		CheckIn:       &now,
		WorkHours:     8.5,
		OvertimeHours: 0.5,
		LateMinutes:   15,
		Status:        StatusLate,
		Employee:      employee.Employee{FirstName: "John", LastName: "Doe"},
		Shift:         &shift.Shift{Name: "General Shift"},
	}

	res := ToAttendanceResponse(att)
	if res.Status != "LATE" {
		t.Errorf("expected Status LATE, got %s", res.Status)
	}
	if res.WorkHours != 8.5 {
		t.Errorf("expected WorkHours 8.5, got %f", res.WorkHours)
	}
	if res.LateMinutes != 15 {
		t.Errorf("expected LateMinutes 15, got %d", res.LateMinutes)
	}
	if res.ShiftName != "General Shift" {
		t.Errorf("expected ShiftName General Shift, got %s", res.ShiftName)
	}
	if res.EmployeeName != "John Doe" {
		t.Errorf("expected EmployeeName John Doe, got %s", res.EmployeeName)
	}
}

func TestMonthlyAttendanceSummary(t *testing.T) {
	summary := &MonthlyAttendanceSummary{
		EmployeeID:         uuid.New().String(),
		EmployeeName:       "Jane Doe",
		TotalWorkingDays:   22,
		PresentDays:        20,
		LateDays:           2,
		AbsentDays:         1,
		LeaveDays:          1,
		TotalWorkHours:     160.0,
		TotalOvertimeHours: 10.0,
	}

	if summary.PresentDays+summary.AbsentDays+summary.LeaveDays != 22 {
		t.Errorf("expected sum of present, absent and leave days to equal working days")
	}
	if summary.TotalOvertimeHours != 10.0 {
		t.Errorf("expected 10.0 overtime hours")
	}
}
