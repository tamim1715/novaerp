package employee

import (
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/modules/department"
)

type Employee struct {
	model.BaseModel

	DepartmentID uuid.UUID `gorm:"type:uuid;not null"`

	FirstName string `gorm:"size:100;not null"`

	LastName string `gorm:"size:100"`

	Email string `gorm:"size:255;uniqueIndex;not null"`

	Phone string `gorm:"size:20;uniqueIndex"`

	Designation string `gorm:"size:100"`

	JoiningDate time.Time

	Salary float64 `gorm:"type:numeric(12,2)"`

	Status bool `gorm:"default:true"`

	Department department.Department `gorm:"foreignKey:DepartmentID"`
}
