package allowance

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	CreateType(ctx context.Context, at *AllowanceType) error
	FindAllTypes(ctx context.Context, req pagination.PageRequest) ([]AllowanceType, int64, error)
	FindTypeByID(ctx context.Context, id string) (*AllowanceType, error)
	FindTypeByCode(ctx context.Context, code string) (*AllowanceType, error)
	UpdateType(ctx context.Context, at *AllowanceType) error
	DeleteType(ctx context.Context, id string) error

	AssignEmployeeAllowance(ctx context.Context, ea *EmployeeAllowance) error
	FindEmployeeAllowanceByID(ctx context.Context, id string) (*EmployeeAllowance, error)
	FindActiveAllowancesByEmployeeID(ctx context.Context, employeeID uuid.UUID) ([]EmployeeAllowance, error)
	FindEmployeeAllowances(ctx context.Context, req pagination.PageRequest) ([]EmployeeAllowance, int64, error)
	DeleteEmployeeAllowance(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateType(ctx context.Context, at *AllowanceType) error {
	return r.db.WithContext(ctx).Create(at).Error
}

func (r *repository) FindAllTypes(ctx context.Context, req pagination.PageRequest) ([]AllowanceType, int64, error) {
	var list []AllowanceType
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&AllowanceType{})

	if req.Search != "" {
		search := "%" + req.Search + "%"
		db = db.Where("name ILIKE ? OR code ILIKE ?", search, search)
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

func (r *repository) FindTypeByID(ctx context.Context, id string) (*AllowanceType, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var at AllowanceType
	err = r.db.WithContext(ctx).First(&at, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &at, nil
}

func (r *repository) FindTypeByCode(ctx context.Context, code string) (*AllowanceType, error) {
	var at AllowanceType
	err := r.db.WithContext(ctx).First(&at, "code = ?", code).Error
	if err != nil {
		return nil, err
	}
	return &at, nil
}

func (r *repository) UpdateType(ctx context.Context, at *AllowanceType) error {
	return r.db.WithContext(ctx).Save(at).Error
}

func (r *repository) DeleteType(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return r.db.WithContext(ctx).Delete(&AllowanceType{}, "id = ?", uid).Error
}

func (r *repository) AssignEmployeeAllowance(ctx context.Context, ea *EmployeeAllowance) error {
	return r.db.WithContext(ctx).Create(ea).Error
}

func (r *repository) FindEmployeeAllowanceByID(ctx context.Context, id string) (*EmployeeAllowance, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var ea EmployeeAllowance
	err = r.db.WithContext(ctx).Preload("Employee").Preload("AllowanceType").First(&ea, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &ea, nil
}

func (r *repository) FindActiveAllowancesByEmployeeID(ctx context.Context, employeeID uuid.UUID) ([]EmployeeAllowance, error) {
	var list []EmployeeAllowance
	err := r.db.WithContext(ctx).
		Preload("AllowanceType").
		Where("employee_id = ? AND is_active = true", employeeID).
		Find(&list).Error
	return list, err
}

func (r *repository) FindEmployeeAllowances(ctx context.Context, req pagination.PageRequest) ([]EmployeeAllowance, int64, error) {
	var list []EmployeeAllowance
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&EmployeeAllowance{}).Preload("Employee").Preload("AllowanceType")

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := db.Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(req.Offset()).
		Find(&list).Error

	return list, total, err
}

func (r *repository) DeleteEmployeeAllowance(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return r.db.WithContext(ctx).Delete(&EmployeeAllowance{}, "id = ?", uid).Error
}
