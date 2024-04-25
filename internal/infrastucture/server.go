package infrastucture

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Server struct {
	Engine *echo.Echo
}

func NewServer(lc fx.Lifecycle) *Server {
	s := Server{
		Engine: echo.New(),
	}
	s.initRoutes()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				s.Engine.Logger.Fatal(s.Engine.Start(fmt.Sprintf(":%s", viper.GetString("PORT"))))
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
