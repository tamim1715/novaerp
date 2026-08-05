package pagination

import (
	"testing"
)

func TestPageRequest_Normalize(t *testing.T) {
	tests := []struct {
		name     string
		input    PageRequest
		expected PageRequest
	}{
		{
			name:  "Zero or negative values get defaults",
			input: PageRequest{Page: 0, Size: -5, SortBy: "", Order: ""},
			expected: PageRequest{
				Page:   1,
				Size:   10,
				SortBy: "created_at",
				Order:  "desc",
			},
		},
		{
			name:  "Size limit capped at 100",
			input: PageRequest{Page: 2, Size: 150, SortBy: "name", Order: "asc"},
			expected: PageRequest{
				Page:   2,
				Size:   100,
				SortBy: "name",
				Order:  "asc",
			},
		},
		{
			name:  "Invalid SortBy SQL injection attempt resets to default",
			input: PageRequest{Page: 1, Size: 10, SortBy: "id; DROP TABLE users;--", Order: "asc"},
			expected: PageRequest{
				Page:   1,
				Size:   10,
				SortBy: "created_at",
				Order:  "asc",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.input
			req.Normalize()

			if req.Page != tt.expected.Page {
				t.Errorf("expected Page %d, got %d", tt.expected.Page, req.Page)
			}
			if req.Size != tt.expected.Size {
				t.Errorf("expected Size %d, got %d", tt.expected.Size, req.Size)
			}
			if req.SortBy != tt.expected.SortBy {
				t.Errorf("expected SortBy %q, got %q", tt.expected.SortBy, req.SortBy)
			}
			if req.Order != tt.expected.Order {
				t.Errorf("expected Order %q, got %q", tt.expected.Order, req.Order)
			}
		})
	}
}

func TestPageRequest_Offset(t *testing.T) {
	req := PageRequest{Page: 3, Size: 20}
	expectedOffset := 40
	if req.Offset() != expectedOffset {
		t.Errorf("expected offset %d, got %d", expectedOffset, req.Offset())
	}
}
