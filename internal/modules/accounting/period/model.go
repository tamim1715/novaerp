package period

import (
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/enums"
)

type Status = enums.Status

const (
	StatusOpen   = enums.StatusOpen
	StatusClosed = enums.StatusClosed
)

// FiscalYear represents a financial reporting year (e.g. FY 2026).
type FiscalYear struct {
	model.BaseModel

	Name      string    `gorm:"size:50;uniqueIndex;not null" json:"name"`
	StartDate time.Time `gorm:"type:date;not null" json:"startDate"`
	EndDate   time.Time `gorm:"type:date;not null" json:"endDate"`
	IsClosed  bool      `gorm:"default:false" json:"isClosed"`

	Periods []AccountingPeriod `gorm:"foreignKey:FiscalYearID" json:"periods,omitempty"`
}

// AccountingPeriod represents monthly financial sub-periods within a fiscal year.
type AccountingPeriod struct {
	model.BaseModel

	FiscalYearID uuid.UUID    `gorm:"type:uuid;not null;index" json:"fiscalYearId"`
	PeriodNumber int          `gorm:"not null" json:"periodNumber"` // 1 to 12
	Name         string       `gorm:"size:50;not null" json:"name"` // e.g. "2026-01 (Jan)"
	StartDate    time.Time    `gorm:"type:date;not null" json:"startDate"`
	EndDate      time.Time    `gorm:"type:date;not null" json:"endDate"`
	Status       enums.Status `gorm:"size:20;default:'OPEN';not null" json:"status"` // OPEN, CLOSED

	FiscalYear *FiscalYear `gorm:"foreignKey:FiscalYearID" json:"fiscalYear,omitempty"`
}
