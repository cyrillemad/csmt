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
	httpClient    *http.Client
	rateLimiter   *rate.Limiter
	retryAttempts int
	retryDelay    time.Duration
	userAgent     string
}

func NewClient(options ...Option) *Client {
	client := &Client{
		httpClient:    &http.Client{Timeout: 10 * time.Second},
		rateLimiter:   rate.NewLimiter(rate.Limit(5), 10),
		retryAttempts: 3,
		retryDelay:    time.Millisecond * 50,
		userAgent:     "csmt",
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

	for attempt := 0; attempt <= client.retryAttempts; attempt++ {
		debug.SLog(fmt.Sprintf(
			"HTTP request %s to %s, try %d",
			request.Method,
			request.URL,
			attempt))
		response, err := client.httpClient.Do(request)
		defer response.Body.Close()

		if err != nil {
			if shouldRetryError(err) {
				time.Sleep(client.retryDelay)
				continue
			}
			return err
		}

		if shouldRetryStatus(response.StatusCode) {
			if err := response.Body.Close(); err != nil {
				return err
			}
			time.Sleep(client.retryDelay)
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

		return nil
	}
	return fmt.Errorf("max retries exceeded: %d", client.retryAttempts)
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
