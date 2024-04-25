package infrastucture

import (
	"crypto/subtle"
	"net/http"

	"github.com/h3xry/assessment-tax/pkg/deductions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
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
		admin.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(username), []byte(viper.GetString("ADMIN_USERNAME"))) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(viper.GetString("ADMIN_PASSWORD"))) == 1 {
				return true, nil
			}
			return false, nil
		}))
		deductions.NewHandler(admin.Group("/deductions"), deductionsUsecase)
	}
}
