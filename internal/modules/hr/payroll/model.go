package payroll

import (
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/modules/employee"
)

// PayrollPeriod represents a monthly payroll batch (e.g. Month 8, Year 2026).
type PayrollPeriod struct {
	model.BaseModel

	Month       int       `gorm:"not null;uniqueIndex:idx_month_year" json:"month"`
	Year        int       `gorm:"not null;uniqueIndex:idx_month_year" json:"year"`
	StartDate   time.Time `gorm:"type:date;not null" json:"startDate"`
	EndDate     time.Time `gorm:"type:date;not null" json:"endDate"`
	Status      string    `gorm:"size:20;default:'DRAFT'" json:"status"` // DRAFT, PROCESSING, APPROVED, PAID
	TotalGross  float64   `gorm:"type:numeric(14,2);default:0" json:"totalGross"`
	TotalNet    float64   `gorm:"type:numeric(14,2);default:0" json:"totalNet"`
	ProcessedAt *time.Time`gorm:"type:timestamp" json:"processedAt,omitempty"`

	Payslips []Payslip `gorm:"foreignKey:PayrollPeriodID" json:"payslips,omitempty"`
}

// Payslip represents individual employee salary slip for a payroll period.
type Payslip struct {
	model.BaseModel

	PayrollPeriodID     uuid.UUID `gorm:"type:uuid;not null;index" json:"payrollPeriodId"`
	EmployeeID          uuid.UUID `gorm:"type:uuid;not null;index" json:"employeeId"`
	BasicSalary         float64   `gorm:"type:numeric(12,2);not null" json:"basicSalary"`
	Allowances          float64   `gorm:"type:numeric(12,2);default:0" json:"allowances"`
	Deductions          float64   `gorm:"type:numeric(12,2);default:0" json:"deductions"`
	UnpaidLeaveDeduction float64  `gorm:"type:numeric(12,2);default:0" json:"unpaidLeaveDeduction"`
	GrossSalary         float64   `gorm:"type:numeric(12,2);not null" json:"grossSalary"`
	NetSalary           float64   `gorm:"type:numeric(12,2);not null" json:"netSalary"`
	Status              string    `gorm:"size:20;default:'DRAFT'" json:"status"` // DRAFT, PAID
	PaymentDate         *time.Time`gorm:"type:timestamp" json:"paymentDate,omitempty"`
	Notes               string    `gorm:"size:255" json:"notes"`

	Employee      employee.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	PayrollPeriod PayrollPeriod     `gorm:"foreignKey:PayrollPeriodID" json:"payrollPeriod,omitempty"`
}
