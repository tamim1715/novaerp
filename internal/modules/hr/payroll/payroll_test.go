package payroll

import (
	"testing"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/modules/employee"
)

func TestToPayslipResponse(t *testing.T) {
	empID := uuid.New()
	periodID := uuid.New()

	ps := &Payslip{
		PayrollPeriodID:      periodID,
		EmployeeID:           empID,
		BasicSalary:          3000.00,
		Allowances:           200.00,
		OvertimeHours:        5.0,
		OvertimePay:          125.00,
		Deductions:           100.00,
		AbsentDays:           1,
		AbsentDeduction:      136.36,
		UnpaidLeaveDays:      0,
		UnpaidLeaveDeduction: 0,
		PresentDays:          21,
		WorkingDays:          22,
		GrossSalary:          3325.00,
		NetSalary:            3088.64,
		Status:               StatusDraft,
		Employee:             employee.Employee{FirstName: "Jane", LastName: "Smith", Designation: "Engineer"},
	}

	res := ToPayslipResponse(ps)
	if res.BasicSalary != 3000.00 {
		t.Errorf("expected BasicSalary 3000, got %f", res.BasicSalary)
	}
	if res.OvertimePay != 125.00 {
		t.Errorf("expected OvertimePay 125, got %f", res.OvertimePay)
	}
	if res.AbsentDays != 1 {
		t.Errorf("expected AbsentDays 1, got %d", res.AbsentDays)
	}
	if res.NetSalary != 3088.64 {
		t.Errorf("expected NetSalary 3088.64, got %f", res.NetSalary)
	}
	if res.EmployeeName != "Jane Smith" {
		t.Errorf("expected EmployeeName Jane Smith, got %s", res.EmployeeName)
	}
}
