package httpclient

import (
	"context"
	"io"
	"net/http"

	"github.com/cyrillemad/csmt/types"
)

func (client *Client) newRequest(
	ctx context.Context,
	method string,
	path string,
	body io.Reader,
	auth types.Authorize,
) (*http.Request, error) {

	url := path

	request, err := http.NewRequestWithContext(
		ctx,
		method,
		url,
		body,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", client.userAgent)
	request.Header.Set("Accept", "application/json")
	if auth.Key != "" &&
		auth.Header != "" {
		request.Header.Set(
			auth.Header,
			auth.Key,
		)
	}

	return request, nil
}
