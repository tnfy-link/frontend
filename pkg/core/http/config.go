package http

import "github.com/gofiber/fiber/v2"

type Config struct {
	Address     string
	ProxyHeader string
	Proxies     []string
}

type Options struct {
	getOnly      bool
	views        fiber.Views
	errorHandler fiber.ErrorHandler
}

type Option func(*Options) *Options

func WithGetOnly() Option {
	return func(o *Options) *Options {
		o.getOnly = true
		return o
	}
}

func WithViews(views fiber.Views) Option {
	return func(o *Options) *Options {
		o.views = views
		return o
	}
}

func WithErrorHandler(handler fiber.ErrorHandler) Option {
	return func(o *Options) *Options {
		o.errorHandler = handler
		return o
	}
}
