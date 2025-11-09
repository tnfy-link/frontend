package config

import (
	"fmt"
	"time"

	"github.com/tnfy-link/core/config"
)

type HTTPConfig struct {
	Address     string   `envconfig:"HTTP__ADDRESS"`
	ProxyHeader string   `envconfig:"HTTP__PROXY_HEADER"`
	Proxies     []string `envconfig:"HTTP__PROXIES"`
}

type QueueConfig struct {
	URL string `envconfig:"QUEUE__URL"`
}

type LinksConfig struct {
	URL     string        `envconfig:"LINKS__URL"`
	Timeout time.Duration `envconfig:"LINKS__TIMEOUT" default:"300ms"`
}

type Config struct {
	HTTP  HTTPConfig
	Queue QueueConfig
	Links LinksConfig
}

func Default() Config {
	//nolint:exhaustruct,mnd // default values
	return Config{
		HTTP: HTTPConfig{
			Address: "127.0.0.1:3001",
		},
		Queue: QueueConfig{
			URL: "redis://localhost:6379/0",
		},
		Links: LinksConfig{
			URL:     "http://localhost:3000/api/",
			Timeout: time.Millisecond * 300,
		},
	}
}

func New() (Config, error) {
	instance := Default()

	if err := config.Load(&instance); err != nil {
		return instance, fmt.Errorf("failed to load config: %w", err)
	}

	return instance, nil
}
