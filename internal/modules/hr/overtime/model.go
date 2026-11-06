package overtime

import (
	"github.com/tamim1715/novaerp/internal/common/model"
)

// OvertimePolicy defines the rules and rate multipliers for overtime calculation.
type OvertimePolicy struct {
	model.BaseModel

	Name                   string  `gorm:"size:100;not null" json:"name"`
	Code                   string  `gorm:"size:50;not null;uniqueIndex" json:"code"`
	MinOvertimeMinutes     int     `gorm:"default:30" json:"minOvertimeMinutes"`                        // Minimum threshold before OT is credited
	MaxOvertimeHoursPerDay float64 `gorm:"type:numeric(5,2);default:4.0" json:"maxOvertimeHoursPerDay"` // Daily cap
	WeekdayRateMultiplier  float64 `gorm:"type:numeric(4,2);default:1.5" json:"weekdayRateMultiplier"`  // e.g. 1.5x
	WeekendRateMultiplier  float64 `gorm:"type:numeric(4,2);default:2.0" json:"weekendRateMultiplier"`  // e.g. 2.0x
	HolidayRateMultiplier  float64 `gorm:"type:numeric(4,2);default:2.0" json:"holidayRateMultiplier"`  // e.g. 2.0x
	IsDefault              bool    `gorm:"default:false" json:"isDefault"`
	IsActive               bool    `gorm:"default:true" json:"isActive"`
	Description            string  `gorm:"size:255" json:"description"`
}
