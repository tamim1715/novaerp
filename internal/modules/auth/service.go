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

type Service interface {
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
}

type service struct {
	userRepo  user.Repository
	jwtSecret string
	logger    *zap.Logger
}

func NewService(userRepo user.Repository, jwtSecret string, logger *zap.Logger) Service {
	return &service{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		logger:    logger,
	}
}

func (s *service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
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

	token, err := GenerateToken(u.ID.String(), u.Email, s.jwtSecret, 24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &LoginResponse{
		Token: token,
		User:  user.ToResponse(u),
	}, nil
}
