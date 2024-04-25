package deductions

import (
	"github.com/h3xry/assessment-tax/pkg/domain"
	"github.com/h3xry/assessment-tax/pkg/model"
)

type useCase struct {
	repo domain.DeductionsRepository
}

func NewUseCase(repo domain.DeductionsRepository) domain.DeductionsUsecase {
	return &useCase{
		repo: repo,
	}
}

func (u *useCase) Find(name string) (*model.Deductions, error) {
	return u.repo.Find(name)
}

func (u *useCase) Update(model *model.Deductions) error {
	return u.repo.Update(model)
}
