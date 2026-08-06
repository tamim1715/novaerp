package leavetype

type CreateLeaveTypeRequest struct {
	Name           string `json:"name" binding:"required,max=100" example:"Casual Leave"`
	Code           string `json:"code" binding:"required,max=20" example:"CL"`
	MaxDaysPerYear int    `json:"maxDaysPerYear" binding:"required,min=1" example:"14"`
	IsPaid         *bool  `json:"isPaid" example:"true"`
	Description    string `json:"description" example:"Standard annual casual leave"`
}

type UpdateLeaveTypeRequest struct {
	Name           string `json:"name" example:"Casual Leave"`
	Code           string `json:"code" example:"CL"`
	MaxDaysPerYear *int   `json:"maxDaysPerYear" example:"15"`
	IsPaid         *bool  `json:"isPaid" example:"true"`
	Description    string `json:"description" example:"Updated casual leave policy"`
}

type LeaveTypeResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Code           string `json:"code"`
	MaxDaysPerYear int    `json:"maxDaysPerYear"`
	IsPaid         bool   `json:"isPaid"`
	Description    string `json:"description"`
	CreatedAt      int64  `json:"createdAt"`
}
