package home

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Module() fx.Option {
	return fx.Module(
		"home",
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("home")
		}),
		fx.Provide(New),
	)
}
