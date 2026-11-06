package overtime

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestToOvertimePolicyResponse(t *testing.T) {
	p := &OvertimePolicy{
		Name:                   "Standard OT",
		Code:                   "OT_STD",
		MinOvertimeMinutes:     30,
		MaxOvertimeHoursPerDay: 4.0,
		WeekdayRateMultiplier:  1.5,
		WeekendRateMultiplier:  2.0,
		HolidayRateMultiplier:  2.0,
		IsDefault:              true,
		IsActive:               true,
	}
	p.ID = uuid.New()
	p.CreatedAt = time.Now()

	resp := ToOvertimePolicyResponse(p)
	if resp.Code != "OT_STD" {
		t.Errorf("expected Code OT_STD, got %s", resp.Code)
	}
	if resp.WeekdayRateMultiplier != 1.5 {
		t.Errorf("expected WeekdayRateMultiplier 1.5, got %f", resp.WeekdayRateMultiplier)
	}
	if !resp.IsDefault {
		t.Errorf("expected IsDefault true")
	}
}
