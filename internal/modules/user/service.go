package user

import (
	"errors"

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
	Create(request CreateUserRequest) (*User, error)
	FindAll(request pagination.PageRequest) ([]User, int64, error)
	FindByID(string) (*User, error)
	Update(string, UpdateUserRequest) (*User, error)
	Delete(string) error
}

type service struct {
	repository Repository
	logger     *zap.Logger
}

func (s *service) Create(req CreateUserRequest) (*User, error) {

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashPassword(req.Password),
		IsActive: true,
	}

	if err := s.repository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) FindAll(request pagination.PageRequest) ([]User, int64, error) {
	return s.repository.FindAll(request)
}

func (s *service) FindByID(id string) (*User, error) {
	return s.repository.FindByID(id)
}

func (s *service) Update(id string, req UpdateUserRequest) (*User, error) {

	user, err := s.repository.FindByID(id)

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

	if err := s.repository.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) Delete(id string) error {

	_, err := s.repository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	return s.repository.Delete(id)
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}
