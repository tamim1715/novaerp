package allowance

import (
	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/modules/employee"
)

// AllowanceCalculationType defines whether allowance is fixed or percentage of basic salary.
type AllowanceCalculationType string

const (
	CalculationFixed      AllowanceCalculationType = "FIXED"
	CalculationPercentage AllowanceCalculationType = "PERCENTAGE"
)

// AllowanceType defines recurring compensation components (e.g. HRA, Medical, Travel).
type AllowanceType struct {
	model.BaseModel

	Name            string                   `gorm:"size:100;not null" json:"name"`
	Code            string                   `gorm:"size:50;not null;uniqueIndex" json:"code"`
	CalculationType AllowanceCalculationType `gorm:"size:20;default:'FIXED'" json:"calculationType"` // FIXED, PERCENTAGE
	DefaultAmount   float64                  `gorm:"type:numeric(12,2);default:0" json:"defaultAmount"`
	IsTaxable       bool                     `gorm:"default:false" json:"isTaxable"`
	IsActive        bool                     `gorm:"default:true" json:"isActive"`
	Description     string                   `gorm:"size:255" json:"description"`
}

// EmployeeAllowance assigns an allowance type to a specific employee.
type EmployeeAllowance struct {
	model.BaseModel

	EmployeeID      uuid.UUID `gorm:"type:uuid;not null;index" json:"employeeId"`
	AllowanceTypeID uuid.UUID `gorm:"type:uuid;not null;index" json:"allowanceTypeId"`
	Amount          float64   `gorm:"type:numeric(12,2);default:0" json:"amount"` // Override amount; if 0, uses DefaultAmount
	IsActive        bool      `gorm:"default:true" json:"isActive"`

	Employee      employee.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	AllowanceType AllowanceType     `gorm:"foreignKey:AllowanceTypeID" json:"allowanceType,omitempty"`
}
