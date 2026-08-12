// Package enums provides centralized domain enum types and constants for the ERP system.
package enums

// Status represents the standard status enum type used across NovaERP domains.
type Status string

// String returns the string representation of the Status.
func (s Status) String() string {
	return string(s)
}

// Domain-specific status type aliases for semantic clarity.
type EmployeeStatus = Status
type AttendanceStatus = Status
type LeaveStatus = Status
type PayrollStatus = Status
type PeriodStatus = Status
type JournalStatus = Status

const (
	// Employee Statuses
	StatusActive     Status = "active"
	StatusInactive   Status = "inactive"
	StatusTerminated Status = "terminated"
	StatusOnLeave    Status = "on_leave"

	// Attendance Statuses
	StatusPresent AttendanceStatus = "PRESENT"
	StatusAbsent  AttendanceStatus = "ABSENT"
	StatusLate    AttendanceStatus = "LATE"
	StatusHalfDay AttendanceStatus = "HALF_DAY"

	// Leave Request Statuses
	StatusPending   LeaveStatus = "PENDING"
	StatusApproved  LeaveStatus = "APPROVED"
	StatusRejected  LeaveStatus = "REJECTED"
	StatusCancelled LeaveStatus = "CANCELLED"

	// Payroll Batch & Payslip Statuses
	StatusDraft      PayrollStatus = "DRAFT"
	StatusProcessing PayrollStatus = "PROCESSING"
	StatusPaid       PayrollStatus = "PAID"

	// Accounting Period Statuses
	StatusOpen   PeriodStatus = "OPEN"
	StatusClosed PeriodStatus = "CLOSED"

	// Journal Entry Statuses
	StatusPosted JournalStatus = "POSTED"
	StatusVoid   JournalStatus = "VOID"
)

// ValidEmployeeStatuses returns all valid status values for employees.
func ValidEmployeeStatuses() []Status {
	return []Status{
		StatusActive,
		StatusInactive,
		StatusTerminated,
		StatusOnLeave,
	}
}

// ValidAttendanceStatuses returns all valid status values for daily attendance.
func ValidAttendanceStatuses() []Status {
	return []Status{
		StatusPresent,
		StatusAbsent,
		StatusLate,
		StatusHalfDay,
		StatusOnLeave,
	}
}

// ValidLeaveStatuses returns all valid status values for employee leave requests.
func ValidLeaveStatuses() []Status {
	return []Status{
		StatusPending,
		StatusApproved,
		StatusRejected,
		StatusCancelled,
	}
}

// ValidPayrollPeriodStatuses returns all valid status values for payroll processing periods.
func ValidPayrollPeriodStatuses() []Status {
	return []Status{
		StatusDraft,
		StatusProcessing,
		StatusApproved,
		StatusPaid,
	}
}

// ValidPayslipStatuses returns all valid status values for individual payslips.
func ValidPayslipStatuses() []Status {
	return []Status{
		StatusDraft,
		StatusPaid,
	}
}

// ValidAccountingPeriodStatuses returns all valid status values for accounting periods.
func ValidAccountingPeriodStatuses() []Status {
	return []Status{
		StatusOpen,
		StatusClosed,
	}
}

// ValidJournalStatuses returns all valid status values for general ledger journal entries.
func ValidJournalStatuses() []Status {
	return []Status{
		StatusDraft,
		StatusPosted,
		StatusVoid,
	}
}

// IsValidEmployee checks whether the status is a valid employee status.
func (s Status) IsValidEmployee() bool {
	for _, v := range ValidEmployeeStatuses() {
		if s == v {
			return true
		}
	}
	return false
}

// IsValidAttendance checks whether the status is a valid attendance status.
func (s Status) IsValidAttendance() bool {
	for _, v := range ValidAttendanceStatuses() {
		if s == v {
			return true
		}
	}
	return false
}

// IsValidLeave checks whether the status is a valid leave request status.
func (s Status) IsValidLeave() bool {
	for _, v := range ValidLeaveStatuses() {
		if s == v {
			return true
		}
	}
	return false
}

// IsValidPayrollPeriod checks whether the status is a valid payroll period status.
func (s Status) IsValidPayrollPeriod() bool {
	for _, v := range ValidPayrollPeriodStatuses() {
		if s == v {
			return true
		}
	}
	return false
}

// IsValidPayslip checks whether the status is a valid payslip status.
func (s Status) IsValidPayslip() bool {
	for _, v := range ValidPayslipStatuses() {
		if s == v {
			return true
		}
	}
	return false
}

// IsValidAccountingPeriod checks whether the status is a valid accounting period status.
func (s Status) IsValidAccountingPeriod() bool {
	for _, v := range ValidAccountingPeriodStatuses() {
		if s == v {
			return true
		}
	}
	return false
}

// IsValidJournal checks whether the status is a valid journal entry status.
func (s Status) IsValidJournal() bool {
	for _, v := range ValidJournalStatuses() {
		if s == v {
			return true
		}
	}
	return false
}
