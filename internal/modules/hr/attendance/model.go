package attendance

import (
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/modules/employee"
)

// Attendance records daily work hours and check-in/out times.
type Attendance struct {
	model.BaseModel

	EmployeeID    uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:idx_employee_date" json:"employeeId"`
	Date          time.Time  `gorm:"type:date;not null;uniqueIndex:idx_employee_date" json:"date"`
	CheckIn       *time.Time `gorm:"type:timestamp" json:"checkIn"`
	CheckOut      *time.Time `gorm:"type:timestamp" json:"checkOut"`
	WorkHours     float64    `gorm:"type:numeric(5,2);default:0" json:"workHours"`
	OvertimeHours float64    `gorm:"type:numeric(5,2);default:0" json:"overtimeHours"`
	Status        string     `gorm:"size:20;default:'PRESENT'" json:"status"` // PRESENT, ABSENT, LATE, HALF_DAY, ON_LEAVE
	Notes         string     `gorm:"size:255" json:"notes"`

	Employee employee.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
}
