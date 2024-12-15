package config

import (
	"github.com/tnfy-link/frontend/internal/links"
	"github.com/tnfy-link/frontend/pkg/core/http"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"config",
	fx.Provide(New),
	fx.Provide(func(c Config) http.Config {
		return http.Config{
			Address:     c.Http.Address,
			ProxyHeader: c.Http.ProxyHeader,
			Proxies:     c.Http.Proxies,
		}
	}),
	fx.Provide(func(c Config) links.Config {
		return links.Config{
			URL:     c.Links.URL,
			Timeout: c.Links.Timeout,
		}
	}),
)
