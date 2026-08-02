package employee

func ToResponse(employee *Employee) EmployeeResponse {
	return EmployeeResponse{
		ID:           employee.ID.String(),
		Email:        employee.Email,
		DepartmentID: employee.DepartmentID.String(),
		CreatedAt:    employee.CreatedAt.Unix(),
	}
}

func ToResponses(employees []Employee) []EmployeeResponse {

	responses := make([]EmployeeResponse, 0, len(employees))

	for _, employee := range employees {
		responses = append(responses, ToResponse(&employee))
	}
	return responses
}
