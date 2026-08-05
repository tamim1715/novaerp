// Package model defines base models shared across domain entities.
package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel provides common fields for GORM entities.
type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hooks into GORM's creation process to ensure a valid UUID is assigned.
func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}
