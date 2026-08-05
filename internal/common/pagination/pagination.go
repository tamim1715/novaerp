// Package pagination provides helpers for handling list pagination requests and responses.
package pagination

import "regexp"

var safeSortRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

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

	// Ensure SortBy contains only valid identifier characters to prevent SQL injection
	if p.SortBy == "" || !safeSortRegex.MatchString(p.SortBy) {
		p.SortBy = "created_at"
	}

	if p.Order != "asc" && p.Order != "desc" {
		p.Order = "desc"
	}
}

func (p *PageRequest) Offset() int {
	return (p.Page - 1) * p.Size
}
