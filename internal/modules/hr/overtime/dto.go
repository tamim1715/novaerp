package overtime

type CreateOvertimePolicyRequest struct {
	Name                   string  `json:"name" binding:"required" example:"Standard Overtime Policy"`
	Code                   string  `json:"code" binding:"required" example:"OT_STD"`
	MinOvertimeMinutes     int     `json:"minOvertimeMinutes" example:"30"`
	MaxOvertimeHoursPerDay float64 `json:"maxOvertimeHoursPerDay" example:"4.0"`
	WeekdayRateMultiplier  float64 `json:"weekdayRateMultiplier" example:"1.5"`
	WeekendRateMultiplier  float64 `json:"weekendRateMultiplier" example:"2.0"`
	HolidayRateMultiplier  float64 `json:"holidayRateMultiplier" example:"2.0"`
	IsDefault              bool    `json:"isDefault" example:"true"`
	Description            string  `json:"description" example:"Standard corporate OT policy"`
}

type UpdateOvertimePolicyRequest struct {
	Name                   string   `json:"name" example:"Standard Overtime Policy"`
	MinOvertimeMinutes     *int     `json:"minOvertimeMinutes" example:"30"`
	MaxOvertimeHoursPerDay *float64 `json:"maxOvertimeHoursPerDay" example:"4.0"`
	WeekdayRateMultiplier  *float64 `json:"weekdayRateMultiplier" example:"1.5"`
	WeekendRateMultiplier  *float64 `json:"weekendRateMultiplier" example:"2.0"`
	HolidayRateMultiplier  *float64 `json:"holidayRateMultiplier" example:"2.0"`
	IsDefault              *bool    `json:"isDefault" example:"true"`
	IsActive               *bool    `json:"isActive" example:"true"`
	Description            string   `json:"description" example:"Updated OT policy"`
}

type OvertimePolicyResponse struct {
	ID                     string  `json:"id"`
	Name                   string  `json:"name"`
	Code                   string  `json:"code"`
	MinOvertimeMinutes     int     `json:"minOvertimeMinutes"`
	MaxOvertimeHoursPerDay float64 `json:"maxOvertimeHoursPerDay"`
	WeekdayRateMultiplier  float64 `json:"weekdayRateMultiplier"`
	WeekendRateMultiplier  float64 `json:"weekendRateMultiplier"`
	HolidayRateMultiplier  float64 `json:"holidayRateMultiplier"`
	IsDefault              bool    `json:"isDefault"`
	IsActive               bool    `json:"isActive"`
	Description            string  `json:"description"`
	CreatedAt              int64   `json:"createdAt"`
}
