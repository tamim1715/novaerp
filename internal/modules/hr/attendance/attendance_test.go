package attendance

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/modules/employee"
)

func TestToAttendanceResponse(t *testing.T) {
	empID := uuid.New()
	now := time.Now()

	att := &Attendance{
		EmployeeID:    empID,
		Date:          now,
		CheckIn:       &now,
		WorkHours:     8.5,
		OvertimeHours: 0.5,
		Status:        StatusPresent,
		Employee:      employee.Employee{FirstName: "John", LastName: "Doe"},
	}

	res := ToAttendanceResponse(att)
	if res.Status != "PRESENT" {
		t.Errorf("expected Status PRESENT, got %s", res.Status)
	}
	if res.WorkHours != 8.5 {
		t.Errorf("expected WorkHours 8.5, got %f", res.WorkHours)
	}
	if res.EmployeeName != "John Doe" {
		t.Errorf("expected EmployeeName John Doe, got %s", res.EmployeeName)
	}
}
