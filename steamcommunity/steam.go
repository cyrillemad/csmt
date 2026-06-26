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
	Country          string
	Currency         types.Currency
	AppID            int
	APIPath          string
	Language         types.Language
	Timeout          time.Duration
	EmptyFieldsRetry types.RetryConfig
	HTTPOptions      []net.Option
	Cookie           string
}

func NewClient(
	httpClient *net.Client,
	options ...Option) *Client {
	config := Config{
		Language: types.English,
		Country:  "RU",
		Currency: types.RUB,
		AppID:    730,
		APIPath:  "https://steamcommunity.com/",
		Timeout:  5 * time.Second,
		EmptyFieldsRetry: types.RetryConfig{
			Delay:    500 * time.Millisecond,
			Attempts: 5,
		},
		HTTPOptions: []net.Option{
			net.WithRateLimit(2, 4),
		},
	}

	for _, option := range options {
		option(&config)
	}

	return &Client{
		Client: httpClient,
		config: config,
	}
}
