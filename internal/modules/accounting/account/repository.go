package account

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, account *Account) error
	FindAll(ctx context.Context, req pagination.PageRequest, accountType string) ([]Account, int64, error)
	FindByID(ctx context.Context, id string) (*Account, error)
	FindByCode(ctx context.Context, code string) (*Account, error)
	GetTree(ctx context.Context) ([]Account, error)
	Update(ctx context.Context, account *Account) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
	SeedDefaults(ctx context.Context) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, account *Account) error {
	return r.db.WithContext(ctx).Create(account).Error
}

func (r *repository) FindAll(ctx context.Context, req pagination.PageRequest, accountType string) ([]Account, int64, error) {
	var accounts []Account
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&Account{})

	if accountType != "" {
		db = db.Where("type = ?", accountType)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := req.Offset()
	err := db.Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(offset).
		Find(&accounts).Error

	return accounts, total, err
}

func (r *repository) FindByID(ctx context.Context, id string) (*Account, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var account Account
	err = r.db.WithContext(ctx).Preload("Parent").Preload("Children").First(&account, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *repository) FindByCode(ctx context.Context, code string) (*Account, error) {
	var account Account
	err := r.db.WithContext(ctx).First(&account, "code = ?", code).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *repository) GetTree(ctx context.Context) ([]Account, error) {
	var rootAccounts []Account
	err := r.db.WithContext(ctx).
		Where("parent_id IS NULL").
		Preload("Children").
		Preload("Children.Children").
		Order("code ASC").
		Find(&rootAccounts).Error
	return rootAccounts, err
}

func (r *repository) Update(ctx context.Context, account *Account) error {
	return r.db.WithContext(ctx).Save(account).Error
}

func (r *repository) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}

	// Check if system account
	var acc Account
	if err := r.db.WithContext(ctx).First(&acc, "id = ?", uid).Error; err != nil {
		return err
	}
	if acc.IsSystem {
		return errors.New("cannot delete protected system account")
	}

	// Check if it has child accounts
	var childCount int64
	r.db.WithContext(ctx).Model(&Account{}).Where("parent_id = ?", uid).Count(&childCount)
	if childCount > 0 {
		return errors.New("cannot delete account with existing child sub-accounts")
	}

	return r.db.WithContext(ctx).Delete(&Account{}, "id = ?", uid).Error
}

func (r *repository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&Account{}).Count(&count).Error
	return count, err
}

func (r *repository) SeedDefaults(ctx context.Context) error {
	return SeedStandardChartOfAccounts(r.db.WithContext(ctx))
}
