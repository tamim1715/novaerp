package employee

type CreateEmployeeRequest struct {
	EmployeeID string `json:"employeeId" binding:"required"`

	FirstName string `json:"firstName" binding:"required,max=100"`

	LastName string `json:"lastName"`

	Email string `json:"email" binding:"required,email"`

	Phone string `json:"phone"`

	Designation string `json:"designation"`

	JoiningDate int64 `json:"joiningDate"`

	Salary float64 `json:"salary"`

	Status bool `json:"status"`
}

type UpdateEmployeeRequest struct {
	EmployeeID string `json:"employeeId"`

	FirstName string `json:"firstName"`

	LastName string `json:"lastName"`

	Email string `json:"email"`

	Phone string `json:"phone"`

	Designation string `json:"designation"`

	JoiningDate int64 `json:"joiningDate"`

	Salary float64 `json:"salary"`

	Status bool `json:"status"`
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

	JoiningDate int64 `json:"joiningDate"`

	Salary float64 `json:"salary"`

	Status bool `json:"status"`

	CreatedAt int64 `json:"createdAt"`
}
