package leavetype

import (
	"github.com/tamim1715/novaerp/internal/common/model"
)

// LeaveType defines categories of leave (e.g. Annual, Sick, Casual).
type LeaveType struct {
	model.BaseModel

	Name           string `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Code           string `gorm:"size:20;not null;uniqueIndex" json:"code"`
	MaxDaysPerYear int    `gorm:"not null" json:"maxDaysPerYear"`
	IsPaid         bool   `gorm:"default:true" json:"isPaid"`
	Description    string `gorm:"size:255" json:"description"`
}
