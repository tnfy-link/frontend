package config

import (
	"time"

	"github.com/tnfy-link/core/config"
)

type HttpConfig struct {
	Address     string   `envconfig:"HTTP__ADDRESS"`
	ProxyHeader string   `envconfig:"HTTP__PROXY_HEADER"`
	Proxies     []string `envconfig:"HTTP__PROXIES"`
}

type LinksConfig struct {
	URL     string        `envconfig:"LINKS__URL"`
	Timeout time.Duration `envconfig:"LINKS__TIMEOUT" default:"300ms"`
}

type Config struct {
	Http  HttpConfig
	Links LinksConfig
}

var instance = Config{
	Http: HttpConfig{Address: "127.0.0.1:3001"},
	Links: LinksConfig{
		URL:     "http://localhost:3000/api/",
		Timeout: time.Millisecond * 300,
	},
}

func New() (Config, error) {
	return instance, config.Load(&instance)
}
