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
		PayrollPeriodID: periodID,
		EmployeeID:      empID,
		BasicSalary:     3000.00,
		Allowances:      200.00,
		Deductions:      100.00,
		GrossSalary:     3200.00,
		NetSalary:       3100.00,
		Status:          StatusDraft,
		Employee:        employee.Employee{FirstName: "Jane", LastName: "Smith", Designation: "Engineer"},
	}

	res := ToPayslipResponse(ps)
	if res.BasicSalary != 3000.00 {
		t.Errorf("expected BasicSalary 3000, got %f", res.BasicSalary)
	}
	if res.NetSalary != 3100.00 {
		t.Errorf("expected NetSalary 3100, got %f", res.NetSalary)
	}
	if res.EmployeeName != "Jane Smith" {
		t.Errorf("expected EmployeeName Jane Smith, got %s", res.EmployeeName)
	}
}
