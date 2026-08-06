package stock

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
