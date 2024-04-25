package domain

import "github.com/h3xry/assessment-tax/pkg/model"

type DeductionsRepository interface {
	Find(name string) (*model.Deductions, error)
	Update(model *model.Deductions) error
}

type DeductionsUsecase interface {
	Find(name string) (*model.Deductions, error)
	Update(model *model.Deductions) error
}
