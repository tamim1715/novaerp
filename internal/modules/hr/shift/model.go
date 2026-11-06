package shift

import (
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/modules/employee"
)

// Shift represents a work shift timing configuration.
type Shift struct {
	model.BaseModel

	Name               string  `gorm:"size:100;not null" json:"name"`
	Code               string  `gorm:"size:50;not null;uniqueIndex" json:"code"`
	StartTime          string  `gorm:"size:10;not null" json:"startTime"` // e.g. "09:00"
	EndTime            string  `gorm:"size:10;not null" json:"endTime"`   // e.g. "17:00"
	GracePeriodMinutes int     `gorm:"default:15" json:"gracePeriodMinutes"`
	WorkHours          float64 `gorm:"type:numeric(5,2);default:8.0" json:"workHours"`
	HalfDayHours       float64 `gorm:"type:numeric(5,2);default:4.0" json:"halfDayHours"`
	IsNightShift       bool    `gorm:"default:false" json:"isNightShift"`
	IsActive           bool    `gorm:"default:true" json:"isActive"`
	Description        string  `gorm:"size:255" json:"description"`
}

// EmployeeShift maps an employee to an assigned work shift.
type EmployeeShift struct {
	model.BaseModel

	EmployeeID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"employeeId"`
	ShiftID       uuid.UUID  `gorm:"type:uuid;not null;index" json:"shiftId"`
	EffectiveFrom time.Time  `gorm:"type:date;not null" json:"effectiveFrom"`
	EffectiveTo   *time.Time `gorm:"type:date" json:"effectiveTo,omitempty"`
	IsActive      bool       `gorm:"default:true" json:"isActive"`

	Employee employee.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	Shift    Shift             `gorm:"foreignKey:ShiftID" json:"shift,omitempty"`
}
