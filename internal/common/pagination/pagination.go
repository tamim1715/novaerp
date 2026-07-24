package pagination

type PageRequest struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

type PageResponse struct {
	Page       int         `json:"page"`
	Size       int         `json:"size"`
	TotalItems int64       `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
	Data       interface{} `json:"data"`
}
