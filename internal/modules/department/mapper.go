package department

func ToResponse(department *Department) DepartmentResponse {
	return DepartmentResponse{
		ID:          department.ID.String(),
		Name:        department.Name,
		Code:        department.Code,
		Description: department.Description,
		CreatedAt:   department.CreatedAt.Unix(),
	}
}

func ToResponses(departments []Department) []DepartmentResponse {

	responses := make([]DepartmentResponse, 0, len(departments))

	for _, department := range departments {
		responses = append(responses, ToResponse(&department))
	}

	return responses
}
