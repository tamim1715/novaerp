package user

import "github.com/tamim1715/novaerp/internal/common/model"

type User struct {
	model.BaseModel

	Username string `gorm:"size:50;uniqueIndex;not null"`
	Email    string `gorm:"size:150;uniqueIndex;not null"`
	Password string `gorm:"size:255;not null"`
	Role     string `gorm:"size:30;default:'user'"`
	IsActive bool   `gorm:"default:true"`
}
