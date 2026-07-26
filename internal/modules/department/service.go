package department

import (
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewService(
	repository Repository,
	logger *zap.Logger,
) Service {

	return &service{
		repository: repository,
		logger:     logger,
	}
}

type Service interface {
	Create(CreateDepartmentRequest) (*Department, error)
	FindAll() ([]Department, error)
	FindByID(string) (*Department, error)
	Update(string, UpdateDepartmentRequest) (*Department, error)
	Delete(string) error
}

type service struct {
	repository Repository
	logger     *zap.Logger
}

func (s *service) Create(req CreateDepartmentRequest) (*Department, error) {

	department := &Department{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	}

	if err := s.repository.Create(department); err != nil {
		return nil, err
	}

	return department, nil
}

func (s *service) FindAll() ([]Department, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(id string) (*Department, error) {
	return s.repository.FindByID(id)
}

func (s *service) Update(id string, req UpdateDepartmentRequest) (*Department, error) {

	department, err := s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	department.Name = req.Name
	department.Code = req.Code
	department.Description = req.Description

	if err := s.repository.Update(department); err != nil {
		return nil, err
	}

	return department, nil
}

func (s *service) Delete(id string) error {

	_, err := s.repository.FindByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	return s.repository.Delete(id)
}
