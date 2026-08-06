package stock

import (
	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/modules/inventory/product"
	"github.com/tamim1715/novaerp/internal/modules/inventory/warehouse"
)

// Stock tracks the quantity of a product in a specific warehouse.
type Stock struct {
	model.BaseModel

	WarehouseID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_warehouse_product" json:"warehouseId"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_warehouse_product" json:"productId"`
	Quantity    int       `gorm:"default:0;not null" json:"quantity"`

	Warehouse warehouse.Warehouse `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
	Product   product.Product     `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
