package department

import "github.com/tamim1715/novaerp/internal/common/model"

type Department struct {
	model.BaseModel

	Name        string `gorm:"size:100;not null;unique" json:"name"`
	Code        string `gorm:"size:20;not null;unique" json:"code"`
	Description string `gorm:"size:255" json:"description"`
}
