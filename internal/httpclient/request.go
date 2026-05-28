package httpclient

import (
	"context"
	"io"
	"net/http"
)

func (client *Client) newRequest(
	ctx context.Context,
	method string,
	path string,
	body io.Reader,
) (*http.Request, error) {

	url := client.baseURL + path

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

	return request, nil
}
