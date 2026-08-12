package employee

func ToResponse(employee *Employee) EmployeeResponse {
	return EmployeeResponse{
		ID:           employee.ID.String(),
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		Phone:        employee.Phone,
		Email:        employee.Email,
		Designation:  employee.Designation,
		JoiningDate:  employee.JoiningDate.String(),
		Salary:       employee.Salary,
		Status:       employee.Status.String(),
		DepartmentID: employee.DepartmentID,
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
