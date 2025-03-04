package api

import "github.com/gofiber/fiber/v2"

type API struct {
	links *Links
}

func (a *API) Register(r fiber.Router) {
	a.links.Register(r.Group("/links"))
}

func New(links *Links) *API {
	if links == nil {
		panic("links is nil")
	}

	return &API{
		links: links,
	}
}
