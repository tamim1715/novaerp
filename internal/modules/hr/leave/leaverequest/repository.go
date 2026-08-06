package leaverequest

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, lr *LeaveRequest) error
	FindAll(ctx context.Context, req pagination.PageRequest) ([]LeaveRequest, int64, error)
	FindByID(ctx context.Context, id string) (*LeaveRequest, error)
	Update(ctx context.Context, lr *LeaveRequest) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, lr *LeaveRequest) error {
	return r.db.WithContext(ctx).Create(lr).Error
}

func (r *repository) FindAll(ctx context.Context, req pagination.PageRequest) ([]LeaveRequest, int64, error) {
	var list []LeaveRequest
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&LeaveRequest{}).
		Preload("Employee").
		Preload("LeaveType")

	if req.Search != "" {
		search := "%" + req.Search + "%"
		db = db.Where("status ILIKE ? OR reason ILIKE ?", search, search)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(req.Offset()).
		Find(&list).Error

	return list, total, err
}

func (r *repository) FindByID(ctx context.Context, id string) (*LeaveRequest, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var lr LeaveRequest
	err = r.db.WithContext(ctx).
		Preload("Employee").
		Preload("LeaveType").
		First(&lr, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (r *repository) Update(ctx context.Context, lr *LeaveRequest) error {
	return r.db.WithContext(ctx).Save(lr).Error
}
