package enums

import "testing"

func TestStatusString(t *testing.T) {
	if StatusActive.String() != "active" {
		t.Errorf("expected 'active', got '%s'", StatusActive.String())
	}
	if StatusPosted.String() != "POSTED" {
		t.Errorf("expected 'POSTED', got '%s'", StatusPosted.String())
	}
}

func TestStatusValidation(t *testing.T) {
	// Employee status
	if !StatusActive.IsValidEmployee() {
		t.Errorf("expected StatusActive to be valid employee status")
	}
	if StatusPosted.IsValidEmployee() {
		t.Errorf("expected StatusPosted to NOT be valid employee status")
	}

	// Attendance status
	if !StatusPresent.IsValidAttendance() || !StatusLate.IsValidAttendance() {
		t.Errorf("expected StatusPresent and StatusLate to be valid attendance status")
	}
	if StatusClosed.IsValidAttendance() {
		t.Errorf("expected StatusClosed to NOT be valid attendance status")
	}

	// Leave status
	if !StatusPending.IsValidLeave() || !StatusApproved.IsValidLeave() {
		t.Errorf("expected StatusPending and StatusApproved to be valid leave status")
	}
	if StatusActive.IsValidLeave() {
		t.Errorf("expected StatusActive to NOT be valid leave status")
	}

	// PayrollPeriod status
	if !StatusDraft.IsValidPayrollPeriod() || !StatusProcessing.IsValidPayrollPeriod() || !StatusPaid.IsValidPayrollPeriod() {
		t.Errorf("expected draft, processing, and paid to be valid payroll period status")
	}

	// Payslip status
	if !StatusDraft.IsValidPayslip() || !StatusPaid.IsValidPayslip() {
		t.Errorf("expected draft and paid to be valid payslip status")
	}
	if StatusProcessing.IsValidPayslip() {
		t.Errorf("expected processing to NOT be valid payslip status")
	}

	// Accounting period status
	if !StatusOpen.IsValidAccountingPeriod() || !StatusClosed.IsValidAccountingPeriod() {
		t.Errorf("expected open and closed to be valid accounting period status")
	}
	if StatusDraft.IsValidAccountingPeriod() {
		t.Errorf("expected draft to NOT be valid accounting period status")
	}

	// Journal status
	if !StatusDraft.IsValidJournal() || !StatusPosted.IsValidJournal() || !StatusVoid.IsValidJournal() {
		t.Errorf("expected draft, posted, and void to be valid journal status")
	}
	if StatusOpen.IsValidJournal() {
		t.Errorf("expected open to NOT be valid journal status")
	}
}
