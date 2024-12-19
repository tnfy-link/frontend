package home

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tnfy-link/frontend/internal/links"
	"go.uber.org/zap"
)

type Controller struct {
	links *links.Service

	log *zap.Logger
}

func (c *Controller) index(ctx *fiber.Ctx) error {
	return ctx.Render("index", contextIndex{API_URL: c.links.URL()}, "layouts/main")
}

func (c *Controller) redirect(ctx *fiber.Ctx) error {
	linkID := ctx.Params("id")
	link, err := c.links.Get(ctx.Context(), linkID)
	if err != nil {
		c.log.Error("failed to get link", zap.Error(err))
		return fiber.NewError(fiber.StatusNotFound, linkNotFoundMessage)
	}

	go func(id, query string) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		if err := c.links.Redirect(ctx, id, query); err != nil {
			c.log.Error("failed to register redirect", zap.Error(err))
		}
	}(strings.Clone(linkID), strings.Clone(ctx.Context().QueryArgs().String()))

	return ctx.Redirect(link.TargetURL, fiber.StatusTemporaryRedirect)
}

func (c *Controller) Register(r fiber.Router) {
	r.Get("/", c.index)
	r.Get("/:id", c.redirect)
}

func New(links *links.Service, log *zap.Logger) *Controller {
	if links == nil {
		panic("links service is nil")
	}

	if log == nil {
		panic("logger is nil")
	}

	return &Controller{
		links: links,

		log: log,
	}
}
