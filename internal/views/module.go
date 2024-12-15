package views

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"views",
	fx.Provide(New),
	fx.Invoke(func(app *fiber.App) {
		app.Use("/static", filesystem.New(filesystem.Config{
			Root:       http.FS(Static),
			PathPrefix: "static",
			Browse:     false,
			MaxAge:     3600,
		}))
	}),
)
