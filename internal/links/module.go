package links

import (
	"context"

	"github.com/tnfy-link/client-go/api"
	"github.com/tnfy-link/client-go/queue"
	"github.com/tnfy-link/frontend/pkg/cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func Module() fx.Option {
	return fx.Module(
		"links",
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("links")
		}),
		fx.Provide(func(config Config) *api.Client {
			return api.New(config.URL)
		}, fx.Private),
		fx.Provide(queue.NewStatsQueue, fx.Private),
		fx.Provide(
			func() Cache {
				return cache.New[api.Link]()
			},
			fx.Private,
		),
		fx.Provide(New),
		fx.Invoke(func(cache Cache, lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					go cache.Start()
					return nil
				},
				OnStop: func(_ context.Context) error {
					cache.Stop()
					return nil
				},
			})
		}),
	)
}
