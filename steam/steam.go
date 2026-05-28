package steam

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
	Country  string
	Currency types.Currency
	AppID    int
	Language types.Language
	Timeout  time.Duration
}

func NewClient(options ...Option) *Client {
	config := Config{
		Language: types.English,
		Country:  "RU",
		Currency: types.RUB,
		AppID:    730,
		Timeout:  5 * time.Second,
	}

	httpClient := net.NewClient("https://steamcommunity.com",
		net.WithRateLimit(8, 15),
	)

	for _, option := range options {
		option(&config)
	}

	return &Client{
		Client: httpClient,
		config: config,
	}
}
