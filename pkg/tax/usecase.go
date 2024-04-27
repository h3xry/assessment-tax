package tax

import (
	"github.com/h3xry/assessment-tax/pkg/domain"
	"github.com/samber/lo"
)

type useCase struct {
}

func NewUseCase() domain.TaxUsecase {
	return &useCase{}
}

func (uc *useCase) CalculateTax(income float64, wht float64, allowances []domain.TaxAllowance) (float64, []domain.TaxLevel) {
	totalIncome := income
	for _, allowance := range allowances {
		allowance.Validate()
		totalIncome -= allowance.Amount
	}
	taxLevel := []domain.TaxLevel{}
	var tax, tempTax float64
	if totalIncome > 2000000 {
		tempTax = (2000000 - 1000000) * 0.2
		tax += tempTax
		totalIncome = 2000000
	}
	taxLevel = append(taxLevel, domain.TaxLevel{
		Level: "2,000,001 ขึ้นไป",
		Tax:   tempTax,
	})
	if totalIncome > 1000000 {
		tempTax = (totalIncome - 1000000) * 0.2
		tax += tempTax
		totalIncome = 1000000
	}
	taxLevel = append(taxLevel, domain.TaxLevel{
		Level: "1,000,001-2,000,000",
		Tax:   tempTax,
	})
	if totalIncome > 500000 {
		tempTax = (totalIncome - 500000) * 0.15
		tax += tempTax
		totalIncome = 500000
	}
	taxLevel = append(taxLevel, domain.TaxLevel{
		Level: "500,001-1,000,000",
		Tax:   tempTax,
	})
	if totalIncome > 150000 {
		tempTax = (totalIncome - 150000) * 0.1
		tax += tempTax
		totalIncome = 150000
	}
	taxLevel = append(taxLevel, domain.TaxLevel{
		Level: "150,001-500,000",
		Tax:   tempTax,
	})
	taxLevel = append(taxLevel, domain.TaxLevel{
		Level: "0-150,000",
		Tax:   0,
	})
	return tax - wht, lo.Reverse(taxLevel)
}
