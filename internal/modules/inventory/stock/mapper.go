package stock

func ToStockResponse(s *Stock) StockResponse {
	resp := StockResponse{
		ID:          s.ID.String(),
		WarehouseID: s.WarehouseID.String(),
		ProductID:   s.ProductID.String(),
		Quantity:    s.Quantity,
		CreatedAt:   s.CreatedAt.Unix(),
	}

	if s.Warehouse.Name != "" {
		resp.WarehouseName = s.Warehouse.Name
	}
	if s.Product.Name != "" {
		resp.ProductName = s.Product.Name
		resp.ProductSKU = s.Product.SKU
	}

	return resp
}

func ToStockResponseList(stocks []Stock) []StockResponse {
	list := make([]StockResponse, len(stocks))
	for i, s := range stocks {
		list[i] = ToStockResponse(&s)
	}
	return list
}
