package views

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"go.uber.org/fx"
)

func Module() fx.Option {
	const staticMaxAge = 3600

	return fx.Module(
		"views",
		fx.Provide(New),
		fx.Invoke(func(app *fiber.App) {
			app.Use(favicon.New(favicon.Config{
				File:       "static/favicon.ico",
				FileSystem: http.FS(Static),
			}))
			app.Use("/static", filesystem.New(filesystem.Config{
				Root:       http.FS(Static),
				PathPrefix: "static",
				Browse:     false,
				MaxAge:     staticMaxAge,
			}))
		}),
	)
}
