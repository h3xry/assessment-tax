package deductions

import (
	"github.com/h3xry/assessment-tax/pkg/models"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) Update(model *models.Deductions) error {
	return r.DB.Where("name = ?", model.Name).Save(model).Error
}

func (r *repository) Find(name string) (*models.Deductions, error) {
	var deductions models.Deductions
	if err := r.DB.Where("name = ?", name).First(&deductions).Error; err != nil {
		return nil, err
	}
	return &deductions, nil
}
