package deductions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewHandler(route *echo.Group) *Handler {
	handler := Handler{}
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

		return c.JSON(http.StatusOK, echo.Map{
			"message": req.Amount,
		})
	}
}
