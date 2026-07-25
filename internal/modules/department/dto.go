package department

type CreateDepartmentRequest struct {
	Name        string `json:"name" validate:"required,max=100"`
	Code        string `json:"code" validate:"required,max=20"`
	Description string `json:"description"`
}

type UpdateDepartmentRequest struct {
	Name        string `json:"name" validate:"required,max=100"`
	Code        string `json:"code" validate:"required,max=20"`
	Description string `json:"description"`
}

type DepartmentResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}
