package overtime

import (
	"context"
	"errors"
	"fmt"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	CreatePolicy(ctx context.Context, req CreateOvertimePolicyRequest) (*OvertimePolicy, error)
	FindAllPolicies(ctx context.Context, req pagination.PageRequest) ([]OvertimePolicy, int64, error)
	FindPolicyByID(ctx context.Context, id string) (*OvertimePolicy, error)
	GetDefaultPolicy(ctx context.Context) (*OvertimePolicy, error)
	UpdatePolicy(ctx context.Context, id string, req UpdateOvertimePolicyRequest) (*OvertimePolicy, error)
	DeletePolicy(ctx context.Context, id string) error
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

func (s *service) CreatePolicy(ctx context.Context, req CreateOvertimePolicyRequest) (*OvertimePolicy, error) {
	existing, err := s.repo.FindPolicyByCode(ctx, req.Code)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("overtime policy with code %s already exists", req.Code)
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	minMinutes := req.MinOvertimeMinutes
	if minMinutes < 0 {
		minMinutes = 30
	}
	maxHours := req.MaxOvertimeHoursPerDay
	if maxHours <= 0 {
		maxHours = 4.0
	}
	weekdayMul := req.WeekdayRateMultiplier
	if weekdayMul <= 0 {
		weekdayMul = 1.5
	}
	weekendMul := req.WeekendRateMultiplier
	if weekendMul <= 0 {
		weekendMul = 2.0
	}
	holidayMul := req.HolidayRateMultiplier
	if holidayMul <= 0 {
		holidayMul = 2.0
	}

	policy := &OvertimePolicy{
		Name:                   req.Name,
		Code:                   req.Code,
		MinOvertimeMinutes:     minMinutes,
		MaxOvertimeHoursPerDay: maxHours,
		WeekdayRateMultiplier:  weekdayMul,
		WeekendRateMultiplier:  weekendMul,
		HolidayRateMultiplier:  holidayMul,
		IsDefault:              req.IsDefault,
		IsActive:               true,
		Description:            req.Description,
	}

	if err := s.repo.CreatePolicy(ctx, policy); err != nil {
		return nil, fmt.Errorf("failed to create overtime policy: %w", err)
	}

	return policy, nil
}

func (s *service) FindAllPolicies(ctx context.Context, req pagination.PageRequest) ([]OvertimePolicy, int64, error) {
	return s.repo.FindAllPolicies(ctx, req)
}

func (s *service) FindPolicyByID(ctx context.Context, id string) (*OvertimePolicy, error) {
	return s.repo.FindPolicyByID(ctx, id)
}

func (s *service) GetDefaultPolicy(ctx context.Context) (*OvertimePolicy, error) {
	policy, err := s.repo.FindDefaultPolicy(ctx)
	if err != nil {
		// If none configured, return fallback standard policy in-memory
		return &OvertimePolicy{
			Name:                   "Standard Overtime Policy",
			Code:                   "OT_DEFAULT",
			MinOvertimeMinutes:     30,
			MaxOvertimeHoursPerDay: 4.0,
			WeekdayRateMultiplier:  1.5,
			WeekendRateMultiplier:  2.0,
			HolidayRateMultiplier:  2.0,
			IsDefault:              true,
			IsActive:               true,
		}, nil
	}
	return policy, nil
}

func (s *service) UpdatePolicy(ctx context.Context, id string, req UpdateOvertimePolicyRequest) (*OvertimePolicy, error) {
	policy, err := s.repo.FindPolicyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		policy.Name = req.Name
	}
	if req.MinOvertimeMinutes != nil {
		policy.MinOvertimeMinutes = *req.MinOvertimeMinutes
	}
	if req.MaxOvertimeHoursPerDay != nil {
		policy.MaxOvertimeHoursPerDay = *req.MaxOvertimeHoursPerDay
	}
	if req.WeekdayRateMultiplier != nil {
		policy.WeekdayRateMultiplier = *req.WeekdayRateMultiplier
	}
	if req.WeekendRateMultiplier != nil {
		policy.WeekendRateMultiplier = *req.WeekendRateMultiplier
	}
	if req.HolidayRateMultiplier != nil {
		policy.HolidayRateMultiplier = *req.HolidayRateMultiplier
	}
	if req.IsDefault != nil {
		policy.IsDefault = *req.IsDefault
	}
	if req.IsActive != nil {
		policy.IsActive = *req.IsActive
	}
	if req.Description != "" {
		policy.Description = req.Description
	}

	if err := s.repo.UpdatePolicy(ctx, policy); err != nil {
		return nil, err
	}

	return policy, nil
}

func (s *service) DeletePolicy(ctx context.Context, id string) error {
	return s.repo.DeletePolicy(ctx, id)
}
