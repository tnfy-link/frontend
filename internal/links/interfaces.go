package links

import (
	"time"

	"github.com/tnfy-link/client-go/api"
)

type Cache interface {
	Start()
	Set(key string, value api.Link, ttl time.Duration)
	Get(key string) (api.Link, bool)
	Stop()
}
