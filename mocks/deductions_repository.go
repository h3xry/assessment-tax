package mocks

import (
	"github.com/h3xry/assessment-tax/pkg/models"
	"github.com/stretchr/testify/mock"
)

type DeductionsRepo struct {
	mock.Mock
}

func (m *DeductionsRepo) Update(model *models.Deductions) error {
	args := m.Called(model)
	return args.Error(0)
}

func (m *DeductionsRepo) Find(name string) (*models.Deductions, error) {
	args := m.Called(name)
	return args.Get(0).(*models.Deductions), args.Error(1)
}
