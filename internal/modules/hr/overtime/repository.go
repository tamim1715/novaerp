package overtime

import (
	"context"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
)

type Repository interface {
	CreatePolicy(ctx context.Context, p *OvertimePolicy) error
	FindAllPolicies(ctx context.Context, req pagination.PageRequest) ([]OvertimePolicy, int64, error)
	FindPolicyByID(ctx context.Context, id string) (*OvertimePolicy, error)
	FindPolicyByCode(ctx context.Context, code string) (*OvertimePolicy, error)
	FindDefaultPolicy(ctx context.Context) (*OvertimePolicy, error)
	UpdatePolicy(ctx context.Context, p *OvertimePolicy) error
	DeletePolicy(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreatePolicy(ctx context.Context, p *OvertimePolicy) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *repository) FindAllPolicies(ctx context.Context, req pagination.PageRequest) ([]OvertimePolicy, int64, error) {
	var list []OvertimePolicy
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&OvertimePolicy{})

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

func (r *repository) FindPolicyByID(ctx context.Context, id string) (*OvertimePolicy, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var p OvertimePolicy
	err = r.db.WithContext(ctx).First(&p, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) FindPolicyByCode(ctx context.Context, code string) (*OvertimePolicy, error) {
	var p OvertimePolicy
	err := r.db.WithContext(ctx).First(&p, "code = ?", code).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) FindDefaultPolicy(ctx context.Context) (*OvertimePolicy, error) {
	var p OvertimePolicy
	err := r.db.WithContext(ctx).Where("is_default = true AND is_active = true").First(&p).Error
	if err != nil {
		// Fallback to any active policy
		err = r.db.WithContext(ctx).Where("is_active = true").First(&p).Error
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *repository) UpdatePolicy(ctx context.Context, p *OvertimePolicy) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *repository) DeletePolicy(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	return r.db.WithContext(ctx).Delete(&OvertimePolicy{}, "id = ?", uid).Error
}
