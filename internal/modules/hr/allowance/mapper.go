package allowance

func ToAllowanceTypeResponse(at *AllowanceType) AllowanceTypeResponse {
	return AllowanceTypeResponse{
		ID:              at.ID.String(),
		Name:            at.Name,
		Code:            at.Code,
		CalculationType: string(at.CalculationType),
		DefaultAmount:   at.DefaultAmount,
		IsTaxable:       at.IsTaxable,
		IsActive:        at.IsActive,
		Description:     at.Description,
		CreatedAt:       at.CreatedAt.Unix(),
	}
}

func ToAllowanceTypeResponseList(list []AllowanceType) []AllowanceTypeResponse {
	resp := make([]AllowanceTypeResponse, len(list))
	for i, item := range list {
		resp[i] = ToAllowanceTypeResponse(&item)
	}
	return resp
}

func ToEmployeeAllowanceResponse(ea *EmployeeAllowance) EmployeeAllowanceResponse {
	resp := EmployeeAllowanceResponse{
		ID:              ea.ID.String(),
		EmployeeID:      ea.EmployeeID.String(),
		AllowanceTypeID: ea.AllowanceTypeID.String(),
		Amount:          ea.Amount,
		IsActive:        ea.IsActive,
		CreatedAt:       ea.CreatedAt.Unix(),
	}

	if ea.Employee.FirstName != "" {
		resp.EmployeeName = ea.Employee.FirstName + " " + ea.Employee.LastName
	}

	if ea.AllowanceType.Name != "" {
		atResp := ToAllowanceTypeResponse(&ea.AllowanceType)
		resp.AllowanceType = &atResp
	}

	return resp
}

func ToEmployeeAllowanceResponseList(list []EmployeeAllowance) []EmployeeAllowanceResponse {
	resp := make([]EmployeeAllowanceResponse, len(list))
	for i, item := range list {
		resp[i] = ToEmployeeAllowanceResponse(&item)
	}
	return resp
}
