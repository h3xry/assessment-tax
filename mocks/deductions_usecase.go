package mocks

import (
	"github.com/h3xry/assessment-tax/pkg/models"
	"github.com/stretchr/testify/mock"
)

type DeductionsUsecase struct {
	mock.Mock
}

func (m *DeductionsUsecase) Find(name string) error {
	ret := m.Called(name)
	return ret.Error(0)
}

func (m *DeductionsUsecase) Update(model *models.Deductions) error {
	ret := m.Called(model)
	return ret.Error(0)
}
