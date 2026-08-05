package inventory

func ToWarehouseResponse(w *Warehouse) WarehouseResponse {
	if w == nil {
		return WarehouseResponse{}
	}
	return WarehouseResponse{
		ID:        w.ID.String(),
		Name:      w.Name,
		Code:      w.Code,
		Location:  w.Location,
		IsActive:  w.IsActive,
		CreatedAt: w.CreatedAt.Unix(),
	}
}

func ToWarehouseResponses(warehouses []Warehouse) []WarehouseResponse {
	responses := make([]WarehouseResponse, len(warehouses))
	for i, w := range warehouses {
		responses[i] = ToWarehouseResponse(&w)
	}
	return responses
}

func ToProductResponse(p *Product) ProductResponse {
	if p == nil {
		return ProductResponse{}
	}
	return ProductResponse{
		ID:            p.ID.String(),
		SKU:           p.SKU,
		Name:          p.Name,
		Category:      p.Category,
		Unit:          p.Unit,
		UnitPrice:     p.UnitPrice,
		MinStockLevel: p.MinStockLevel,
		Description:   p.Description,
		CreatedAt:     p.CreatedAt.Unix(),
	}
}

func ToProductResponses(products []Product) []ProductResponse {
	responses := make([]ProductResponse, len(products))
	for i, p := range products {
		responses[i] = ToProductResponse(&p)
	}
	return responses
}

func ToStockResponse(s *Stock) StockResponse {
	if s == nil {
		return StockResponse{}
	}
	res := StockResponse{
		ID:          s.ID.String(),
		WarehouseID: s.WarehouseID.String(),
		ProductID:   s.ProductID.String(),
		Quantity:    s.Quantity,
		CreatedAt:   s.CreatedAt.Unix(),
	}
	if s.Warehouse.Name != "" {
		res.WarehouseName = s.Warehouse.Name
	}
	if s.Product.Name != "" {
		res.ProductName = s.Product.Name
		res.ProductSKU = s.Product.SKU
	}
	return res
}

func ToStockResponses(stocks []Stock) []StockResponse {
	responses := make([]StockResponse, len(stocks))
	for i, s := range stocks {
		responses[i] = ToStockResponse(&s)
	}
	return responses
}
