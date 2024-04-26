package domain

type TaxUsecase interface {
	CalculateTax(income float64, wht float64, allowances []TaxAllowance) float64
}

type TaxAllowance struct {
	AllowanceType string  `json:"allowanceType" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
}

const (
	TaxDonationMaxDeduction = 100000.0
)

func (tax *TaxAllowance) Validate() {
	if tax.AllowanceType == "donation" && tax.Amount > TaxDonationMaxDeduction {
		tax.Amount = TaxDonationMaxDeduction
	}
}
