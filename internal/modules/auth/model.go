package auth

import (
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/model"
	"github.com/tamim1715/novaerp/internal/modules/user"
)

// RefreshToken stores issued refresh tokens for revocation & token rotation.
type RefreshToken struct {
	model.BaseModel

	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"userId"`
	Token     string    `gorm:"size:255;not null;uniqueIndex" json:"token"`
	ExpiresAt time.Time `gorm:"type:timestamp;not null" json:"expiresAt"`
	IsRevoked bool      `gorm:"default:false" json:"isRevoked"`
	UserAgent string    `gorm:"size:255" json:"userAgent"`
	ClientIP  string    `gorm:"size:45" json:"clientIp"`

	User user.User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
