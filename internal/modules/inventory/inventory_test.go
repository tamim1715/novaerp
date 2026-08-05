package inventory

import (
	"testing"
)

func TestToWarehouseResponse(t *testing.T) {
	w := &Warehouse{
		Name:     "Central Warehouse",
		Code:     "WH-CENTRAL",
		Location: "Dhaka",
		IsActive: true,
	}

	res := ToWarehouseResponse(w)
	if res.Name != w.Name {
		t.Errorf("expected Name %s, got %s", w.Name, res.Name)
	}
	if res.Code != w.Code {
		t.Errorf("expected Code %s, got %s", w.Code, res.Code)
	}
	if res.Location != w.Location {
		t.Errorf("expected Location %s, got %s", w.Location, res.Location)
	}
	if !res.IsActive {
		t.Error("expected IsActive to be true")
	}
}

func TestToProductResponse(t *testing.T) {
	p := &Product{
		SKU:           "SKU-100",
		Name:          "Gaming Laptop",
		Category:      "Electronics",
		Unit:          "pcs",
		UnitPrice:     1200.50,
		MinStockLevel: 5,
		Description:   "High-end gaming laptop",
	}

	res := ToProductResponse(p)
	if res.SKU != p.SKU {
		t.Errorf("expected SKU %s, got %s", p.SKU, res.SKU)
	}
	if res.Name != p.Name {
		t.Errorf("expected Name %s, got %s", p.Name, res.Name)
	}
	if res.UnitPrice != p.UnitPrice {
		t.Errorf("expected UnitPrice %f, got %f", p.UnitPrice, res.UnitPrice)
	}
}
