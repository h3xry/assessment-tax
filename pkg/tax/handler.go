package tax

import (
	"net/http"

	"github.com/h3xry/assessment-tax/pkg/domain"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	userCase domain.TaxUsecase
}

func NewHandler(route *echo.Group, userCase domain.TaxUsecase) *Handler {
	handler := Handler{
		userCase: userCase,
	}
	route.POST("/calculation", handler.handleCalculation())
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
		tax := h.userCase.CalculateTax(req.TotalIncome, req.Wht, req.Allowances)
		return c.JSON(http.StatusOK, echo.Map{
			"tax": tax,
		})
	}
}
