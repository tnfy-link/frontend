package links

import (
	"context"
	"time"

	"github.com/tnfy-link/client-go/api"
)

type Service struct {
	api    *api.Client
	apiURL string

	timeout time.Duration
}

func (s *Service) URL() string {
	return s.apiURL
}

func (s *Service) Get(ctx context.Context, linkID string) (api.Link, error) {
	ctx, cancel := s.prepareContext(ctx)
	defer cancel()

	res, err := s.api.GetLink(ctx, linkID)
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
		api:    api,
		apiURL: config.URL,

		timeout: config.Timeout,
	}
}
