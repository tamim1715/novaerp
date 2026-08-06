package product

import (
	"github.com/tamim1715/novaerp/internal/common/model"
)

// Product represents an item stored in inventory.
type Product struct {
	model.BaseModel

	SKU           string  `gorm:"size:50;not null;uniqueIndex" json:"sku"`
	Name          string  `gorm:"size:150;not null" json:"name"`
	Category      string  `gorm:"size:100" json:"category"`
	Unit          string  `gorm:"size:20;default:'pcs'" json:"unit"`
	UnitPrice     float64 `gorm:"type:numeric(12,2);default:0" json:"unitPrice"`
	MinStockLevel int     `gorm:"default:0" json:"minStockLevel"`
	Description   string  `gorm:"size:255" json:"description"`
}
