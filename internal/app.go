package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnfy-link/frontend/internal/config"
	"github.com/tnfy-link/frontend/internal/home"
	"github.com/tnfy-link/frontend/internal/links"
	"github.com/tnfy-link/frontend/internal/views"
	"github.com/tnfy-link/frontend/pkg/core/http"
	"github.com/tnfy-link/frontend/pkg/core/logger"
	"github.com/tnfy-link/frontend/pkg/core/validator"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Run() {
	fx.New(
		// Core Modules
		logger.Module,
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			logOption := fxevent.ZapLogger{Logger: logger}
			logOption.UseLogLevel(zapcore.DebugLevel)
			return &logOption
		}),
		fx.Provide(func(views fiber.Views, logger *zap.Logger) http.Options {
			opts := http.Options{}
			http.WithViews(views)(&opts)
			http.WithErrorHandler(http.NewViewsErrorHandler(logger, "error"))(&opts)
			http.WithGetOnly()(&opts)
			return opts
		}),
		http.Module,
		validator.Module,
		// App Modules
		config.Module,
		links.Module,
		views.Module,
		home.Module,
		// Kickstarter
		fx.Invoke(func(app *fiber.App, home *home.Controller) {
			home.Register(app)
		}),
	).
		Run()
}
