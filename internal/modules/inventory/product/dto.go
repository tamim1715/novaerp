package product

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
