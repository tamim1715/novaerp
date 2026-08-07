package account

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	Code          string     `json:"code" binding:"required,min=2,max=50"`
	Name          string     `json:"name" binding:"required,min=2,max=150"`
	Type          string     `json:"type" binding:"required,oneof=ASSET LIABILITY EQUITY REVENUE EXPENSE"`
	NormalBalance string     `json:"normalBalance" binding:"omitempty,oneof=DEBIT CREDIT"`
	ParentID      *uuid.UUID `json:"parentId,omitempty"`
	Currency      string     `json:"currency" binding:"omitempty,len=3"`
	IsActive      *bool      `json:"isActive,omitempty"`
	Description   string     `json:"description" binding:"omitempty,max=255"`
}

type UpdateAccountRequest struct {
	Name        *string    `json:"name,omitempty" binding:"omitempty,min=2,max=150"`
	ParentID    *uuid.UUID `json:"parentId,omitempty"`
	Currency    *string    `json:"currency,omitempty" binding:"omitempty,len=3"`
	IsActive    *bool      `json:"isActive,omitempty"`
	Description *string    `json:"description,omitempty" binding:"omitempty,max=255"`
}

type AccountResponse struct {
	ID            uuid.UUID  `json:"id"`
	Code          string     `json:"code"`
	Name          string     `json:"name"`
	Type          string     `json:"type"`
	NormalBalance string     `json:"normalBalance"`
	ParentID      *uuid.UUID `json:"parentId,omitempty"`
	Currency      string     `json:"currency"`
	IsActive      bool       `json:"isActive"`
	IsSystem      bool       `json:"isSystem"`
	Description   string     `json:"description"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
}

type AccountTreeResponse struct {
	AccountResponse
	Children []AccountTreeResponse `json:"children,omitempty"`
}
