package main

import (
	"github.com/h3xry/assessment-tax/internal/config"
	"github.com/h3xry/assessment-tax/internal/infrastucture"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.NopLogger,
		fx.Provide(config.NewConfig),
		fx.Provide(infrastucture.NewServer),
		fx.Invoke(func(s *infrastucture.Server) {}),
	).Run()
}
