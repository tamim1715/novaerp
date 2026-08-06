package warehouse

import (
	"github.com/tamim1715/novaerp/internal/common/model"
)

// Warehouse represents a storage location.
type Warehouse struct {
	model.BaseModel

	Name     string `gorm:"size:100;not null" json:"name"`
	Code     string `gorm:"size:30;not null;uniqueIndex" json:"code"`
	Location string `gorm:"size:255" json:"location"`
	IsActive bool   `gorm:"default:true" json:"isActive"`
}
