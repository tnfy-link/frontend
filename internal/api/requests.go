package api

type LinksPostRequest struct {
	TargetURL string `json:"targetUrl" validate:"required,http_url"`
}
