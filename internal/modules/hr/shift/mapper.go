package shift

func ToShiftResponse(s *Shift) ShiftResponse {
	return ShiftResponse{
		ID:                 s.ID.String(),
		Name:               s.Name,
		Code:               s.Code,
		StartTime:          s.StartTime,
		EndTime:            s.EndTime,
		GracePeriodMinutes: s.GracePeriodMinutes,
		WorkHours:          s.WorkHours,
		HalfDayHours:       s.HalfDayHours,
		IsNightShift:       s.IsNightShift,
		IsActive:           s.IsActive,
		Description:        s.Description,
		CreatedAt:          s.CreatedAt.Unix(),
	}
}

func ToShiftResponseList(list []Shift) []ShiftResponse {
	resp := make([]ShiftResponse, len(list))
	for i, item := range list {
		resp[i] = ToShiftResponse(&item)
	}
	return resp
}

func ToEmployeeShiftResponse(es *EmployeeShift) EmployeeShiftResponse {
	resp := EmployeeShiftResponse{
		ID:            es.ID.String(),
		EmployeeID:    es.EmployeeID.String(),
		ShiftID:       es.ShiftID.String(),
		EffectiveFrom: es.EffectiveFrom.Format("2006-01-02"),
		IsActive:      es.IsActive,
		CreatedAt:     es.CreatedAt.Unix(),
	}

	if es.EffectiveTo != nil {
		resp.EffectiveTo = es.EffectiveTo.Format("2006-01-02")
	}

	if es.Employee.FirstName != "" {
		resp.EmployeeName = es.Employee.FirstName + " " + es.Employee.LastName
	}

	if es.Shift.Name != "" {
		shResp := ToShiftResponse(&es.Shift)
		resp.Shift = &shResp
	}

	return resp
}

func ToEmployeeShiftResponseList(list []EmployeeShift) []EmployeeShiftResponse {
	resp := make([]EmployeeShiftResponse, len(list))
	for i, item := range list {
		resp[i] = ToEmployeeShiftResponse(&item)
	}
	return resp
}
