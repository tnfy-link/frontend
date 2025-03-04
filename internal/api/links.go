package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tnfy-link/core/handler"
	"github.com/tnfy-link/frontend/internal/links"
	"go.uber.org/zap"
)

type Links struct {
	handler.Base

	links *links.Service
}

func (l *Links) post(ctx *fiber.Ctx) error {
	req := LinksPostRequest{}
	if err := l.BodyParserValidator(ctx, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	link, err := l.links.Shorten(ctx.Context(), req.TargetURL)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to shorten link"})
	}

	return ctx.JSON(LinksPostResponse{
		URL:        link.URL,
		ValidUntil: link.ValidUntil,
	})
}

func (l *Links) Register(r fiber.Router) {
	r.Post("/", l.post)
}

func NewLinks(links *links.Service, validator *validator.Validate, log *zap.Logger) *Links {
	if links == nil {
		panic("links service is nil")
	}

	if log == nil {
		panic("logger is nil")
	}

	return &Links{
		links: links,
		Base: handler.Base{
			Logger:    log,
			Validator: validator,
		},
	}
}
