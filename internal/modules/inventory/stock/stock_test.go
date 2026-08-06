package stock

import (
	"testing"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/modules/inventory/product"
	"github.com/tamim1715/novaerp/internal/modules/inventory/warehouse"
)

func TestToStockResponse(t *testing.T) {
	wID := uuid.New()
	pID := uuid.New()
	s := &Stock{
		WarehouseID: wID,
		ProductID:   pID,
		Quantity:    100,
		Warehouse:   warehouse.Warehouse{Name: "Main Warehouse"},
		Product:     product.Product{Name: "Laptop", SKU: "LTP-101"},
	}

	res := ToStockResponse(s)
	if res.Quantity != 100 {
		t.Errorf("expected Quantity 100, got %d", res.Quantity)
	}
	if res.WarehouseName != "Main Warehouse" {
		t.Errorf("expected WarehouseName Main Warehouse, got %s", res.WarehouseName)
	}
	if res.ProductName != "Laptop" {
		t.Errorf("expected ProductName Laptop, got %s", res.ProductName)
	}
}
