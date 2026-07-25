package department

import "gorm.io/gorm"

type Repository interface {
	Create(*Department) error
	FindAll() ([]Department, error)
	FindByID(string) (*Department, error)
	Update(*Department) error
	Delete(string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(department *Department) error {
	return r.db.Create(department).Error
}

func (r *repository) FindAll() ([]Department, error) {

	var departments []Department

	err := r.db.Find(&departments).Error

	return departments, err
}

func (r *repository) FindByID(id string) (*Department, error) {

	var department Department

	err := r.db.First(&department, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (r *repository) Update(department *Department) error {
	return r.db.Save(department).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&Department{}, "id = ?", id).Error
}
