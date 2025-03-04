package links

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/tnfy-link/client-go/api"
	"github.com/tnfy-link/client-go/queue"
	"go.uber.org/zap"
)

const (
	maxUTMLabelValueLength = 64
)

var utmLabels = map[string]string{
	"utm_source":   "source",
	"utm_medium":   "medium",
	"utm_campaign": "campaign",
}

type Service struct {
	api   *api.Client
	queue *queue.StatsQueue

	cache Cache

	log *zap.Logger

	timeout time.Duration
}

func (s *Service) Shorten(ctx context.Context, targetURL string) (api.Link, error) {
	ctx, cancel := s.prepareContext(ctx)
	defer cancel()

	resp, err := s.api.Shorten(ctx, targetURL)
	if err != nil {
		s.log.Error("failed to shorten URL", zap.String("targetURL", targetURL), zap.Error(err))
		return api.Link{}, fmt.Errorf("failed to shorten URL: %w", err)
	}

	return resp.Link, nil
}

func (s *Service) Get(ctx context.Context, linkID string) (api.Link, error) {
	if res, ok := s.cache.Get(linkID); ok {
		s.log.Debug("got link from cache", zap.String("id", linkID))
		return res, nil
	}

	ctx, cancel := s.prepareContext(ctx)
	defer cancel()

	res, err := s.api.GetLink(ctx, linkID)
	if err == nil && !res.Link.ValidUntil.IsZero() && time.Until(res.Link.ValidUntil) > time.Second {
		s.cache.Set(linkID, res.Link, time.Until(res.Link.ValidUntil))
	}

	return res.Link, err
}

func (s *Service) Redirect(ctx context.Context, id, query string) error {
	values, err := url.ParseQuery(query)
	if err != nil {
		// not a fatal error, just log
		s.log.Warn("failed to parse query", zap.Error(err))
	}

	labels := labels{}

	for k, v := range utmLabels {
		if val := values.Get(k); val != "" {
			if err := validateUTMValue(val); err != nil {
				s.log.Warn("invalid utm value", zap.String("id", id), zap.String("label", v), zap.String("value", val), zap.Error(err))
				continue
			}

			if len(val) > maxUTMLabelValueLength {
				s.log.Warn("label value too long", zap.String("id", id), zap.String("label", v), zap.String("value", val))
				val = val[:maxUTMLabelValueLength]
			}
			labels[v] = val
		}
	}

	return s.queue.Enqueue(ctx, queue.StatsIncrEvent{
		LinkID: id,
		Labels: labels,
	})
}

func (s *Service) prepareContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, s.timeout)
}

func New(api *api.Client, queue *queue.StatsQueue, cache Cache, log *zap.Logger, config Config) *Service {
	if api == nil {
		panic("api client is nil")
	}

	if queue == nil {
		panic("queue is nil")
	}

	if cache == nil {
		panic("cache is nil")
	}

	if log == nil {
		panic("logger is nil")
	}

	return &Service{
		api:   api,
		queue: queue,

		cache: cache,

		log: log,

		timeout: config.Timeout,
	}
}
