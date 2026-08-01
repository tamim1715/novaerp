package department

type CreateDepartmentRequest struct {
	Name        string `json:"name" validate:"required,max=100" example:"hr"`
	Code        string `json:"code" validate:"required,max=20" example:"1234"`
	Description string `json:"description" example:"This is a hr department"`
}

type UpdateDepartmentRequest struct {
	Name        string `json:"name" validate:"required,max=100" example:"hr"`
	Code        string `json:"code" validate:"required,max=20" example:"1234"`
	Description string `json:"description" example:"This is an new hr department"`
}

type DepartmentResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}
