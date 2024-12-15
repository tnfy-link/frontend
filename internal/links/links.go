package links

import (
	"context"
	"time"

	"github.com/tnfy-link/frontend/pkg/api"
)

type Service struct {
	api *api.Client

	timeout time.Duration
}

func (s *Service) Shorten(ctx context.Context, targetURL string) (api.Link, error) {
	ctx, cancel := s.prepareContext(ctx)
	defer cancel()

	res, err := s.api.Shorten(ctx, targetURL)
	return res.Link, err
}

func (s *Service) prepareContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, s.timeout)
}

func New(api *api.Client, config Config) *Service {
	if api == nil {
		panic("api client is nil")
	}

	return &Service{
		api: api,

		timeout: config.Timeout,
	}
}
