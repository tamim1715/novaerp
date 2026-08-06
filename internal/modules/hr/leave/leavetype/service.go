package leavetype

import (
	"context"
	"fmt"

	"github.com/tamim1715/novaerp/internal/common/pagination"
	"go.uber.org/zap"
)

type Service interface {
	CreateLeaveType(ctx context.Context, req CreateLeaveTypeRequest) (*LeaveType, error)
	FindAllLeaveTypes(ctx context.Context, req pagination.PageRequest) ([]LeaveType, int64, error)
	FindLeaveTypeByID(ctx context.Context, id string) (*LeaveType, error)
	UpdateLeaveType(ctx context.Context, id string, req UpdateLeaveTypeRequest) (*LeaveType, error)
	DeleteLeaveType(ctx context.Context, id string) error
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

func (s *service) CreateLeaveType(ctx context.Context, req CreateLeaveTypeRequest) (*LeaveType, error) {
	isPaid := true
	if req.IsPaid != nil {
		isPaid = *req.IsPaid
	}

	lt := &LeaveType{
		Name:           req.Name,
		Code:           req.Code,
		MaxDaysPerYear: req.MaxDaysPerYear,
		IsPaid:         isPaid,
		Description:    req.Description,
	}

	if err := s.repo.Create(ctx, lt); err != nil {
		return nil, fmt.Errorf("failed to create leave type: %w", err)
	}

	return lt, nil
}

func (s *service) FindAllLeaveTypes(ctx context.Context, req pagination.PageRequest) ([]LeaveType, int64, error) {
	return s.repo.FindAll(ctx, req)
}

func (s *service) FindLeaveTypeByID(ctx context.Context, id string) (*LeaveType, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) UpdateLeaveType(ctx context.Context, id string, req UpdateLeaveTypeRequest) (*LeaveType, error) {
	lt, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		lt.Name = req.Name
	}
	if req.Code != "" {
		lt.Code = req.Code
	}
	if req.MaxDaysPerYear != nil {
		lt.MaxDaysPerYear = *req.MaxDaysPerYear
	}
	if req.IsPaid != nil {
		lt.IsPaid = *req.IsPaid
	}
	if req.Description != "" {
		lt.Description = req.Description
	}

	if err := s.repo.Update(ctx, lt); err != nil {
		return nil, err
	}

	return lt, nil
}

func (s *service) DeleteLeaveType(ctx context.Context, id string) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
