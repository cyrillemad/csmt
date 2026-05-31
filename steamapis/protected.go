package steamapis

import (
	"time"

	net "github.com/cyrillemad/csmt/internal/httpclient"
)

type AuthClient struct {
	*Client
	apikey string
}

func NewAuthClient(
	httpClient *net.Client,
	apikey string,
	options ...Option) *AuthClient {
	config := Config{
		AppID:   730,
		APIPath: "https://api.steamapis.com/",
		Timeout: 5 * time.Second,
		HTTPOptions: []net.Option{
			net.WithRateLimit(2, 4),
		},
	}

	for _, option := range options {
		option(&config)
	}

	client := NewClient(
		httpClient,
		options...)
	return &AuthClient{
		Client: client,
		apikey: apikey,
	}
}
