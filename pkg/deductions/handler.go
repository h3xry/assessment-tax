package deductions

import (
	"net/http"

	"github.com/h3xry/assessment-tax/pkg/domain"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	useCase domain.DeductionsUsecase
}

func NewHandler(route *echo.Group, useCase domain.DeductionsUsecase) *Handler {
	handler := Handler{
		useCase: useCase,
	}
	route.POST("/k-receipt", handler.setKRecepit())
	return &handler
}

type requestSetKRecepit struct {
	Amount float64 `json:"amount" validate:"required,min=1,max=100000"`
}

func (h *Handler) setKRecepit() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(requestSetKRecepit)
		if err := c.Bind(req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}
		deduction, err := h.useCase.Find("K-Recepit")
		if err != nil {
			return err
		}
		deduction.Amount = req.Amount
		if err := h.useCase.Update(deduction); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": deduction.Amount,
		})
	}
}
