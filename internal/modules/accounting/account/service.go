package account

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) (*Account, error)
	FindAllAccounts(ctx context.Context, req pagination.PageRequest, accountType string) ([]Account, int64, error)
	FindAccountByID(ctx context.Context, id string) (*Account, error)
	GetAccountTree(ctx context.Context) ([]Account, error)
	UpdateAccount(ctx context.Context, id string, req UpdateAccountRequest) (*Account, error)
	DeleteAccount(ctx context.Context, id string) error
	SeedChartOfAccounts(ctx context.Context) error
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

func (s *service) CreateAccount(ctx context.Context, req CreateAccountRequest) (*Account, error) {
	// Determine normal balance based on account type if not explicitly provided
	normalBal := req.NormalBalance
	if normalBal == "" {
		if req.Type == TypeAsset || req.Type == TypeExpense {
			normalBal = BalanceDebit
		} else {
			normalBal = BalanceCredit
		}
	}

	curr := "USD"
	if req.Currency != "" {
		curr = req.Currency
	}

	isActive := true
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	// Validate parent if provided
	if req.ParentID != nil {
		if _, err := s.repo.FindByID(ctx, req.ParentID.String()); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("parent account not found")
			}
			return nil, err
		}
	}

	account := &Account{
		Code:          req.Code,
		Name:          req.Name,
		Type:          req.Type,
		NormalBalance: normalBal,
		ParentID:      req.ParentID,
		Currency:      curr,
		IsActive:      isActive,
		Description:   req.Description,
	}

	if err := s.repo.Create(ctx, account); err != nil {
		return nil, err
	}

	return account, nil
}

func (s *service) FindAllAccounts(ctx context.Context, req pagination.PageRequest, accountType string) ([]Account, int64, error) {
	return s.repo.FindAll(ctx, req, accountType)
}

func (s *service) FindAccountByID(ctx context.Context, id string) (*Account, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) GetAccountTree(ctx context.Context) ([]Account, error) {
	return s.repo.GetTree(ctx)
}

func (s *service) UpdateAccount(ctx context.Context, id string, req UpdateAccountRequest) (*Account, error) {
	account, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		account.Name = *req.Name
	}
	if req.Currency != nil {
		account.Currency = *req.Currency
	}
	if req.IsActive != nil {
		account.IsActive = *req.IsActive
	}
	if req.Description != nil {
		account.Description = *req.Description
	}
	if req.ParentID != nil {
		if *req.ParentID == account.ID {
			return nil, errors.New("account cannot be its own parent")
		}
		if *req.ParentID != uuid.Nil {
			if _, err := s.repo.FindByID(ctx, req.ParentID.String()); err != nil {
				return nil, errors.New("parent account not found")
			}
			account.ParentID = req.ParentID
		} else {
			account.ParentID = nil
		}
	}

	if err := s.repo.Update(ctx, account); err != nil {
		return nil, err
	}

	return account, nil
}

func (s *service) DeleteAccount(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) SeedChartOfAccounts(ctx context.Context) error {
	return s.repo.SeedDefaults(ctx)
}
