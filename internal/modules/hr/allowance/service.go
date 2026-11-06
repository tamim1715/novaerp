package allowance

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	CreateType(ctx context.Context, req CreateAllowanceTypeRequest) (*AllowanceType, error)
	FindAllTypes(ctx context.Context, req pagination.PageRequest) ([]AllowanceType, int64, error)
	FindTypeByID(ctx context.Context, id string) (*AllowanceType, error)
	UpdateType(ctx context.Context, id string, req UpdateAllowanceTypeRequest) (*AllowanceType, error)
	DeleteType(ctx context.Context, id string) error

	AssignAllowance(ctx context.Context, req AssignAllowanceRequest) (*EmployeeAllowance, error)
	FindEmployeeAllowances(ctx context.Context, req pagination.PageRequest) ([]EmployeeAllowance, int64, error)
	CalculateEmployeeAllowances(ctx context.Context, employeeID uuid.UUID, basicSalary float64) (float64, []CalculatedAllowanceItem, error)
	DeleteEmployeeAllowance(ctx context.Context, id string) error
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

func (s *service) CreateType(ctx context.Context, req CreateAllowanceTypeRequest) (*AllowanceType, error) {
	existing, err := s.repo.FindTypeByCode(ctx, req.Code)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("allowance type with code %s already exists", req.Code)
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	calcType := AllowanceCalculationType(req.CalculationType)
	if calcType != CalculationPercentage {
		calcType = CalculationFixed
	}

	at := &AllowanceType{
		Name:            req.Name,
		Code:            req.Code,
		CalculationType: calcType,
		DefaultAmount:   req.DefaultAmount,
		IsTaxable:       req.IsTaxable,
		IsActive:        true,
		Description:     req.Description,
	}

	if err := s.repo.CreateType(ctx, at); err != nil {
		return nil, fmt.Errorf("failed to create allowance type: %w", err)
	}

	return at, nil
}

func (s *service) FindAllTypes(ctx context.Context, req pagination.PageRequest) ([]AllowanceType, int64, error) {
	return s.repo.FindAllTypes(ctx, req)
}

func (s *service) FindTypeByID(ctx context.Context, id string) (*AllowanceType, error) {
	return s.repo.FindTypeByID(ctx, id)
}

func (s *service) UpdateType(ctx context.Context, id string, req UpdateAllowanceTypeRequest) (*AllowanceType, error) {
	at, err := s.repo.FindTypeByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		at.Name = req.Name
	}
	if req.CalculationType != "" {
		calcType := AllowanceCalculationType(req.CalculationType)
		if calcType == CalculationPercentage || calcType == CalculationFixed {
			at.CalculationType = calcType
		}
	}
	if req.DefaultAmount != nil {
		at.DefaultAmount = *req.DefaultAmount
	}
	if req.IsTaxable != nil {
		at.IsTaxable = *req.IsTaxable
	}
	if req.IsActive != nil {
		at.IsActive = *req.IsActive
	}
	if req.Description != "" {
		at.Description = req.Description
	}

	if err := s.repo.UpdateType(ctx, at); err != nil {
		return nil, err
	}

	return at, nil
}

func (s *service) DeleteType(ctx context.Context, id string) error {
	return s.repo.DeleteType(ctx, id)
}

func (s *service) AssignAllowance(ctx context.Context, req AssignAllowanceRequest) (*EmployeeAllowance, error) {
	empID, err := uuid.Parse(req.EmployeeID)
	if err != nil {
		return nil, errors.New("invalid employee ID format")
	}

	atID, err := uuid.Parse(req.AllowanceTypeID)
	if err != nil {
		return nil, errors.New("invalid allowance type ID format")
	}

	if _, err := s.employeeRepo.FindByID(ctx, req.EmployeeID); err != nil {
		return nil, fmt.Errorf("employee not found: %w", err)
	}

	at, err := s.repo.FindTypeByID(ctx, req.AllowanceTypeID)
	if err != nil {
		return nil, fmt.Errorf("allowance type not found: %w", err)
	}

	amount := req.Amount
	if amount == 0 {
		amount = at.DefaultAmount
	}

	ea := &EmployeeAllowance{
		EmployeeID:      empID,
		AllowanceTypeID: atID,
		Amount:          amount,
		IsActive:        true,
	}

	if err := s.repo.AssignEmployeeAllowance(ctx, ea); err != nil {
		return nil, fmt.Errorf("failed to assign allowance: %w", err)
	}

	return s.repo.FindEmployeeAllowanceByID(ctx, ea.ID.String())
}

func (s *service) FindEmployeeAllowances(ctx context.Context, req pagination.PageRequest) ([]EmployeeAllowance, int64, error) {
	return s.repo.FindEmployeeAllowances(ctx, req)
}

func (s *service) CalculateEmployeeAllowances(ctx context.Context, employeeID uuid.UUID, basicSalary float64) (float64, []CalculatedAllowanceItem, error) {
	allowances, err := s.repo.FindActiveAllowancesByEmployeeID(ctx, employeeID)
	if err != nil {
		return 0, nil, err
	}

	var total float64
	var items []CalculatedAllowanceItem

	for _, ea := range allowances {
		var calcVal float64
		amt := ea.Amount
		if amt == 0 && ea.AllowanceType.DefaultAmount > 0 {
			amt = ea.AllowanceType.DefaultAmount
		}

		if ea.AllowanceType.CalculationType == CalculationPercentage {
			calcVal = math.Round((basicSalary*(amt/100.0))*100) / 100
		} else {
			calcVal = math.Round(amt*100) / 100
		}

		total += calcVal
		items = append(items, CalculatedAllowanceItem{
			AllowanceTypeID: ea.AllowanceTypeID.String(),
			Code:            ea.AllowanceType.Code,
			Name:            ea.AllowanceType.Name,
			CalculationType: string(ea.AllowanceType.CalculationType),
			RateOrAmount:    amt,
			CalculatedTotal: calcVal,
		})
	}

	return math.Round(total*100) / 100, items, nil
}

func (s *service) DeleteEmployeeAllowance(ctx context.Context, id string) error {
	return s.repo.DeleteEmployeeAllowance(ctx, id)
}
