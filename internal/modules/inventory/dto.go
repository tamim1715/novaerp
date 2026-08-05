package inventory

// Warehouse DTOs
type CreateWarehouseRequest struct {
	Name     string `json:"name" binding:"required,max=100" example:"Main Warehouse"`
	Code     string `json:"code" binding:"required,max=30" example:"WH-001"`
	Location string `json:"location" example:"Building A, Industrial Area"`
}

type UpdateWarehouseRequest struct {
	Name     string `json:"name" example:"Main Warehouse Updated"`
	Code     string `json:"code" example:"WH-001"`
	Location string `json:"location" example:"Building A, Industrial Area"`
	IsActive *bool  `json:"isActive" example:"true"`
}

type WarehouseResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Location  string `json:"location"`
	IsActive  bool   `json:"isActive"`
	CreatedAt int64  `json:"createdAt"`
}

// Product DTOs
type CreateProductRequest struct {
	SKU           string  `json:"sku" binding:"required,max=50" example:"PRD-1001"`
	Name          string  `json:"name" binding:"required,max=150" example:"Laptop Dell XPS 15"`
	Category      string  `json:"category" example:"Electronics"`
	Unit          string  `json:"unit" example:"pcs"`
	UnitPrice     float64 `json:"unitPrice" example:"1499.99"`
	MinStockLevel int     `json:"minStockLevel" example:"5"`
	Description   string  `json:"description" example:"High performance laptop"`
}

type UpdateProductRequest struct {
	SKU           string  `json:"sku" example:"PRD-1001"`
	Name          string  `json:"name" example:"Laptop Dell XPS 15"`
	Category      string  `json:"category" example:"Electronics"`
	Unit          string  `json:"unit" example:"pcs"`
	UnitPrice     float64 `json:"unitPrice" example:"1499.99"`
	MinStockLevel *int    `json:"minStockLevel" example:"5"`
	Description   string  `json:"description" example:"High performance laptop"`
}

type ProductResponse struct {
	ID            string  `json:"id"`
	SKU           string  `json:"sku"`
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	Unit          string  `json:"unit"`
	UnitPrice     float64 `json:"unitPrice"`
	MinStockLevel int     `json:"minStockLevel"`
	Description   string  `json:"description"`
	CreatedAt     int64   `json:"createdAt"`
}

// Stock DTOs
type AdjustStockRequest struct {
	WarehouseID   string `json:"warehouseId" binding:"required" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`
	ProductID     string `json:"productId" binding:"required" example:"8a393bbd-9790-4045-853e-26a2c732ee09"`
	QuantityDelta int    `json:"quantityDelta" binding:"required" example:"50"` // Positive to add, negative to reduce
}

type StockResponse struct {
	ID            string `json:"id"`
	WarehouseID   string `json:"warehouseId"`
	ProductID     string `json:"productId"`
	Quantity      int    `json:"quantity"`
	WarehouseName string `json:"warehouseName,omitempty"`
	ProductName   string `json:"productName,omitempty"`
	ProductSKU    string `json:"productSku,omitempty"`
	CreatedAt     int64  `json:"createdAt"`
}
