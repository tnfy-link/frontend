package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnfy-link/core/http"
	"github.com/tnfy-link/core/http/jsonify"
	"github.com/tnfy-link/core/logger"
	"github.com/tnfy-link/core/redis"
	"github.com/tnfy-link/core/validator"
	"github.com/tnfy-link/frontend/internal/api"
	"github.com/tnfy-link/frontend/internal/config"
	"github.com/tnfy-link/frontend/internal/home"
	"github.com/tnfy-link/frontend/internal/links"
	"github.com/tnfy-link/frontend/internal/views"
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
			return *(&http.Options{}).
				WithViews(views).
				WithErrorHandler(http.NewViewsErrorHandler(logger, "error", "layouts/main"))
		}),
		http.Module,
		validator.Module,
		redis.Module,
		// App Modules
		config.Module,
		links.Module,
		views.Module,
		home.Module,
		api.Module,
		// Kickstarter
		fx.Invoke(func(app *fiber.App, home *home.Controller, api *api.API) {
			home.Register(app)
			api.Register(app.Group("/api").Use(jsonify.New()))

			app.Use(func(c *fiber.Ctx) error {
				return fiber.NewError(fiber.StatusNotFound, "Not Found")
			})
		}),
	).
		Run()
}
