package deductions

import (
	"github.com/h3xry/assessment-tax/pkg/model"
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

func (r *repository) Update(model *model.Deductions) error {
	return r.DB.Save(model).Error
}

func (r *repository) Find(name string) (*model.Deductions, error) {
	var deductions model.Deductions
	if err := r.DB.Where("name = ?", name).First(&deductions).Error; err != nil {
		return nil, err
	}
	return &deductions, nil
}
