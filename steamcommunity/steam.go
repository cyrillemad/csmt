package steamcommunity

import (
	"time"

	net "github.com/cyrillemad/csmt/internal/httpclient"
	"github.com/cyrillemad/csmt/types"
)

type Client struct {
	*net.Client
	config Config
}

type Config struct {
	Country     string
	Currency    types.Currency
	AppID       int
	Language    types.Language
	Timeout     time.Duration
	HTTPOptions []net.Option
}

func NewClient(options ...Option) *Client {
	config := Config{
		Language: types.English,
		Country:  "RU",
		Currency: types.RUB,
		AppID:    730,
		Timeout:  5 * time.Second,
		HTTPOptions: []net.Option{
			net.WithRateLimit(2, 4),
		},
	}

	for _, option := range options {
		option(&config)
	}

	httpClient := net.NewClient("https://steamcommunity.com",
		config.HTTPOptions...,
	)

	return &Client{
		Client: httpClient,
		config: config,
	}
}
