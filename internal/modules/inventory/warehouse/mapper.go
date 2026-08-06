package warehouse

func ToWarehouseResponse(w *Warehouse) WarehouseResponse {
	return WarehouseResponse{
		ID:        w.ID.String(),
		Name:      w.Name,
		Code:      w.Code,
		Location:  w.Location,
		IsActive:  w.IsActive,
		CreatedAt: w.CreatedAt.Unix(),
	}
}

func ToWarehouseResponseList(warehouses []Warehouse) []WarehouseResponse {
	list := make([]WarehouseResponse, len(warehouses))
	for i, w := range warehouses {
		list[i] = ToWarehouseResponse(&w)
	}
	return list
}
