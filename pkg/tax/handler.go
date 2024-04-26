package tax

import "github.com/labstack/echo/v4"

type Handler struct{}

func NewHandler(route *echo.Group) *Handler {
	handler := Handler{}
	route.POST("/calculation", handler.handleCalculation())
	return &Handler{}
}

func (h *Handler) handleCalculation() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
