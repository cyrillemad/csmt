package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cyrillemad/csmt/internal/debug"
	"github.com/cyrillemad/csmt/types"
	"golang.org/x/time/rate"
)

type Client struct {
	httpClient  *http.Client
	rateLimiter *rate.Limiter
	retryConfig types.RetryConfig
	userAgent   string
}

func NewClient(options ...Option) *Client {
	client := &Client{
		httpClient:  &http.Client{Timeout: 10 * time.Second},
		rateLimiter: rate.NewLimiter(rate.Limit(5), 10),
		retryConfig: types.RetryConfig{
			Attempts: 3,
			Delay:    5 * time.Millisecond * 50},
		userAgent: "csmt",
	}

	for _, option := range options {
		option(client)
	}
	return client
}

func (client *Client) do(request *http.Request, v any) error {
	if err := client.rateLimiter.Wait(request.Context()); err != nil {
		return err
	}

	for attempt := 0; attempt <= client.retryConfig.Attempts; attempt++ {
		debug.SLog(fmt.Sprintf(
			"HTTP request %s to %s, try %d",
			request.Method,
			request.URL,
			attempt))
		response, err := client.httpClient.Do(request)

		if err != nil {
			if shouldRetryError(err) {
				time.Sleep(client.retryConfig.Delay)
				continue
			}
			return err
		}

		if shouldRetryStatus(response.StatusCode) {
			if err := response.Body.Close(); err != nil {
				return err
			}
			time.Sleep(client.retryConfig.Delay)
			continue
		}

		if response.StatusCode >= 400 {
			return fmt.Errorf("http error: %d", response.StatusCode)
		}

		if v != nil {
			if err := json.NewDecoder(response.Body).Decode(v); err != nil {
				return err
			}
		}
		if err := response.Body.Close(); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("max retries exceeded: %d", client.retryConfig.Attempts)
}

//todo: add post method

func (client *Client) Get(
	ctx context.Context,
	path string,
	auth types.Authorize,
	v any,
) error {

	request, err := client.newRequest(
		ctx,
		http.MethodGet,
		path,
		nil,
		auth,
	)

	if err != nil {
		return err
	}

	return client.do(request, v)
}
