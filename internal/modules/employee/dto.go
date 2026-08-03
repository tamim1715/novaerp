package employee

type CreateEmployeeRequest struct {
	DepartmentID string `json:"department_id" binding:"required" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`

	FirstName string `json:"firstName" binding:"required,max=100" example:"Shahadath Hossain"`

	LastName string `json:"lastName" example:"Tamim"`

	Email string `json:"email" binding:"required,email" example:"shahadat@gmail.com"`

	Phone string `json:"phone" example:"123456"`

	Designation string `json:"designation" example:"assistant manager"`

	JoiningDate string `json:"joiningDate" example:"2006-01-02"`

	Salary float64 `json:"salary" example:"50000"`

	Status string `json:"status" example:"active"`
}

type UpdateEmployeeRequest struct {
	DepartmentID string `json:"department_id" example:"3f393bbd-9790-4045-853e-26a2c732ee06"`

	FirstName string `json:"firstName" example:"Shahadath Hossain"`

	LastName string `json:"lastName" example:"Tamim"`

	Email string `json:"email" example:"shahadat@gmail.com"`

	Phone string `json:"phone" example:"123456"`

	Designation string `json:"designation" example:"assistant manager"`

	JoiningDate string `json:"joiningDate" example:"2006-01-02"`

	Salary float64 `json:"salary" example:"50000"`

	Status string `json:"status" example:"active"`
}

type EmployeeResponse struct {
	ID string `json:"id"`

	DepartmentID string `json:"departmentId"`

	DepartmentName string `json:"departmentName"`

	FirstName string `json:"firstName"`

	LastName string `json:"lastName"`

	Email string `json:"email"`

	Phone string `json:"phone"`

	Designation string `json:"designation"`

	JoiningDate string `json:"joiningDate"`

	Salary float64 `json:"salary"`

	Status string `json:"status"`

	CreatedAt int64 `json:"createdAt"`
}
