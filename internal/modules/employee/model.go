package employee

import (
	"time"

	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/modules/department"
)

type Employee struct {
	model.BaseModel

	DepartmentID string `gorm:"type:uuid;not null"`

	FirstName string `gorm:"size:100;not null"`

	LastName string `gorm:"size:100"`

	Email string `gorm:"size:255;uniqueIndex;not null"`

	Phone string `gorm:"size:20;uniqueIndex"`

	Designation string `gorm:"size:100"`

	JoiningDate time.Time `gorm:"type:timestamp;not null"`

	Salary float64 `gorm:"type:numeric(12,2)"`

	Status string `gorm:"default:'active'"`

	Department department.Department `gorm:"foreignKey:DepartmentID"`
}
