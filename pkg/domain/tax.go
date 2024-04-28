package domain

type TaxUsecase interface {
	CalculateTax(income float64, wht float64, allowances []TaxAllowance) (float64, float64, []TaxLevel)
}

type TaxAllowance struct {
	AllowanceType string  `json:"allowanceType" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
	MaxDeduction  float64 `json:"-"`
}

type TaxLevel struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

const (
	TaxDonationMaxDeduction = 100000.0
)

func (tax *TaxAllowance) Validate() {
	if tax.AllowanceType == "donation" && tax.Amount > TaxDonationMaxDeduction {
		tax.Amount = TaxDonationMaxDeduction
		return
	}
	if tax.AllowanceType == "personalDeduction" && tax.Amount > tax.MaxDeduction {
		tax.Amount = tax.MaxDeduction
		return
	}
	if tax.AllowanceType == "k-receipt" && tax.Amount > tax.MaxDeduction {
		tax.Amount = tax.MaxDeduction
		return
	}
}
