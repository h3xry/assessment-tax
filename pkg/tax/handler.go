package tax

import (
	"encoding/csv"
	"errors"
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
		personal, err := h.deductionUsecase.Find("personalDeduction")
		if err != nil {
			return err
		}
		req.Allowances = append(req.Allowances, domain.TaxAllowance{
			AllowanceType: "personalDeduction",
			Amount:        personal.Amount,
			MaxDeduction:  personal.Amount,
		})
		for i, v := range req.Allowances {
			if ok := lo.Contains([]string{"donation", "personalDeduction", "k-receipt"}, v.AllowanceType); !ok {
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
			if v.AllowanceType == "k-receipt" {
				kReceipt, err := h.deductionUsecase.Find("kReceipt")
				if err != nil {
					return err
				}
				req.Allowances[i].MaxDeduction = kReceipt.Amount
				break
			}
		}
		tax, taxRefund, taxLevel := h.userCase.CalculateTax(req.TotalIncome, req.Wht, req.Allowances)
		response := echo.Map{
			"tax":      tax,
			"taxLevel": taxLevel,
		}
		if taxRefund > 0 {
			response["taxRefund"] = taxRefund
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (h *Handler) handleCalculationCSV() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("taxFile")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		reader := csv.NewReader(src)
		records, err := reader.ReadAll()
		if err != nil {
			return err
		}
		taxs := []requestCalculation{}
		for i, record := range records {
			if i == 0 {
				continue
			}
			totalIncome, err := strconv.ParseFloat(record[0], 64)
			if err != nil {
				return err
			}
			wht, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				return err
			}
			if wht >= totalIncome {
				return errors.New("with holding tax must be less than total income")
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
		result := []echo.Map{}
		for _, v := range taxs {
			personal, err := h.deductionUsecase.Find("personalDeduction")
			if err != nil {
				return err
			}
			v.Allowances = append(v.Allowances, domain.TaxAllowance{
				AllowanceType: "personalDeduction",
				Amount:        personal.Amount,
			})
			tax, refund, _ := h.userCase.CalculateTax(v.TotalIncome, v.Wht, v.Allowances)
			if refund > 0 {
				result = append(result, echo.Map{
					"totalIncome": v.TotalIncome,
					"tax":         tax,
					"taxRefund":   refund,
				})
				continue
			}
			result = append(result, echo.Map{
				"totalIncome": v.TotalIncome,
				"tax":         tax,
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"taxes": result,
		})
	}
}
