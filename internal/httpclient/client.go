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
	userAgent   string
}

func NewClient(options ...Option) *Client {
	client := &Client{
		httpClient:  &http.Client{Timeout: 10 * time.Second},
		rateLimiter: rate.NewLimiter(rate.Limit(5), 10),
		userAgent:   "csmt",
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

	debug.SLog(fmt.Sprintf("HTTP request %s to %s", request.Method, request.URL))
	response, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return fmt.Errorf("http error: %s", response.Status)
	}

	if v != nil {
		if err := json.NewDecoder(response.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}

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
