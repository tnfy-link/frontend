package links

import (
	"github.com/tnfy-link/client-go/api"
	"github.com/tnfy-link/client-go/queue"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"links",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("links")
	}),
	fx.Provide(func(config Config) *api.Client {
		return api.New(config.URL)
	}, fx.Private),
	fx.Provide(queue.NewStatsQueue, fx.Private),
	fx.Provide(New),
)
