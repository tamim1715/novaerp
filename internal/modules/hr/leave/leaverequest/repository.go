package leaverequest

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, lr *LeaveRequest) error
	FindAll(ctx context.Context, req pagination.PageRequest) ([]LeaveRequest, int64, error)
	FindByID(ctx context.Context, id string) (*LeaveRequest, error)
	FindApprovedByEmployeeAndDate(ctx context.Context, employeeID uuid.UUID, date time.Time) (*LeaveRequest, error)
	FindApprovedByEmployeeAndDateRange(ctx context.Context, employeeID uuid.UUID, startDate, endDate time.Time) ([]LeaveRequest, error)
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

func (r *repository) FindApprovedByEmployeeAndDate(ctx context.Context, employeeID uuid.UUID, date time.Time) (*LeaveRequest, error) {
	var lr LeaveRequest
	targetDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	err := r.db.WithContext(ctx).
		Preload("LeaveType").
		Where("employee_id = ? AND status = ? AND start_date <= ? AND end_date >= ?", employeeID, StatusApproved, targetDate, targetDate).
		First(&lr).Error
	if err != nil {
		return nil, err
	}
	return &lr, nil
}

func (r *repository) FindApprovedByEmployeeAndDateRange(ctx context.Context, employeeID uuid.UUID, startDate, endDate time.Time) ([]LeaveRequest, error) {
	var list []LeaveRequest
	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	end := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, time.UTC)
	err := r.db.WithContext(ctx).
		Preload("LeaveType").
		Where("employee_id = ? AND status = ? AND ((start_date BETWEEN ? AND ?) OR (end_date BETWEEN ? AND ?) OR (start_date <= ? AND end_date >= ?))",
			employeeID, StatusApproved, start, end, start, end, start, end).
		Find(&list).Error
	return list, err
}

func (r *repository) Update(ctx context.Context, lr *LeaveRequest) error {
	return r.db.WithContext(ctx).Save(lr).Error
}
