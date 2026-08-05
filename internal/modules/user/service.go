package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func NewService(
	repository Repository,
	logger *zap.Logger,
) Service {

	return &service{
		repository: repository,
		logger:     logger,
	}
}

type Service interface {
	Create(ctx context.Context, request CreateUserRequest) (*User, error)
	FindAll(ctx context.Context, request pagination.PageRequest) ([]User, int64, error)
	FindByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, id string, req UpdateUserRequest) (*User, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository Repository
	logger     *zap.Logger
}

func (s *service) Create(ctx context.Context, req CreateUserRequest) (*User, error) {

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		IsActive: true,
	}

	if err := s.repository.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) FindAll(ctx context.Context, request pagination.PageRequest) ([]User, int64, error) {
	return s.repository.FindAll(ctx, request)
}

func (s *service) FindByID(ctx context.Context, id string) (*User, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, id string, req UpdateUserRequest) (*User, error) {

	user, err := s.repository.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.repository.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) Delete(ctx context.Context, id string) error {

	_, err := s.repository.FindByID(ctx, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	return s.repository.Delete(ctx, id)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
