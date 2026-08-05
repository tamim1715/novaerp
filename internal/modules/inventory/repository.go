package inventory

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrNilDB = errors.New("database connection is not initialized")

type Repository interface {
	// Warehouse operations
	CreateWarehouse(ctx context.Context, warehouse *Warehouse) error
	FindAllWarehouses(ctx context.Context, req pagination.PageRequest) ([]Warehouse, int64, error)
	FindWarehouseByID(ctx context.Context, id string) (*Warehouse, error)
	UpdateWarehouse(ctx context.Context, warehouse *Warehouse) error
	DeleteWarehouse(ctx context.Context, id string) error

	// Product operations
	CreateProduct(ctx context.Context, product *Product) error
	FindAllProducts(ctx context.Context, req pagination.PageRequest) ([]Product, int64, error)
	FindProductByID(ctx context.Context, id string) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id string) error

	// Stock operations
	GetStock(ctx context.Context, warehouseID, productID string) (*Stock, error)
	FindAllStocks(ctx context.Context, req pagination.PageRequest) ([]Stock, int64, error)
	AdjustStock(ctx context.Context, warehouseID, productID uuid.UUID, delta int) (*Stock, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Warehouse implementations
func (r *repository) CreateWarehouse(ctx context.Context, warehouse *Warehouse) error {
	if r == nil || r.db == nil {
		return ErrNilDB
	}
	return r.db.WithContext(ctx).Create(warehouse).Error
}

func (r *repository) FindAllWarehouses(ctx context.Context, req pagination.PageRequest) ([]Warehouse, int64, error) {
	if r == nil || r.db == nil {
		return nil, 0, ErrNilDB
	}
	req.Normalize()

	var warehouses []Warehouse
	var total int64

	query := r.db.WithContext(ctx).Model(&Warehouse{})
	if req.Search != "" {
		search := "%" + req.Search + "%"
		query = query.Where("name ILIKE ? OR code ILIKE ? OR location ILIKE ?", search, search, search)
	}

	query.Count(&total)

	err := query.
		Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(req.Offset()).
		Find(&warehouses).Error

	return warehouses, total, err
}

func (r *repository) FindWarehouseByID(ctx context.Context, id string) (*Warehouse, error) {
	if r == nil || r.db == nil {
		return nil, ErrNilDB
	}
	var warehouse Warehouse
	err := r.db.WithContext(ctx).First(&warehouse, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *repository) UpdateWarehouse(ctx context.Context, warehouse *Warehouse) error {
	if r == nil || r.db == nil {
		return ErrNilDB
	}
	return r.db.WithContext(ctx).Save(warehouse).Error
}

func (r *repository) DeleteWarehouse(ctx context.Context, id string) error {
	if r == nil || r.db == nil {
		return ErrNilDB
	}
	return r.db.WithContext(ctx).Delete(&Warehouse{}, "id = ?", id).Error
}

// Product implementations
func (r *repository) CreateProduct(ctx context.Context, product *Product) error {
	if r == nil || r.db == nil {
		return ErrNilDB
	}
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *repository) FindAllProducts(ctx context.Context, req pagination.PageRequest) ([]Product, int64, error) {
	if r == nil || r.db == nil {
		return nil, 0, ErrNilDB
	}
	req.Normalize()

	var products []Product
	var total int64

	query := r.db.WithContext(ctx).Model(&Product{})
	if req.Search != "" {
		search := "%" + req.Search + "%"
		query = query.Where("sku ILIKE ? OR name ILIKE ? OR category ILIKE ?", search, search, search)
	}

	query.Count(&total)

	err := query.
		Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(req.Offset()).
		Find(&products).Error

	return products, total, err
}

func (r *repository) FindProductByID(ctx context.Context, id string) (*Product, error) {
	if r == nil || r.db == nil {
		return nil, ErrNilDB
	}
	var product Product
	err := r.db.WithContext(ctx).First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *repository) UpdateProduct(ctx context.Context, product *Product) error {
	if r == nil || r.db == nil {
		return ErrNilDB
	}
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *repository) DeleteProduct(ctx context.Context, id string) error {
	if r == nil || r.db == nil {
		return ErrNilDB
	}
	return r.db.WithContext(ctx).Delete(&Product{}, "id = ?", id).Error
}

// Stock implementations
func (r *repository) GetStock(ctx context.Context, warehouseID, productID string) (*Stock, error) {
	if r == nil || r.db == nil {
		return nil, ErrNilDB
	}
	var stock Stock
	err := r.db.WithContext(ctx).
		Preload("Warehouse").
		Preload("Product").
		First(&stock, "warehouse_id = ? AND product_id = ?", warehouseID, productID).Error
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *repository) FindAllStocks(ctx context.Context, req pagination.PageRequest) ([]Stock, int64, error) {
	if r == nil || r.db == nil {
		return nil, 0, ErrNilDB
	}
	req.Normalize()

	var stocks []Stock
	var total int64

	query := r.db.WithContext(ctx).Model(&Stock{}).Preload("Warehouse").Preload("Product")
	query.Count(&total)

	err := query.
		Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(req.Offset()).
		Find(&stocks).Error

	return stocks, total, err
}

func (r *repository) AdjustStock(ctx context.Context, warehouseID, productID uuid.UUID, delta int) (*Stock, error) {
	if r == nil || r.db == nil {
		return nil, ErrNilDB
	}

	var stock Stock
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Upsert stock entry using OnConflict
		stock = Stock{
			WarehouseID: warehouseID,
			ProductID:   productID,
			Quantity:    delta,
		}

		err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "warehouse_id"}, {Name: "product_id"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"quantity":   gorm.Expr("stocks.quantity + ?", delta),
				"updated_at": gorm.Expr("NOW()"),
			}),
		}).Create(&stock).Error

		if err != nil {
			return err
		}

		return tx.Preload("Warehouse").Preload("Product").First(&stock, "warehouse_id = ? AND product_id = ?", warehouseID, productID).Error
	})

	if err != nil {
		return nil, err
	}

	return &stock, nil
}
