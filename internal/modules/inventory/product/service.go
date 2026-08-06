package product

import (
	"context"
	"fmt"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"go.uber.org/zap"
)

type Service interface {
	CreateProduct(ctx context.Context, req CreateProductRequest) (*Product, error)
	FindAllProducts(ctx context.Context, req pagination.PageRequest) ([]Product, int64, error)
	FindProductByID(ctx context.Context, id string) (*Product, error)
	UpdateProduct(ctx context.Context, id string, req UpdateProductRequest) (*Product, error)
	DeleteProduct(ctx context.Context, id string) error
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

func (s *service) CreateProduct(ctx context.Context, req CreateProductRequest) (*Product, error) {
	if req.Unit == "" {
		req.Unit = "pcs"
	}

	product := &Product{
		SKU:           req.SKU,
		Name:          req.Name,
		Category:      req.Category,
		Unit:          req.Unit,
		UnitPrice:     req.UnitPrice,
		MinStockLevel: req.MinStockLevel,
		Description:   req.Description,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}

func (s *service) FindAllProducts(ctx context.Context, req pagination.PageRequest) ([]Product, int64, error) {
	return s.repo.FindAll(ctx, req)
}

func (s *service) FindProductByID(ctx context.Context, id string) (*Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateProduct(ctx context.Context, id string, req UpdateProductRequest) (*Product, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.SKU != "" {
		product.SKU = req.SKU
	}
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Category != "" {
		product.Category = req.Category
	}
	if req.Unit != "" {
		product.Unit = req.Unit
	}
	if req.UnitPrice != 0 {
		product.UnitPrice = req.UnitPrice
	}
	if req.MinStockLevel != nil {
		product.MinStockLevel = *req.MinStockLevel
	}
	if req.Description != "" {
		product.Description = req.Description
	}

	if err := s.repo.Update(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) DeleteProduct(ctx context.Context, id string) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
