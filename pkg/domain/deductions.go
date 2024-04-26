package domain

import (
	"github.com/h3xry/assessment-tax/pkg/models"
)

type DeductionsRepository interface {
	Find(name string) (*models.Deductions, error)
	Update(model *models.Deductions) error
}

type DeductionsUsecase interface {
	Find(name string) (*models.Deductions, error)
	Update(model *models.Deductions) error
}
