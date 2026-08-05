package inventory

import (
	"github.com/google/uuid"
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

// Stock tracks the quantity of a product in a specific warehouse.
type Stock struct {
	model.BaseModel

	WarehouseID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_warehouse_product" json:"warehouseId"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_warehouse_product" json:"productId"`
	Quantity    int       `gorm:"default:0;not null" json:"quantity"`

	Warehouse Warehouse `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
