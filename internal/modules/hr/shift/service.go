package shift

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	CreateShift(ctx context.Context, req CreateShiftRequest) (*Shift, error)
	FindAllShifts(ctx context.Context, req pagination.PageRequest) ([]Shift, int64, error)
	FindShiftByID(ctx context.Context, id string) (*Shift, error)
	UpdateShift(ctx context.Context, id string, req UpdateShiftRequest) (*Shift, error)
	DeleteShift(ctx context.Context, id string) error

	AssignShift(ctx context.Context, req AssignShiftRequest) (*EmployeeShift, error)
	FindActiveEmployeeShift(ctx context.Context, employeeID uuid.UUID, date time.Time) (*EmployeeShift, error)
	FindEmployeeShifts(ctx context.Context, req pagination.PageRequest) ([]EmployeeShift, int64, error)
}

type service struct {
	repo         Repository
	employeeRepo employee.Repository
	logger       *zap.Logger
}

func NewService(repo Repository, employeeRepo employee.Repository, logger *zap.Logger) Service {
	return &service{
		repo:         repo,
		employeeRepo: employeeRepo,
		logger:       logger,
	}
}

func (s *service) CreateShift(ctx context.Context, req CreateShiftRequest) (*Shift, error) {
	existing, err := s.repo.FindShiftByCode(ctx, req.Code)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("shift with code %s already exists", req.Code)
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	workHours := req.WorkHours
	if workHours <= 0 {
		workHours = 8.0
	}
	halfDayHours := req.HalfDayHours
	if halfDayHours <= 0 {
		halfDayHours = 4.0
	}
	graceMinutes := req.GracePeriodMinutes
	if graceMinutes < 0 {
		graceMinutes = 15
	}

	sh := &Shift{
		Name:               req.Name,
		Code:               req.Code,
		StartTime:          req.StartTime,
		EndTime:            req.EndTime,
		GracePeriodMinutes: graceMinutes,
		WorkHours:          workHours,
		HalfDayHours:       halfDayHours,
		IsNightShift:       req.IsNightShift,
		IsActive:           true,
		Description:        req.Description,
	}

	if err := s.repo.CreateShift(ctx, sh); err != nil {
		return nil, fmt.Errorf("failed to create shift: %w", err)
	}

	return sh, nil
}

func (s *service) FindAllShifts(ctx context.Context, req pagination.PageRequest) ([]Shift, int64, error) {
	return s.repo.FindAllShifts(ctx, req)
}

func (s *service) FindShiftByID(ctx context.Context, id string) (*Shift, error) {
	return s.repo.FindShiftByID(ctx, id)
}

func (s *service) UpdateShift(ctx context.Context, id string, req UpdateShiftRequest) (*Shift, error) {
	sh, err := s.repo.FindShiftByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		sh.Name = req.Name
	}
	if req.StartTime != "" {
		sh.StartTime = req.StartTime
	}
	if req.EndTime != "" {
		sh.EndTime = req.EndTime
	}
	if req.GracePeriodMinutes > 0 {
		sh.GracePeriodMinutes = req.GracePeriodMinutes
	}
	if req.WorkHours > 0 {
		sh.WorkHours = req.WorkHours
	}
	if req.HalfDayHours > 0 {
		sh.HalfDayHours = req.HalfDayHours
	}
	sh.IsNightShift = req.IsNightShift
	if req.IsActive != nil {
		sh.IsActive = *req.IsActive
	}
	if req.Description != "" {
		sh.Description = req.Description
	}

	if err := s.repo.UpdateShift(ctx, sh); err != nil {
		return nil, err
	}

	return sh, nil
}

func (s *service) DeleteShift(ctx context.Context, id string) error {
	return s.repo.DeleteShift(ctx, id)
}

func (s *service) AssignShift(ctx context.Context, req AssignShiftRequest) (*EmployeeShift, error) {
	empID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		return nil, errors.New("invalid employee ID format")
	}

	shID, err := uuid.Parse(req.ShiftID)
	if err != nil {
		return nil, errors.New("invalid shift ID format")
	}

	if _, err := s.employeeRepo.FindByID(ctx, req.EmployeeID); err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	if _, err := s.repo.FindShiftByID(ctx, req.ShiftID); err != nil {
		return nil, fmt.Errorf("shift not found: %w", err)
	}

	effFrom, err := time.Parse("2006-01-02", req.EffectiveFrom)
	if err != nil {
		return nil, errors.New("invalid effectiveFrom format (use YYYY-MM-DD)")
	}

	var effTo *time.Time
	if req.EffectiveTo != "" {
		t, err := time.Parse("2006-01-02", req.EffectiveTo)
		if err != nil {
			return nil, errors.New("invalid effectiveTo format (use YYYY-MM-DD)")
		}
		effTo = &t
	}

	es := &EmployeeShift{
		EmployeeID:    empID,
		ShiftID:       shID,
		EffectiveFrom: effFrom,
		EffectiveTo:   effTo,
		IsActive:      true,
	}

	if err := s.repo.AssignEmployeeShift(ctx, es); err != nil {
		return nil, fmt.Errorf("failed to assign shift: %w", err)
	}

	return s.repo.FindEmployeeShiftByID(ctx, es.ID.String())
}

func (s *service) FindActiveEmployeeShift(ctx context.Context, employeeID uuid.UUID, date time.Time) (*EmployeeShift, error) {
	return s.repo.FindActiveEmployeeShift(ctx, employeeID, date)
}

func (s *service) FindEmployeeShifts(ctx context.Context, req pagination.PageRequest) ([]EmployeeShift, int64, error) {
	return s.repo.FindEmployeeShifts(ctx, req)
}
