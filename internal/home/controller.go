package home

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tnfy-link/frontend/internal/links"
)

type Controller struct {
	links *links.Service
}

func (c *Controller) index(ctx *fiber.Ctx) error {
	return ctx.Render("index", contextIndex{API_URL: c.links.URL()})
}

func (c *Controller) Register(r fiber.Router) {
	r.Get("/", c.index)
}

func New(links *links.Service) *Controller {
	if links == nil {
		panic("links service is nil")
	}

	return &Controller{
		links: links,
	}
}
