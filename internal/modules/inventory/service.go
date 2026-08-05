package inventory

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	// Warehouse methods
	CreateWarehouse(ctx context.Context, req CreateWarehouseRequest) (*Warehouse, error)
	FindAllWarehouses(ctx context.Context, req pagination.PageRequest) ([]Warehouse, int64, error)
	FindWarehouseByID(ctx context.Context, id string) (*Warehouse, error)
	UpdateWarehouse(ctx context.Context, id string, req UpdateWarehouseRequest) (*Warehouse, error)
	DeleteWarehouse(ctx context.Context, id string) error

	// Product methods
	CreateProduct(ctx context.Context, req CreateProductRequest) (*Product, error)
	FindAllProducts(ctx context.Context, req pagination.PageRequest) ([]Product, int64, error)
	FindProductByID(ctx context.Context, id string) (*Product, error)
	UpdateProduct(ctx context.Context, id string, req UpdateProductRequest) (*Product, error)
	DeleteProduct(ctx context.Context, id string) error

	// Stock methods
	FindAllStocks(ctx context.Context, req pagination.PageRequest) ([]Stock, int64, error)
	GetStock(ctx context.Context, warehouseID, productID string) (*Stock, error)
	AdjustStock(ctx context.Context, req AdjustStockRequest) (*Stock, error)
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

// Warehouse implementations
func (s *service) CreateWarehouse(ctx context.Context, req CreateWarehouseRequest) (*Warehouse, error) {
	warehouse := &Warehouse{
		Name:     req.Name,
		Code:     req.Code,
		Location: req.Location,
		IsActive: true,
	}

	if err := s.repo.CreateWarehouse(ctx, warehouse); err != nil {
		return nil, fmt.Errorf("failed to create warehouse: %w", err)
	}

	return warehouse, nil
}

func (s *service) FindAllWarehouses(ctx context.Context, req pagination.PageRequest) ([]Warehouse, int64, error) {
	return s.repo.FindAllWarehouses(ctx, req)
}

func (s *service) FindWarehouseByID(ctx context.Context, id string) (*Warehouse, error) {
	return s.repo.FindWarehouseByID(ctx, id)
}

func (s *service) UpdateWarehouse(ctx context.Context, id string, req UpdateWarehouseRequest) (*Warehouse, error) {
	warehouse, err := s.repo.FindWarehouseByID(ctx, id)
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

	if err := s.repo.UpdateWarehouse(ctx, warehouse); err != nil {
		return nil, err
	}

	return warehouse, nil
}

func (s *service) DeleteWarehouse(ctx context.Context, id string) error {
	_, err := s.repo.FindWarehouseByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.DeleteWarehouse(ctx, id)
}

// Product implementations
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

	if err := s.repo.CreateProduct(ctx, product); err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}

func (s *service) FindAllProducts(ctx context.Context, req pagination.PageRequest) ([]Product, int64, error) {
	return s.repo.FindAllProducts(ctx, req)
}

func (s *service) FindProductByID(ctx context.Context, id string) (*Product, error) {
	return s.repo.FindProductByID(ctx, id)
}

func (s *service) UpdateProduct(ctx context.Context, id string, req UpdateProductRequest) (*Product, error) {
	product, err := s.repo.FindProductByID(ctx, id)
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

	if err := s.repo.UpdateProduct(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) DeleteProduct(ctx context.Context, id string) error {
	_, err := s.repo.FindProductByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.DeleteProduct(ctx, id)
}

// Stock implementations
func (s *service) FindAllStocks(ctx context.Context, req pagination.PageRequest) ([]Stock, int64, error) {
	return s.repo.FindAllStocks(ctx, req)
}

func (s *service) GetStock(ctx context.Context, warehouseID, productID string) (*Stock, error) {
	return s.repo.GetStock(ctx, warehouseID, productID)
}

func (s *service) AdjustStock(ctx context.Context, req AdjustStockRequest) (*Stock, error) {
	wID, err := uuid.Parse(req.WarehouseID)
	if err != nil {
		return nil, errors.New("invalid warehouse ID format")
	}

	pID, err := uuid.Parse(req.ProductID)
	if err != nil {
		return nil, errors.New("invalid product ID format")
	}

	// Verify warehouse and product exist
	if _, err := s.repo.FindWarehouseByID(ctx, req.WarehouseID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("warehouse not found")
		}
		return nil, err
	}

	if _, err := s.repo.FindProductByID(ctx, req.ProductID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return s.repo.AdjustStock(ctx, wID, pID, req.QuantityDelta)
}
