package steamapis

import (
	"time"

	net "github.com/cyrillemad/csmt/internal/httpclient"
)

type Client struct {
	*net.Client
	config Config
}

type Config struct {
	AppID       int
	Timeout     time.Duration
	HTTPOptions []net.Option
} // super mvp - it must be reworked when i get internet.

func NewClient(options ...Option) *Client {
	config := Config{
		AppID:   730,
		Timeout: 5 * time.Second,
		HTTPOptions: []net.Option{
			net.WithRateLimit(2, 4),
		},
	}

	for _, option := range options {
		option(&config)
	}

	httpClient := net.NewClient("https://api.steamapis.com",
		config.HTTPOptions...,
	)

	return &Client{
		Client: httpClient,
		config: config,
	}
}
