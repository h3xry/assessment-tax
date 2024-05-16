package tax

import (
	"encoding/csv"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/h3xry/assessment-tax/pkg/domain"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type Handler struct {
	userCase         domain.TaxUsecase
	deductionUsecase domain.DeductionsUsecase
}

func NewHandler(route *echo.Group, userCase domain.TaxUsecase, deductionUsecase domain.DeductionsUsecase) *Handler {
	handler := Handler{
		userCase:         userCase,
		deductionUsecase: deductionUsecase,
	}
	route.POST("/calculations", handler.handleCalculation())
	route.POST("/calculations/upload-csv", handler.handleCalculationCSV())
	return &handler
}

type requestCalculation struct {
	TotalIncome float64               `json:"totalIncome" validate:"required"`
	Wht         float64               `json:"wht" validate:"min=0"`
	Allowances  []domain.TaxAllowance `json:"allowances,omitempty"`
}

func (h *Handler) handleCalculation() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(requestCalculation)
		if err := c.Bind(req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}

		if err := h.addPersonalDeduction(req); err != nil {
			return err
		}

		if err := h.validateAndAdjustAllowances(req); err != nil {
			return err
		}

		tax, taxRefund, taxLevel := h.userCase.CalculateTax(req.TotalIncome, req.Wht, req.Allowances)
		response := buildResponse(tax, taxRefund, taxLevel)

		return c.JSON(http.StatusOK, response)
	}
}

func (h *Handler) addPersonalDeduction(req *requestCalculation) error {
	personal, err := h.deductionUsecase.Find("personalDeduction")
	if err != nil {
		return err
	}
	req.Allowances = append(req.Allowances, domain.TaxAllowance{
		AllowanceType: "personalDeduction",
		Amount:        personal.Amount,
		MaxDeduction:  personal.Amount,
	})
	return nil
}

func (h *Handler) validateAndAdjustAllowances(req *requestCalculation) error {
	for i, v := range req.Allowances {
		if err := validateAllowance(v); err != nil {
			return err
		}
		if v.AllowanceType == "k-receipt" {
			kReceipt, err := h.deductionUsecase.Find("kReceipt")
			if err != nil {
				return err
			}
			req.Allowances[i].MaxDeduction = kReceipt.Amount
			break
		}
	}
	return nil
}

func validateAllowance(v domain.TaxAllowance) error {
	validTypes := []string{"donation", "personalDeduction", "k-receipt"}
	if !lo.Contains(validTypes, v.AllowanceType) {
		return domain.Error{
			HttpCode: http.StatusBadRequest,
			Message:  "invalid allowance type",
		}
	}
	if v.Amount < 0 {
		return domain.Error{
			HttpCode: http.StatusBadRequest,
			Message:  "allowance amount must be less than 0",
		}
	}
	return nil
}

func buildResponse(tax float64, taxRefund float64, taxLevel []domain.TaxLevel) echo.Map {
	response := echo.Map{
		"tax":      tax,
		"taxLevel": taxLevel,
	}
	if taxRefund > 0 {
		response["taxRefund"] = taxRefund
	}
	return response
}

func (h *Handler) handleCalculationCSV() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("taxFile")
		if err != nil {
			return err
		}
		records, err := readCSV(file)
		taxs, err := h.processCSV(records)
		results, err := h.calculateTaxes(taxs)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"taxes": results,
		})
	}
}

func readCSV(file *multipart.FileHeader) ([][]string, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()
	reader := csv.NewReader(src)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (u *Handler) validateWHT(wht string, totalIncome string) (float64, float64, error) {
	whtFloat, err := strconv.ParseFloat(wht, 64)
	if err != nil {
		return 0, 0, err
	}
	totalIncomeFloat, err := strconv.ParseFloat(totalIncome, 64)
	if err != nil {
		return 0, 0, err
	}
	if whtFloat > totalIncomeFloat {
		return 0, 0, domain.Error{
			HttpCode: http.StatusBadRequest,
			Message:  "WHT must be less than total income",
		}
	}
	return totalIncomeFloat, whtFloat, nil
}

func (h *Handler) processCSV(records [][]string) ([]requestCalculation, error) {
	taxs := []requestCalculation{}
	for i, record := range records {
		if i == 0 {
			continue
		}
		totalIncome, wht, err := h.validateWHT(record[1], record[2])
		if err != nil {
			return nil, err
		}
		donation, err := strconv.ParseFloat(record[2], 64)
		allowances := []domain.TaxAllowance{}
		allowances = append(allowances, domain.TaxAllowance{
			AllowanceType: "donation",
			Amount:        donation,
		})

		taxs = append(taxs, requestCalculation{
			TotalIncome: totalIncome,
			Wht:         wht,
			Allowances:  allowances,
		})
	}
	return taxs, nil
}

func (h *Handler) calculateTaxes(taxs []requestCalculation) ([]echo.Map, error) {
	var results []echo.Map

	for _, v := range taxs {
		personal, err := h.deductionUsecase.Find("personalDeduction")
		if err != nil {
			return nil, err
		}

		v.Allowances = append(v.Allowances, domain.TaxAllowance{
			AllowanceType: "personalDeduction",
			Amount:        personal.Amount,
		})

		tax, refund, _ := h.userCase.CalculateTax(v.TotalIncome, v.Wht, v.Allowances)

		result := echo.Map{
			"totalIncome": v.TotalIncome,
			"tax":         tax,
		}

		if refund > 0 {
			result["taxRefund"] = refund
		}

		results = append(results, result)
	}

	return results, nil
}
