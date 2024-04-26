package tax

import "github.com/h3xry/assessment-tax/pkg/domain"

type useCase struct {
}

func NewUseCase() domain.TaxUsecase {
	return &useCase{}
}

func (uc *useCase) CalculateTax(income float64, wht float64, allowances []domain.TaxAllowance) float64 {
	totalIncome := income
	for _, allowance := range allowances {
		allowance.Validate()
		totalIncome -= allowance.Amount
	}
	var tax float64
	if totalIncome > 2000000 {
		tax += (totalIncome - 2000000) * 0.35
		totalIncome = 2000000
	}
	if totalIncome > 1000000 {
		tax += (totalIncome - 1000000) * 0.2
		totalIncome = 1000000
	}
	if totalIncome > 500000 {
		tax += (totalIncome - 500000) * 0.15
		totalIncome = 500000
	}
	if totalIncome > 150000 {
		tax += (totalIncome - 150000) * 0.1
		totalIncome = 150000
	}
	return tax - wht
}
