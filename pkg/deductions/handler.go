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
	route.POST("/k-receipt", handler.setKReceipt())
	return &handler
}

type requestSetKReceipt struct {
	Amount float64 `json:"amount" validate:"required,min=1,max=100000"`
}

func (h *Handler) setKReceipt() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(requestSetKReceipt)
		if err := c.Bind(req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}
		deduction, err := h.useCase.Find("kReceipt")
		if err != nil {
			return err
		}
		deduction.Amount = req.Amount
		if err := h.useCase.Update(deduction); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"kReceipt": deduction.Amount,
		})
	}
}

type requestSetPersonal struct {
	Amount float64 `json:"amount" validate:"required,min=10000,max=100000"`
}

func (h *Handler) setPersonal() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(requestSetPersonal)
		if err := c.Bind(req); err != nil {
			return err
		}
		if err := c.Validate(req); err != nil {
			return err
		}
		deduction, err := h.useCase.Find("personalDeduction")
		if err != nil {
			return err
		}
		deduction.Amount = req.Amount
		if err := h.useCase.Update(deduction); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"personalDeduction": deduction.Amount,
		})
	}
}
