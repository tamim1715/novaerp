package overtime

func ToOvertimePolicyResponse(p *OvertimePolicy) OvertimePolicyResponse {
	return OvertimePolicyResponse{
		ID:                     p.ID.String(),
		Name:                   p.Name,
		Code:                   p.Code,
		MinOvertimeMinutes:     p.MinOvertimeMinutes,
		MaxOvertimeHoursPerDay: p.MaxOvertimeHoursPerDay,
		WeekdayRateMultiplier:  p.WeekdayRateMultiplier,
		WeekendRateMultiplier:  p.WeekendRateMultiplier,
		HolidayRateMultiplier:  p.HolidayRateMultiplier,
		IsDefault:              p.IsDefault,
		IsActive:               p.IsActive,
		Description:            p.Description,
		CreatedAt:              p.CreatedAt.Unix(),
	}
}

func ToOvertimePolicyResponseList(list []OvertimePolicy) []OvertimePolicyResponse {
	resp := make([]OvertimePolicyResponse, len(list))
	for i, item := range list {
		resp[i] = ToOvertimePolicyResponse(&item)
	}
	return resp
}
