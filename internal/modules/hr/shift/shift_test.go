package shift

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/modules/employee"
)

func TestToShiftResponse(t *testing.T) {
	s := &Shift{
		Name:               "Morning Shift",
		Code:               "MORN_09",
		StartTime:          "09:00",
		EndTime:            "17:00",
		GracePeriodMinutes: 15,
		WorkHours:          8.0,
		HalfDayHours:       4.0,
		IsActive:           true,
	}
	s.ID = uuid.New()
	s.CreatedAt = time.Now()

	resp := ToShiftResponse(s)
	if resp.Code != "MORN_09" {
		t.Errorf("expected Code MORN_09, got %s", resp.Code)
	}
	if resp.GracePeriodMinutes != 15 {
		t.Errorf("expected GracePeriodMinutes 15, got %d", resp.GracePeriodMinutes)
	}
}

func TestToEmployeeShiftResponse(t *testing.T) {
	empID := uuid.New()
	shiftID := uuid.New()
	now := time.Now()

	es := &EmployeeShift{
		EmployeeID:    empID,
		ShiftID:       shiftID,
		EffectiveFrom: now,
		IsActive:      true,
		Employee:      employee.Employee{FirstName: "John", LastName: "Doe"},
		Shift:         Shift{Name: "General Shift", Code: "GEN"},
	}
	es.ID = uuid.New()

	resp := ToEmployeeShiftResponse(es)
	if resp.EmployeeName != "John Doe" {
		t.Errorf("expected EmployeeName John Doe, got %s", resp.EmployeeName)
	}
	if resp.Shift == nil || resp.Shift.Code != "GEN" {
		t.Errorf("expected Shift Code GEN")
	}
}
