package stock

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/modules/inventory/product"
	"github.com/tamim1715/novaerp/internal/modules/inventory/warehouse"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	FindAllStocks(ctx context.Context, req pagination.PageRequest) ([]Stock, int64, error)
	GetStock(ctx context.Context, warehouseID, productID string) (*Stock, error)
	AdjustStock(ctx context.Context, req AdjustStockRequest) (*Stock, error)
}

type service struct {
	repo          Repository
	warehouseRepo warehouse.Repository
	productRepo   product.Repository
	logger        *zap.Logger
}

func NewService(repo Repository, warehouseRepo warehouse.Repository, productRepo product.Repository, logger *zap.Logger) Service {
	return &service{
		repo:          repo,
		warehouseRepo: warehouseRepo,
		productRepo:   productRepo,
		logger:        logger,
	}
}

func (s *service) FindAllStocks(ctx context.Context, req pagination.PageRequest) ([]Stock, int64, error) {
	return s.repo.FindAll(ctx, req)
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
	if _, err := s.warehouseRepo.FindByID(ctx, req.WarehouseID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("warehouse not found")
		}
		return nil, err
	}

	if _, err := s.productRepo.FindByID(ctx, req.ProductID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return s.repo.AdjustStock(ctx, wID, pID, req.QuantityDelta)
}
