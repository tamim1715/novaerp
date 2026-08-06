package product

func ToProductResponse(p *Product) ProductResponse {
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

func ToProductResponseList(products []Product) []ProductResponse {
	list := make([]ProductResponse, len(products))
	for i, p := range products {
		list[i] = ToProductResponse(&p)
	}
	return list
}
