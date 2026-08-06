package auth

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRefreshToken(ctx context.Context, token *RefreshToken) error
	FindRefreshToken(ctx context.Context, tokenStr string) (*RefreshToken, error)
	RevokeRefreshToken(ctx context.Context, tokenStr string) error
	RevokeAllUserRefreshTokens(ctx context.Context, userID uuid.UUID) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateRefreshToken(ctx context.Context, token *RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *repository) FindRefreshToken(ctx context.Context, tokenStr string) (*RefreshToken, error) {
	var token RefreshToken
	err := r.db.WithContext(ctx).
		Preload("User").
		First(&token, "token = ? AND is_revoked = false", tokenStr).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *repository) RevokeRefreshToken(ctx context.Context, tokenStr string) error {
	return r.db.WithContext(ctx).
		Model(&RefreshToken{}).
		Where("token = ?", tokenStr).
		Update("is_revoked", true).Error
}

func (r *repository) RevokeAllUserRefreshTokens(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&RefreshToken{}).
		Where("user_id = ?", userID).
		Update("is_revoked", true).Error
}
