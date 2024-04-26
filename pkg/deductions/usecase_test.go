package deductions

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/h3xry/assessment-tax/mocks"
	"github.com/h3xry/assessment-tax/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestUsecaseFind(t *testing.T) {
	mockDeduction := models.Deductions{}
	gofakeit.Struct(&mockDeduction)

	mockDeductionRepo := new(mocks.DeductionsRepo)
	mockUserUsecase := new(mocks.DeductionsUsecase)

	t.Run("success", func(t *testing.T) {
		mockDeductionRepo.On("Find", mockDeduction.Name).Return(&mockDeduction, nil).Once()
		usecase := NewUseCase(mockDeductionRepo)
		result, err := usecase.Find(mockDeduction.Name)
		assert.NoError(t, err)
		assert.Equal(t, &mockDeduction, result)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDeductionRepo.On("Find", mockDeduction.Name).Return(&models.Deductions{}, assert.AnError).Once()
		usecase := NewUseCase(mockDeductionRepo)
		result, err := usecase.Find(mockDeduction.Name)
		assert.Error(t, err)
		assert.Equal(t, &models.Deductions{}, result)
		mockUserUsecase.AssertExpectations(t)
	})
}

func TestUsecaseUpdate(t *testing.T) {
	mockDeduction := models.Deductions{}
	gofakeit.Struct(&mockDeduction)

	mockDeductionRepo := new(mocks.DeductionsRepo)
	mockUserUsecase := new(mocks.DeductionsUsecase)

	t.Run("success", func(t *testing.T) {
		mockDeductionRepo.On("Update", &mockDeduction).Return(nil).Once()
		usecase := NewUseCase(mockDeductionRepo)
		err := usecase.Update(&mockDeduction)
		assert.NoError(t, err)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDeductionRepo.On("Update", &mockDeduction).Return(assert.AnError).Once()
		usecase := NewUseCase(mockDeductionRepo)
		err := usecase.Update(&mockDeduction)
		assert.Error(t, err)
		mockUserUsecase.AssertExpectations(t)
	})

	t.Run("error amount exceed", func(t *testing.T) {
		mockDeduction.Name = "kReceipt"
		mockDeduction.Amount = 100001
		usecase := NewUseCase(mockDeductionRepo)
		err := usecase.Update(&mockDeduction)
		assert.Error(t, err)
		mockUserUsecase.AssertExpectations(t)
	})
}
