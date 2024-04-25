package infrastucture

import (
	"net/http"

	"github.com/h3xry/assessment-tax/pkg/deductions"
	"github.com/labstack/echo/v4"
)

func (s *Server) initRoutes() {
	s.Engine.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	s.Engine.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})

	deductionsRepository := deductions.NewRepository(s.DB)
	deductionsUsecase := deductions.NewUseCase(deductionsRepository)

	admin := s.Engine.Group("/admin")
	{

		deductions.NewHandler(admin.Group("/deductions"), deductionsUsecase)
	}
}
