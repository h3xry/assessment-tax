package tax

import (
	"net/http"

	"github.com/h3xry/assessment-tax/pkg/domain"
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler(route *echo.Group) *Handler {
	handler := Handler{}
	route.POST("/calculation", handler.handleCalculation())
	return &Handler{}
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
		return c.JSON(http.StatusOK, echo.Map{
			"tax": 29000.0,
		})
	}
}
