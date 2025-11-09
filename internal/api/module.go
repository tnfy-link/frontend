package api

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Module() fx.Option {
	return fx.Module(
		"api",
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("api")
		}),
		fx.Provide(NewLinks, New),
	)
}
