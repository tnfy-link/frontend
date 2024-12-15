package views

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var (
	//go:embed static
	Static embed.FS
	//go:embed templates
	Templates embed.FS
)

func New() (fiber.Views, error) {
	sub, err := fs.Sub(Templates, "templates")
	if err != nil {
		return nil, fmt.Errorf("failed to load templates files: %w", err)
	}

	return html.NewFileSystem(http.FS(sub), ".html"), nil
}
