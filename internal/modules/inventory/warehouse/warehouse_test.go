package warehouse

import "testing"

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
