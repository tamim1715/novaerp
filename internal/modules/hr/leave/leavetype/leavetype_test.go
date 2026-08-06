package leavetype

import "testing"

func TestToLeaveTypeResponse(t *testing.T) {
	lt := &LeaveType{
		Name:           "Annual Leave",
		Code:           "AL",
		MaxDaysPerYear: 20,
		IsPaid:         true,
		Description:    "Paid annual vacation",
	}

	res := ToLeaveTypeResponse(lt)
	if res.Name != lt.Name {
		t.Errorf("expected Name %s, got %s", lt.Name, res.Name)
	}
	if res.Code != lt.Code {
		t.Errorf("expected Code %s, got %s", lt.Code, res.Code)
	}
	if res.MaxDaysPerYear != 20 {
		t.Errorf("expected MaxDaysPerYear 20, got %d", res.MaxDaysPerYear)
	}
}
