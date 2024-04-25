package infrastucture

import (
	"context"
	"fmt"

	"github.com/h3xry/assessment-tax/internal/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Server struct {
	Engine *echo.Echo
	Config *config.ENV
	DB     *gorm.DB
}

func NewServer(lc fx.Lifecycle, cfg *config.ENV, db *gorm.DB) *Server {
	s := Server{
		Engine: echo.New(),
		Config: cfg,
		DB:     db,
	}
	s.initRoutes()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				s.Engine.Logger.Fatal(s.Engine.Start(fmt.Sprintf(":%s", s.Config.Port)))
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("\nshutting down the server")
			return s.Engine.Shutdown(ctx)
		},
	})
	return &s
}
