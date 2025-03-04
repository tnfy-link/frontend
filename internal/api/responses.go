package api

import "time"

type LinksPostResponse struct {
	URL        string    `json:"url"`
	ValidUntil time.Time `json:"validUntil"`
}
