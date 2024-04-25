package deductions

import (
	"github.com/h3xry/assessment-tax/pkg/domain"
	"github.com/h3xry/assessment-tax/pkg/models"
)

type useCase struct {
	repo domain.DeductionsRepository
}

func NewUseCase(repo domain.DeductionsRepository) domain.DeductionsUsecase {
	return &useCase{
		repo: repo,
	}
}

func (u *useCase) Find(name string) (*models.Deductions, error) {
	return u.repo.Find(name)
}

func (u *useCase) Update(model *models.Deductions) error {
	if model.Name == "K-Recepit" && (model.Amount > 100000 || model.Amount < 1) {
		return domain.ErrAmountExceed
	}
	return u.repo.Update(model)
}
