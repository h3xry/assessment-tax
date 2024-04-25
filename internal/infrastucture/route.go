package infrastucture

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) initRoutes() {
	s.Engine.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	s.Engine.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
}
