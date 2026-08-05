package department

import (
	"context"
	"errors"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"go.uber.org/zap"
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
	Create(ctx context.Context, req CreateDepartmentRequest) (*Department, error)
	FindAll(ctx context.Context, request pagination.PageRequest) ([]Department, int64, error)
	FindByID(ctx context.Context, id string) (*Department, error)
	Update(ctx context.Context, id string, req UpdateDepartmentRequest) (*Department, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository Repository
	logger     *zap.Logger
}

func (s *service) Create(ctx context.Context, req CreateDepartmentRequest) (*Department, error) {

	department := &Department{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	}

	if err := s.repository.Create(ctx, department); err != nil {
		return nil, err
	}

	return department, nil
}

func (s *service) FindAll(ctx context.Context, request pagination.PageRequest) ([]Department, int64, error) {
	return s.repository.FindAll(ctx, request)
}

func (s *service) FindByID(ctx context.Context, id string) (*Department, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, id string, req UpdateDepartmentRequest) (*Department, error) {

	department, err := s.repository.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	department.Name = req.Name
	department.Code = req.Code
	department.Description = req.Description

	if err := s.repository.Update(ctx, department); err != nil {
		return nil, err
	}

	return department, nil
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
