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
		totalIncome -= allowance.Amount
	}
	var tax float64 = totalIncome
	if tax > 2000000 {
		tax = (tax - 2000000) * 0.35
	}
	if tax > 1000000 {
		tax = (tax - 1000000) * 0.2
	}
	if tax > 500000 {
		tax = (tax - 500000) * 0.15
	}
	if tax > 150000 {
		tax = (tax - 150000) * 0.1
	}
	return tax - wht
}
