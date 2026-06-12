package httpclient

import (
	"time"

	"golang.org/x/time/rate"
)

type Option func(client *Client)

func WithUserAgent(userAgent string) Option {
	return func(client *Client) {
		if client.userAgent != "" {
			client.userAgent = userAgent
		}
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(client *Client) {
		if timeout > time.Duration(0) {
			client.httpClient.Timeout = timeout
		}
	}
}

func WithRateLimit(rps int, burst int) Option {
	return func(client *Client) {
		if rps <= 0 {
			rps = 1
		}
		if burst <= 0 {
			burst = rps * 2
		}
		client.rateLimiter = rate.NewLimiter(rate.Limit(rps), burst)
	}
}

func WithRetryCount(count int) Option {
	return func(client *Client) {
		if count < 0 {
			count = 0
		}
		client.retryConfig.Attempts = count
	}
}

func WithRetryDelay(delay time.Duration) Option {
	return func(client *Client) {
		if delay < 0 {
			delay = time.Duration(0)
		}
		client.retryConfig.Delay = delay
	}
}
