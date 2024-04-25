package deductions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewHandler(route *echo.Group) *Handler {
	handler := Handler{}
	route.POST("/k-receipt", handler.setKRecepit)
	return &handler
}

func (h *Handler) setKRecepit(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "not implemented yet",
	})
}
