package warehouse

import (
	"context"
	"fmt"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"go.uber.org/zap"
)

type Service interface {
	CreateWarehouse(ctx context.Context, req CreateWarehouseRequest) (*Warehouse, error)
	FindAllWarehouses(ctx context.Context, req pagination.PageRequest) ([]Warehouse, int64, error)
	FindWarehouseByID(ctx context.Context, id string) (*Warehouse, error)
	UpdateWarehouse(ctx context.Context, id string, req UpdateWarehouseRequest) (*Warehouse, error)
	DeleteWarehouse(ctx context.Context, id string) error
}

type service struct {
	repo   Repository
	logger *zap.Logger
}

func NewService(repo Repository, logger *zap.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s *service) CreateWarehouse(ctx context.Context, req CreateWarehouseRequest) (*Warehouse, error) {
	warehouse := &Warehouse{
		Name:     req.Name,
		Code:     req.Code,
		Location: req.Location,
		IsActive: true,
	}

	if err := s.repo.Create(ctx, warehouse); err != nil {
		return nil, fmt.Errorf("failed to create warehouse: %w", err)
	}

	return warehouse, nil
}

func (s *service) FindAllWarehouses(ctx context.Context, req pagination.PageRequest) ([]Warehouse, int64, error) {
	return s.repo.FindAll(ctx, req)
}

func (s *service) FindWarehouseByID(ctx context.Context, id string) (*Warehouse, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateWarehouse(ctx context.Context, id string, req UpdateWarehouseRequest) (*Warehouse, error) {
	warehouse, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		warehouse.Name = req.Name
	}
	if req.Code != "" {
		warehouse.Code = req.Code
	}
	if req.Location != "" {
		warehouse.Location = req.Location
	}
	if req.IsActive != nil {
		warehouse.IsActive = *req.IsActive
	}

	if err := s.repo.Update(ctx, warehouse); err != nil {
		return nil, err
	}

	return warehouse, nil
}

func (s *service) DeleteWarehouse(ctx context.Context, id string) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
