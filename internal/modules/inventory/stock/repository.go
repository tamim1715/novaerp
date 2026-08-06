package stock

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(ctx context.Context, req pagination.PageRequest) ([]Stock, int64, error)
	GetStock(ctx context.Context, warehouseID, productID string) (*Stock, error)
	AdjustStock(ctx context.Context, warehouseID, productID uuid.UUID, delta int) (*Stock, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(ctx context.Context, req pagination.PageRequest) ([]Stock, int64, error) {
	var stocks []Stock
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&Stock{}).
		Preload("Warehouse").
		Preload("Product")

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := req.Offset()
	err := db.Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(offset).
		Find(&stocks).Error

	return stocks, total, err
}

func (r *repository) GetStock(ctx context.Context, warehouseID, productID string) (*Stock, error) {
	wID, err := uuid.Parse(warehouseID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	pID, err := uuid.Parse(productID)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	var stock Stock
	err = r.db.WithContext(ctx).
		Preload("Warehouse").
		Preload("Product").
		First(&stock, "warehouse_id = ? AND product_id = ?", wID, pID).Error
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *repository) AdjustStock(ctx context.Context, warehouseID, productID uuid.UUID, delta int) (*Stock, error) {
	var stock Stock

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("warehouse_id = ? AND product_id = ?", warehouseID, productID).
			First(&stock).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if delta < 0 {
					return errors.New("cannot reduce stock below zero for new record")
				}
				stock = Stock{
					WarehouseID: warehouseID,
					ProductID:   productID,
					Quantity:    delta,
				}
				return tx.Create(&stock).Error
			}
			return err
		}

		newQty := stock.Quantity + delta
		if newQty < 0 {
			return errors.New("insufficient stock quantity")
		}

		stock.Quantity = newQty
		return tx.Save(&stock).Error
	})

	if err != nil {
		return nil, err
	}

	// Reload associations for clean response
	r.db.WithContext(ctx).Preload("Warehouse").Preload("Product").First(&stock, "id = ?", stock.ID)

	return &stock, nil
}
