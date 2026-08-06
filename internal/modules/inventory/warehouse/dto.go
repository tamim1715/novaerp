package warehouse

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
