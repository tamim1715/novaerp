package pagination

type PageRequest struct {
	Page   int    `form:"page"`
	Size   int    `form:"size"`
	Search string `form:"search"`
	SortBy string `form:"sortBy"`
	Order  string `form:"order"`
}

type PageResponse struct {
	Page       int   `json:"page"`
	Size       int   `json:"size"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int   `json:"totalPages"`
	Data       any   `json:"data"`
}

func (p *PageRequest) Normalize() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Size <= 0 {
		p.Size = 10
	}

	if p.Size > 100 {
		p.Size = 100
	}

	if p.SortBy == "" {
		p.SortBy = "created_at"
	}

	if p.Order != "asc" && p.Order != "desc" {
		p.Order = "desc"
	}
}

func (p *PageRequest) Offset() int {
	return (p.Page - 1) * p.Size
}
