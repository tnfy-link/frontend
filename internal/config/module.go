package config

import (
	"github.com/tnfy-link/core/http"
	"github.com/tnfy-link/core/redis"
	"github.com/tnfy-link/frontend/internal/links"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"config",
		fx.Provide(New),
		fx.Provide(func(c Config) http.Config {
			return http.Config{
				Address:     c.HTTP.Address,
				ProxyHeader: c.HTTP.ProxyHeader,
				Proxies:     c.HTTP.Proxies,
			}
		}),
		fx.Provide(func(c Config) redis.Config {
			return redis.Config{
				URL: c.Queue.URL,
			}
		}),
		fx.Provide(func(c Config) links.Config {
			return links.Config{
				URL:     c.Links.URL,
				Timeout: c.Links.Timeout,
			}
		}),
	)
}
