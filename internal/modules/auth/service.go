package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tamim1715/novaerp/internal/modules/user"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	AccessTokenDuration  = 15 * time.Minute
	RefreshTokenDuration = 7 * 24 * time.Hour
)

type Service interface {
	Login(ctx context.Context, req LoginRequest, clientIP, userAgent string) (*LoginResponse, error)
	RefreshToken(ctx context.Context, req RefreshTokenRequest, clientIP, userAgent string) (*RefreshTokenResponse, error)
	Logout(ctx context.Context, req LogoutRequest) error
}

type service struct {
	authRepo   Repository
	userRepo   user.Repository
	keyManager *KeyManager
	logger     *zap.Logger
}

func NewService(authRepo Repository, userRepo user.Repository, keyManager *KeyManager, logger *zap.Logger) Service {
	return &service{
		authRepo:   authRepo,
		userRepo:   userRepo,
		keyManager: keyManager,
		logger:     logger,
	}
}

func (s *service) Login(ctx context.Context, req LoginRequest, clientIP, userAgent string) (*LoginResponse, error) {
	u, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !u.IsActive {
		return nil, errors.New("user account is inactive")
	}

	// Generate RS256 Access Token
	accessToken, err := GenerateAccessToken(u.ID.String(), u.Email, u.Role, s.keyManager.PrivateKey, AccessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// Generate Refresh Token string
	refreshTokenStr, err := GenerateRefreshTokenString()
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// Save Refresh Token to database
	rfToken := &RefreshToken{
		UserID:    u.ID,
		Token:     refreshTokenStr,
		ExpiresAt: time.Now().UTC().Add(RefreshTokenDuration),
		IsRevoked: false,
		UserAgent: userAgent,
		ClientIP:  clientIP,
	}

	if err := s.authRepo.CreateRefreshToken(ctx, rfToken); err != nil {
		return nil, fmt.Errorf("failed to persist refresh token: %w", err)
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
		TokenType:    "Bearer",
		ExpiresIn:    int64(AccessTokenDuration.Seconds()),
		User:         user.ToResponse(u),
	}, nil
}

func (s *service) RefreshToken(ctx context.Context, req RefreshTokenRequest, clientIP, userAgent string) (*RefreshTokenResponse, error) {
	rfToken, err := s.authRepo.FindRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid or revoked refresh token")
		}
		return nil, err
	}

	if rfToken.IsRevoked {
		return nil, errors.New("refresh token has been revoked")
	}

	if time.Now().UTC().After(rfToken.ExpiresAt) {
		_ = s.authRepo.RevokeRefreshToken(ctx, req.RefreshToken)
		return nil, errors.New("refresh token has expired")
	}

	if !rfToken.User.IsActive {
		return nil, errors.New("user account is inactive")
	}

	// Revoke old refresh token (Token Rotation for security)
	if err := s.authRepo.RevokeRefreshToken(ctx, req.RefreshToken); err != nil {
		s.logger.Error("failed to revoke old refresh token", zap.Error(err))
	}

	// Generate new Access Token
	newAccessToken, err := GenerateAccessToken(rfToken.User.ID.String(), rfToken.User.Email, rfToken.User.Role, s.keyManager.PrivateKey, AccessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to generate new access token: %w", err)
	}

	// Generate new rotated Refresh Token
	newRefreshTokenStr, err := GenerateRefreshTokenString()
	if err != nil {
		return nil, fmt.Errorf("failed to generate new refresh token: %w", err)
	}

	newRfToken := &RefreshToken{
		UserID:    rfToken.User.ID,
		Token:     newRefreshTokenStr,
		ExpiresAt: time.Now().UTC().Add(RefreshTokenDuration),
		IsRevoked: false,
		UserAgent: userAgent,
		ClientIP:  clientIP,
	}

	if err := s.authRepo.CreateRefreshToken(ctx, newRfToken); err != nil {
		return nil, fmt.Errorf("failed to persist new refresh token: %w", err)
	}

	return &RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshTokenStr,
		TokenType:    "Bearer",
		ExpiresIn:    int64(AccessTokenDuration.Seconds()),
	}, nil
}

func (s *service) Logout(ctx context.Context, req LogoutRequest) error {
	return s.authRepo.RevokeRefreshToken(ctx, req.RefreshToken)
}
