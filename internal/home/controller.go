package home

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tnfy-link/frontend/internal/links"
)

type Controller struct {
	links *links.Service
}

func (c *Controller) index(ctx *fiber.Ctx) error {
	return ctx.Render("index", nil, "layouts/main")
}

func (c *Controller) post(ctx *fiber.Ctx) error {
	req := PostHomeRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}

	l, err := c.links.Shorten(ctx.Context(), req.TargetURL)
	if err != nil {
		return fmt.Errorf("failed to shorten link: %w", err)
	}

	return ctx.Render("result", l, "layouts/main")
}

func (c *Controller) Register(r fiber.Router) {
	r.Get("/", c.index)
	r.Post("/", c.post)
}

func New(links *links.Service) *Controller {
	if links == nil {
		panic("links service is nil")
	}

	return &Controller{
		links: links,
	}
}
