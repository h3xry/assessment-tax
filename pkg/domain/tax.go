package domain

type Tax interface {
	CalculateTax(income float64, wht float64, allowances []TaxAllowance) float64
}

type TaxAllowance struct {
	AllowanceType string  `json:"allowanceType" validate:"required"`
	Amount        float64 `json:"amount" validate:"required"`
}
