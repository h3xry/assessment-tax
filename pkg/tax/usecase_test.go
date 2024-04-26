package tax

import (
	"testing"

	"github.com/h3xry/assessment-tax/pkg/domain"
)

func TestUsecaseCalculateTax(t *testing.T) {
	// totalIncome = 500000 - 60000 = 440000
	// tax = (totalIncome - 150000) * 0.1 = 29000
	// expected = 29000
	income := 500000.0
	wht := 0.0
	allowances := []domain.TaxAllowance{
		{
			AllowanceType: "donation",
			Amount:        0.0,
		},
		{
			AllowanceType: "personalDeduction",
			Amount:        60000.0,
		},
	}
	expected := 29000.0
	uc := NewUseCase()
	result := uc.CalculateTax(income, wht, allowances)
	if result != expected {
		t.Errorf("expected %f but got %f", expected, result)
	}
}
